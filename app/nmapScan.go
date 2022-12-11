package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Ullaakut/nmap/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Checking the open ports
func nmapScan(update uint16, bot *tgbotapi.BotAPI, group int64, host struct {
	Host  string
	Ports []uint16
}) {
	ports := ""
	for _, port := range host.Ports {
		ports += strconv.FormatUint(uint64(port), 10) + ","
	}

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(host.Host),
		nmap.WithPorts(ports),
	)
	if err != nil {
		log.Fatalf("Unable to create nmap scanner: %v", err)
	}

	result, warnings, err := scanner.Run()
	if err != nil {
		log.Fatalf("Unable to run nmap scan: %v", err)
	}

	if warnings != nil {
		log.Printf("Warnings: \n %v", warnings)
	}

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		for _, port := range host.Ports {
			if port.State.State == "open" {
				fmt.Printf("Host %q  has %d/%s %s %s\n", host.Addresses[0], port.ID, port.Protocol, port.State, port.Service.Name)

				msg := tgbotapi.NewMessage(group, fmt.Sprintf("Host %q  has %d/%s %s %s\n", host.Addresses[0], port.ID, port.Protocol, port.State, port.Service.Name))
				bot.Send(msg)
			}
		}
	}

	time.Sleep(time.Duration(update) * time.Second)
}
