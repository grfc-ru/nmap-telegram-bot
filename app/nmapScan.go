package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Ullaakut/nmap/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Scanning hosts for open ports
func nmapScan(update uint16, bot *tgbotapi.BotAPI, group int64, hosts struct {
	Host      string
	Ports     []string
	Exclusion struct {
		Hosts []string
		Ports []string
	}
},
) {
	ports := ""
	for _, port := range hosts.Ports {
		ports += port + ","
	}

	// Create exclusions hosts
	ehosts := ""
	for _, ehost := range hosts.Exclusion.Hosts {
		ehosts += ehost + ","
	}

	// Create exclusions ports
	eports := ""
	for _, eport := range hosts.Exclusion.Ports {
		eports += eport + ","
	}

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(hosts.Host),
		nmap.WithTargetExclusion(ehosts),
		nmap.WithPorts(ports),
		nmap.WithPortExclusions(eports),
	)
	if err != nil {
		log.Fatalf("Unable to create nmap scanner: %v", err)
	}

	// Infinity scan loop
	for {
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
					fmt.Printf("Host %q  has %d/%s %s %s port\n", host.Addresses[0], port.ID, port.Protocol, port.State, port.Service.Name)

					msg := tgbotapi.NewMessage(group, fmt.Sprintf("Host %q  has %d/%s %s %s port\n", host.Addresses[0], port.ID, port.Protocol, port.State, port.Service.Name))
					bot.Send(msg)
					time.Sleep(time.Second)
				}
			}
		}
		time.Sleep(time.Duration(update) * time.Second)
	}
}
