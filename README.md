# zenTest

Running Redis With Docker Locally

Let’s get started with this tutorial and download the redis docker image and run it using the following 2 docker commands:

$ docker pull redis
$ docker run --name redis-test-instance -p 6379:6379 -d redis

1. **По пути POST `/redis/incr` с json вида**

{
"key": "age",
"value": 19

},

подключается к redis (хост и порт указываются при запуске в параметрах `-host` и `-port`),
инкрементирует значение по ключу, указанному в "key" на значение из value", и
возвращает его в виде:

{
"value": 20

}

1. **По пути POST `/sign/hmacsha512` с json вида**

{
"text": "test",
"key": "test123"

}

и возвращает HMAC-SHA512 подпись значения из "text" по ключу "key" в виде hex строки

1. **По пути POST `/postgres/users` с json вида**

{
"name": "Alex",
"age": 21

}

создает в базе данных postgres таблицу users, если она не существует, добавляет в нее
строку (“Alex”, 21) и возвращает идентификатор добавленного пользователя в виде
{

“id”: 1
}
