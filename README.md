# Telegram bot for organizing constant monitoring of open ports on the network
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/grfc-ru/nmap-telegram-bot/docker-image.yml?label=BUILD%20AND%20PUBLISH%20APPLICATION&logo=GITHUB) [![GitHub](https://img.shields.io/badge/Git-Hub-purple.svg)](https://github.com/grfc-ru/nmap-telegram-bot) [![Docker](https://img.shields.io/badge/Docker-hub-2496ed.svg)](https://hub.docker.com/r/leech001/nmap-telegram-bot)

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
        - 20-23   
      exclusion:            
        hosts:
          - 8.8.8.1         #exclusion host
        ports:
          - 22              #exclusion port
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
