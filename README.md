# GO-DB 

## CMD

## Запуск рутинных операций
Для запуска рутинных операций, такие как статические анализаторы, тесты, всякие генерации и т.п. существует скрипт
run.sh, посмотрите его using, чтобы ознакомиться с его возможностями.
```bash
./run.sh using
```

### Запуск контейнеров для разработки API
```shell script
docker build -t go-db:develop .
```

[//]: # (```shell script)

[//]: # (docker-compose -f deployments/docker-compose.yaml up -d)

[//]: # (```)

### Запуск с мейн:
```txt
Просто запустите main.go
```

### Грамматика языка запросов в виде eBNF:
```curl
query = set_command | get_command | del_command

set_command = "SET" argument argument
get_command = "GET" argument
del_command = "DEL" argument
argument    = punctuation | letter | digit { punctuation | letter | digit }

punctuation = "\*" | "/" | "_" | ...
letter      = "a" | ... | "z" | "A" | ... | "Z"
digit       = "0" | ... | "9"
```

### Пример запросов:
```curl
SET weather_2_pm cold_moscow_weather
GET /etc/nginx/config
DEL user_\*\*\*\*
```

## MAINTAINER
* Blckvia