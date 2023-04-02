package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Update uint32
	}
	Telegram struct {
		Token string
		Group int64
	}
	Scan struct {
		Hosts []struct {
			Host      string
			Ports     []string
			Exclusion struct {
				Hosts []string
				Ports []string
			}
		}
	}
}

func main() {
	// Read config from yaml
	config := Config{}
	filename, _ := filepath.Abs("./conf/config.yaml")
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Parse yaml
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	// Telegram bot
	bot, err := tgbotapi.NewBotAPI(config.Telegram.Token)
	if err != nil {
		panic(err)
	}

	// Debug Telegram bot
	bot.Debug = false

	// Running NMAP scaner
	for _, host := range config.Scan.Hosts {
		go nmapScan(config.App.Update, bot, config.Telegram.Group, host)
	}

	botUpdate(bot, config.Scan.Hosts)
}

// Telegram bot for listening to incoming commands
func botUpdate(bot *tgbotapi.BotAPI, hosts []struct {
	Host      string
	Ports     []string
	Exclusion struct {
		Hosts []string
		Ports []string
	}
}) {

	// Create string for list scanning hosts
	listString := ""
	for _, host := range hosts {
		ports := ""
		for _, port := range host.Ports {
			ports += port + ","
		}

		listString += host.Host + " ports: " + ports + "\n"
	}

	// Telegram bot listener
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 300
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "start":
			msg.Text = "Hi, I am a NMAP scanner bot! Your (group) ID = " + strconv.FormatInt(update.Message.Chat.ID, 10)
		case "list":
			msg.Text = "Scanned hosts:\n" + listString
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
