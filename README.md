# Как запустить
- Установите go (например, [по этому гайду](https://howistart.org/posts/go/1/#setting-up-your-environment))
- Установите в переменную окружения `DENT_CONFIG_PATH` путь до [конфига](https://github.com/GandarfHSE/dentistryBackend/blob/main/utils/config/config.json):
```bash
export DENT_CONFIG_PATH=$PATH_TO_REPO/utils/config/config.json
```
- В корне репозитория выполните (создание пары RSA ключей для аутентификации):
```bash
openssl genrsa -out privatekey.pem 2048
openssl rsa -in privatekey.pem -out publickey.pem -pubout -outform PEM
```
- Перейдите в папку `main`
- Выполните `go build`
- Выполните `./main`
Прибить можно с помощью `Ctrl + C`

# API

Все доступные хендлеры можно найти в [core/handlers.go](https://github.com/GandarfHSE/dentistryBackend/blob/main/core/handlers.go#L13).

Любая ручка в теории может пятисотить.

## /hello
- input: None
- output: "Hello!\n"
- curl example: `curl localhost:8083/hello`

## /user/create
- input: json
- input format: string `login`, string `password`, int `role` (1, 2, 4, 8)
- curl example: `curl localhost:8083/user/create -d '{"login":"kek", "password":"lol", "role":1}'`
- output: json with int `id` or string `err`

Кидает `400`, если юзер существует или если роль некорректна.
Кидает `500`, если не получилось добавить юзера.

## /user/login
- input: json
- input format: string `login`, string `password`
- curl example: `curl localhost:8083/user/login -d '{"login":"kek", "password":"lol"}'`
- output: json with string `token` or string `err`

Кидает `400`, если юзера не существует или если пароль некорректен
Кидает `500`, если не получилось создать токен
