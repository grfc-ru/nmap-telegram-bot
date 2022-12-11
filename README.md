# Telegram bot for organizing constant monitoring of open ports on the network

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
