# Как запустить (инструкция для Ubuntu 22.04)
- Установите go версии хотя бы 1.19 (например, [по этому гайду](https://go.dev/doc/install))
- Установите в переменную окружения `DENT_CONFIG_PATH` путь до [конфига](https://github.com/GandarfHSE/dentistryBackend/blob/main/utils/config/config.json):
```bash
export DENT_CONFIG_PATH=$PATH_TO_REPO/utils/config/config.json
```
- В корне репозитория выполните (создание пары RSA ключей для аутентификации):
```bash
openssl genrsa -out privatekey.pem 2048
openssl rsa -in privatekey.pem -out publickey.pem -pubout -outform PEM
```
- Установите [postgres](https://www.postgresql.org/download/)
- Создайте базу данных `test_dent_db` и [установите пароль](https://stackoverflow.com/questions/12720967/how-can-i-change-a-postgresql-user-password) `postgres` на юзера `postgres` (конфигурируется [тут](https://github.com/GandarfHSE/dentistryBackend/blob/main/utils/config/config.json)), базовые команды psql [тут](https://www.postgresqltutorial.com/postgresql-administration/psql-commands/)
- Перейдите в папку `main`
- Выполните `go build`
- Выполните `./main`

Прибить можно с помощью `Ctrl + C`

# API

Все доступные хендлеры можно найти в [core/handlers.go](https://github.com/GandarfHSE/dentistryBackend/blob/main/core/handlers.go).

Любая ручка может пятисотить.

В случае ошибки возвращается json с полем `err`.

## /hello
- input: None
- output: "Hello!\n"
- curl example: `curl localhost:8083/hello`

## /user/create
- input: json
- input format: string `login`, string `password`, int `role`: 1 (пациент), 2 (доктор), 4 (админ), 8 (разработчик)
- curl example: `curl localhost:8083/user/create -d '{"login":"kek", "password":"lol", "role":1}'`
- output: empty json

Кидает `400`, если юзер существует или если роль некорректна (не 1/2/4/8).

## /user/login
- input: json
- input format: string `login`, string `password`
- curl example: `curl localhost:8083/user/login -d '{"login":"kek", "password":"lol"}'`
- output: json with string `jwt`

Кидает `400`, если юзера не существует или если пароль некорректен.

## /user/list
Выводит список всех юзеров (доступно только для админов и разработчиков)
- input: cookie with `jwt`
- curl example: `curl localhost:8083/user/list -b $(cat cookie.txt)` (cookie.txt: `jwt=input_token_here`)
- output: json with `User` array `userlist`

Кидает `401`, если нет куки.

Кидает `403`, если роль юзера не админ и не разработчик.

## /service/create
- input: json
- input format: string `name`, string `description`, int `cost`, int `duration`
- curl example: `curl localhost:8083/service/create -d '{"name":"Чистка зубов", "description":"оч круто чистим", "cost":300, "duration":42}'`
- output: empty json

Кидает `400`, если услуга с таким именем существует.

## /service/list
Выводит список всех услуг
- input: None
- curl example: `curl localhost:8083/service/list`
- output: json with `Service` array `servicelist`
