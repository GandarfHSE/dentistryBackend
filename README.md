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
- Выполните `./main --hard` (для первого запуска, для последующих `./main`)

Прибить можно с помощью `Ctrl + C`

# Unit tests

Для создайте базу данных `test_db` (конфигурируется [тут](https://github.com/GandarfHSE/dentistryBackend/blob/main/utils/config/config_test.json)) и добавьте путь к тестовому конфигу в переменную окружения `DENT_TEST_CONFIG_PATH`.

# CLI

### --hard
Дропает и пересоздаёт все таблицы. Все данные, находящиеся в БД, будут удалены.

### --https
Сервер слушает запросы на протоколе HTTPS. Необходимо предоставить путь к сертификату в [конфиге](https://github.com/GandarfHSE/dentistryBackend/blob/main/utils/config/config.json).

# API

Все доступные хендлеры можно найти в [core/handlers.go](https://github.com/GandarfHSE/dentistryBackend/blob/main/core/handlers.go).

Любая ручка может пятисотить.

В случае ошибки возвращается json с полем `err`.

## /hello
- input: None
- output: "Hello!\n"
- curl example: `curl localhost:8083/hello`

---

## /image/upload
- input: multipart form
- input format: binary data with content of image `image_data` and string `image_extension`
- input example can be found at [test script](https://github.com/GandarfHSE/dentistryBackend/blob/main/utils/image/test_script/uploadtest.py)
- output: json with string `imageSource` - link to image

---

## /user/create
Создать юзера. Для создания юзера с ролью выше пациента, необходимо предоставить кодовое слово
- input: json
- input format: string `login`, string `password`, int `role`: 1 (пациент), 2 (доктор), 4 (админ), 8 (разработчик), string `keyword` (необязательно)
- curl example: `curl localhost:8083/user/create -d '{"login":"kek", "password":"lol", "role":2, "keyword":"kek"}'`
- output: json with string `err` (see notes: empty json in response)

Кидает `400`, если юзер существует или если роль некорректна (не 1/2/4/8).

Кидает `403`, если роль юзера выше пациента (2/4/8), но кодовое слово не корректно.

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
- output: json with User array `userlist`
- output example: `{"userlist":[{"id":1,"login":"kek","password":"7c5549bc580261e8c7b68655df72a857","role":4},{"id":2,"login":"kek1","password":"7c5549bc580261e8c7b68655df72a857","role":1}]}`

Кидает `401`, если нет куки.

Кидает `403`, если роль юзера не админ и не разработчик.

## /user/whoami
Отдаёт информацию о юзере по куке
- input: cookie with `jwt`
- curl example: `curl localhost:8083/user/whoami -b $(cat cookie.txt)` (cookie.txt: `jwt=input_token_here`)
- output: json with int `Id`, string `Login` and int `Role`
- output example: `{"id":1,"login":"kek","role":1}`

---

## /patient/create
Добавить информацию про пациента
- input: json
- input format: int `uid` (айдишник юзера из таблицы `users`), string `name`, string `passport`, string `photo` (url with photo, optional, default value - empty string)
- curl example: `curl localhost:8083/patient/create -d '{"uid":1, "name":"Carl", "passport":"1234 133742", "photo":"https://storage.yandexcloud.net/dentpicstest/images/uabxxhXirN.jpg"}'`
- output: json with string `err` (see notes: empty json in response)

Кидает `400`, если юзера с таким uid не существует, его роль не пациент или если информацию про него уже существует.

## /patient/get
Получить информацию про пациента по его uid
- input: json
- input format: int `uid`
- curl example: `curl localhost:8083/patient/get -d '{"uid":1}'`
- output: json with PatientInfo array `info`
- output example: `{"info":{"id":1,"uid":1,"name":"Carl","passport":"1234 133742","photo":"https://storage.yandexcloud.net/dentpicstest/images/uabxxhXirN.jpg"}}`

Кидает `400`, если информации про юзера с таким uid не существует

---

## /doctor/create
Добавить информацию про врача
- input: json
- input format: int `uid` (айдишник юзера из таблицы `users`), string `name`, string `post` (должность), int `exp` (типа experience - стаж работы в годах), string `photo` (url with photo, optional, default value - empty string), string `description` (optional, default value - empty string)
- curl example: `curl localhost:8083/doctor/create -d '{"uid":1, "name":"John Doe", "post":"Доктор крутой", "exp":42, "description": "Описание доктора"}'`
- output: json with string `err` (see notes: empty json in response)

Кидает `400`, если юзера с таким uid не существует, его роль не доктор или если информацию про него уже существует

## /doctor/get
Получить информацию про врача по его uid
- input: json
- input format: int `uid`
- curl example: `curl localhost:8083/doctor/get -d '{"uid":1}'`
- output: json with DoctorInfo array `info`
- output example: `{"info":{"id":1,"uid":1,"name":"John Doe","post":"Доктор крутой","exp":42,"photo":"","description":"Описание доктора"}}`

Кидает `400`, если информации про юзера с таким uid не существует

## /doctor/find/namesubstr
Получить список врачей с именем, содержащим данную подстроку (не чувствительно к регистру)
- input: json
- input format: string `name` - подстрока в имени
- curl example: `curl localhost:8083/doctor/find/namesubstr -d '{"name":"oHn"}'`
- output: json with DoctorInfo array `result`
- output example: `{"result":[{"id":1,"uid":1,"name":"John Doe","post":"Доктор крутой","exp":42,"photo":"","description":"Описание доктора"}]}`

---

## /service/create
- input: json
- input format: string `name`, string `description`, int `cost`, int `duration` (длительность в минутах)
- curl example: `curl localhost:8083/service/create -d '{"name":"Чистка зубов", "description":"Профессиональная чистка зубов", "cost":300, "duration":60}'`
- output: json with string `err` (see notes: empty json in response)

Кидает `400`, если услуга с таким именем существует.

## /service/list
Выводит список всех услуг
- input: None
- curl example: `curl localhost:8083/service/list`
- output: json with Service array `servicelist`

---

## /appointment/create/default
Создаёт запись на услугу
- input: json
- input format: int `pid` (айдишник пациента из таблицы `users`), int `did` (айдишник доктора из таблицы `users`), int `sid` (айдишник услуги из таблицы `services`), string `time` (время начала приёма в формате ISO8601)
- curl example: `curl localhost:8083/appointment/create/default -d '{"pid":1, "did":2, "sid":1, "time":"2020-12-09T16:10:53Z"}'`
- output: json with string `err` (see notes: empty json in response)

Кидает `400`, если юзеров с айди `pid` и `did` не существует или услуги с айди `sid` не существует.

Кидает `403`, если у юзеров неправильные роли (у пациента не пациент, у доктора не доктор)

Кидает `409`, если услуга не может быть создана (например, это время занято другой услугой)

## /appointment/create/patient
Создаёт запись на услугу от лица пациента
- input: json and cookie
- input format: int `did`, int `sid`, string `time`
- curl example: `curl localhost:8083/appointment/create/patient -d '{"did":2, "sid":1, "time":"2020-12-09T16:10:53Z"}' -b $(cat cookie.txt)`
- output: json with string `err` (see notes: empty json in response)

Кидает ошибки выше + `401`, если куки нет

## /appointment/create/doctor
Создаёт запись на услугу от лица доктора
- input: json and cookie
- input format: int `pid`, int `sid`, string `time`
- curl example: `curl localhost:8083/appointment/create/doctor -d '{"pid":1, "sid":1, "time":"2020-12-09T16:10:53Z"}' -b $(cat cookie.txt)`
- output: json with string `err` (see notes: empty json in response)

Кидает ошибки выше + `401`, если куки нет

## /appointment/get
Получить запись по айдишнику записи
- input: json
- input format: int `id`
- curl example: `curl localhost:8083/appointment/get -d '{"id":1}'`
- output: json with Appointment `appointment`
- output example: `{"appointment":{"id":1,"pid":1,"did":2,"sid":1,"timebegin":"2020-12-09T16:10:53Z","timeend":"2020-12-09T16:52:53Z"}}`

Кидает `400`, если записи с таким айди нет

## /appointment/list
Получить список всех записей
- input: None
- curl example: `curl localhost:8083/appointment/list`
- output: json with Appointment array `appointmentList`
- output example: `{"appointmentList":[{"id":1,"pid":1,"did":2,"sid":1,"timebegin":"2020-12-09T16:10:53Z","timeend":"2020-12-09T16:52:53Z"}]}`

## /appointment/list/patient
Получить список всех записей пациента по айди
- input: json
- input format: int `pid`
- curl example: `curl localhost:8083/appointment/list/patient -d '{"pid":1}'`
- output: json with Appointment array `appointmentList`

## /appointment/list/doctor
Получить список всех записей доктора по айди
- input: json
- input format: int `did`
- curl example: `curl localhost:8083/appointment/list/doctor -d '{"did":2}'`
- output: json with Appointment array `appointmentList`

## /appointment/free
Получить список свободных таймслотов на запись к доктору. Работает в предположении рабочих часов с 9 до 18 (issue #33).
- input: json
- input format: int `did` (айди доктора из таблицы `users`), int `sid` (айди услуги), string `date` (дата в формате `YYYY-MM-DD`)
- curl example: `curl localhost:8083/appointment/free -d '{"did":2, "sid":1, "date":"2023-08-09"}'`
- output: Timeslot array `freeTimeslots`
- output example: `{"freeTimeslots":[{"timebegin":"2023-08-09T13:00:00Z","timeend":"2023-08-09T15:00:00Z"},{"timebegin":"2023-08-09T15:00:00Z","timeend":"2023-08-09T17:00:00Z"}]}`

---

## /clinic/create
Создать запись о клинике. Выполняется только от админа/разработчика
- input: json and cookie
- input format: string `name`, string `address`, string `phone`
- curl example: `curl localhost:8083/clinic/create -d '{"name":"Клиника нейм", "address":"ул. Пушкина, д. 42", "phone":"88005553535"}' -b $(cat cookie.txt)`
- output: json with string `err` (see notes: empty json in response)

Кидает `401`, если нет куки

Кидает `403`, если роль не разработчик или админ

## /clinic/list
- input: None
- curl example: `curl localhost:8083/clinic/list`
- output: Clinic array `clinicList`
- output example: `{"clinicList":[{"id":1,"name":"Клиника нейм","address":"ул. Пушкина, д. 42","phone":"88005553535"}]}`

## /clinic/find
Сделать поиск клиники по всем трём полям (адрес, имя, телефон)
- input: json
- input format: string `str`
- curl example: `curl localhost:8083/clinic/find -d '{"str":"пуш"}'`
- output: Clinic array `clinicList`

## /clinic/find/name
- input: json
- input format: string `name`
- curl example: `curl localhost:8083/clinic/find/name -d '{"name":"Клиника"}'`
- output: Clinic array `clinicList`

## /clinic/find/address
- input: json
- input format: string `address`
- curl example: `curl localhost:8083/clinic/find/address -d '{"address":"пушкина"}'`
- output: Clinic array `clinicList`

## /clinic/find/phone
- input: json
- input format: string `phone`
- curl example: `curl localhost:8083/clinic/find/phone -d '{"phone":"8800"}'`
- output: Clinic array `clinicList`

---

## /review/create
Создать отзыв
- input: json
- input format: int `cid` (clinic id), int `did` (doctor id from `users` table), int `sid` (service id), int `score` (from [1..5]), string `description`. All fields are optional.
- curl example: `curl localhost:8083/review/create -d '{"did":1, "score": 3, "description": "отзыв бла бла"}'`
- output: json with string `err` (see notes: empty json in response)

Кидает `400`, если айдишник не 0 и не существует или если `score` не от 1 до 5. (пока не кидает, но должно, смотри #32)

## /review/list
- input: None
- curl example: `curl localhost:8083/review/list`
- output: Review list `reviewList`
- output example: `{"reviewList":[{"id":1,"cid":1,"did":0,"sid":0,"score":3,"description":""},{"id":2,"cid":0,"did":1,"sid":0,"score":0,"description":"отзыв бла бла"}]}`

## /review/find/doctor
Найти все отзывы на доктора
- input: json
- input format: int `did`
- curl example: `curl localhost:8083/review/find/doctor -d '{"did":1}'`
- output: Review list `reviewList`

## /review/find/clinic
Найти все отзывы на клинику
- input: json
- input format: int `cid`
- curl example: `curl localhost:8083/review/find/clinic -d '{"cid":1}'`
- output: Review list `reviewList`

## /review/find/service
Найти все отзывы на услугу
- input: json
- input format: int `sid`
- curl example: `curl localhost:8083/review/find/service -d '{"sid":1}'`
- output: Review list `reviewList`

# Notes

### empty json in response
Frontend application struggles with empty jsons. Send json with field `err` and dummy response instead.
