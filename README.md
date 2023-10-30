Техническое задание
Написать на Go и PHP сервисы для вывода данных по стране принадлежности.
a. IPv4 или IPv6 передаётся как GET-параметр.
b. Используется библиотека GeoIP2.
c. Используется кеширование в memcache.

endpoint для проверки 
./?ip={IP}   - принимает формат как ipv4 так и ipv6

инструкция по запуску

если нету .env скопируйте его из .env.example  
    `cp .env.example .env`
Установить требуемые настройки в .env  

Запустить docker compose 
    `docker compose up --build` 
после билда и поднятия можно перейти на приложение по указанному порту из .env OUTPUT_PORTNUM 


Dockerfile билдится под oracle ARM 
если приложение нужно запускать на  x86 поменяйте 
строку  6:
`RUN GOOS=linux GOARCH=arm64 go build -o testovoe main.go`
на : 
`RUN GOOS=linux GOARCH=amd64 go build -o testovoe main.go`
