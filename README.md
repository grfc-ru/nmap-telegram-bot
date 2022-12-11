# Telegram bot for organizing constant monitoring of open ports on the network
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/grfc-ru/nmap-telegram-bot/Publish%20Docker%20image?label=BUILD%20AND%20PUBLISH%20APPLICATION&logo=github) [![GitHub](https://img.shields.io/badge/Git-Hub-purple.svg)](https://github.com/grfc-ru/nmap-telegram-bot) [![Docker](https://img.shields.io/badge/Docker-hub-2496ed.svg)](https://hub.docker.com/r/leech001/nmap-telegram-bot) [![License: WTFPL](https://img.shields.io/badge/license-WTFPL-brightgreen)](https://github.com/grfc-ru/nmap-telegram-bot/blob/master/LICENSE)  

Application for organizing constant monitoring of open ports on nodes. Used to monitor erroneous configurations on network equipment or hacker activity.

## Simple use
Download the required repository;

```bash
$ git clone https://github.com/grfc-ru/nmap-telegram-bot.git
```

Change `conf/config.yaml`

Run `sudo docker-compose up -d`

## Configure

```yaml
app:
  update: 30   #time to rechecking hosts (sec)

telegram:
  token: 244516775:AAGZп55654ASsFFpbjyNA9su6gQU-Qs  #Token for you Telegram BOT
  group: 123456     # Telegram you ID or group ID (use command for BOT /start

scan:
  hosts:
    - host: 8.8.8.0/24      #scan network
      ports:
        - 3389              #scan port
    - host: google.com      #scan host
      ports:
        - 80                #scan port
        - 443               #scan port
```

## Telegram BOT command
```
/start  # Print you ID or group ID need you for config
/list   # Print scanning hosts
```
