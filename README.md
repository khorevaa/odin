# Odin. Remote Administration for 1S.Enterprise Application Servers

![Release](https://img.shields.io/github/release/khorevaa/odin.svg)
![Discord](https://img.shields.io/badge/discord-join%20channel-7289DA)](https://gofiber.io/discord)
![Test](https://github.com/khorevaa/odin/workflows/Test/badge.svg)
![Security](https://github.com/khorevaa/odin/workflows/Security/badge.svg)
![Linter](https://github.com/khorevaa/odin/workflows/Linter/badge.svg)

> ОДИН - мудрец и шаман, знаток рун и сказов, царь-жрец, колдун-воин,
> бог войны и победы, покровитель военной аристократии.

>ОДИН - это FULL REST API SERVER для администрирования и мониторинга серверов приложений 1С. Предприятие.

odin - приложение на golang, которое задействует c сервером платформы 1С по внутреннему протоколу. С его помощью можно прямо взаимодействовать с любым количеством серверов 1С, выполняя автоматизированно рутинные задачи администрирования, которые раньше требовали консоли Администрирования серверов 1С. Предприятие или rac.

Название odin образовалось из германо-скандинавской мифологии. И произносится как ферма Одина.


## Какие задачи решает odin?

Odin выполняет следующие задачи:

1. Администрирование информационных баз на серверах 1С Предприятие (создание, удаление и изменение).

2. Администрирование рабочих серверов в кластере 1С. Предприятие

3. Получение информации по сеансам и соединениям, в том числе их отключение.

4. Получение информации по лицензиям используемыми

5. Получение состояние работоспособности кластера серверов 1С. Предприятие

5. Другие приятные функции позволяющие автоматизировать работы.

В целом приложение по функциональности полностью повторяет стандартное приложение rac от 1С.

## Как запустить?

Для работы требуется запустить приложение с необходимыми ключами работы, предварительно положив в папку с приложением файл временной лицензии (или указав его через специальную опцию):

```shell

Usage: odin [OPTIONS]
Remote Administration for 1S.Enterprise Application Servers

Options:
-v, --version   Show the version and exit
--debug    Enable debug
--prefork   Enable prefork in Production
-p, --port      Port to listen on (default "localhost:3001")
--lcn         License file (default "license.lic")
--dir         App data dir
-c, --config   File to load preconfigure app servers
```

Адрес доступного API http://localhost:3001/api/v1

Полное описание возможностей API можно найти по адресу http://localhost:3001/docs

Полное описание API https://app.swaggerhub.com/apis/khorevaa/odin-remote_administration_1_s_enterprise_server/1.0

## Как настроить и использовать?

Для начала надо запустить службу ras на сервере 1С. Предприятие.

После добавить адрес сервера 1С. Предприятие в приложение используя соответствующую команды API

#### Добавление сервера 1С. Предприятие

```shell
curl -i \
    -H "Accept: application/json" \
    -X POST -d "addr":"srv-app","descr":"ББУ 30","name":"bbu30"\
    http://localhost:3001/api/v1/app
```
Подробное описание модели добавления смотрите в http://localhost:3001/docs/index.html#/app/post_app

Проверка добавленного сервера (получение списка всех подключенных серверов 1С. Предприятие):
```shell
curl -i \
    -H "Accept: application/json" \
    -X GET http://localhost:3001/api/v1/app
```
Ответ:
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "name": "bbu30",
      "addr": "srv-app",
      "port": "1545",
      "version": "9.0",
      "descr": "ББУ 30",
      "agent_version": "8.3.16.1359"
    }]
}
```

### Примеры возможностей приложения

#### Health проверка
Запрос
```shell
curl -i \
    -H "Accept: application/json" \
    -X GET http://localhost:3001/api/v1/health/readiness
```
Ответ
```json
{
  "name": "API Remote Administration for 1S.Enterprise Application Servers",
  "version": "1.0",
  "status": true,
  "apps": [
    {
      "name": "bbu30",
      "host": "srv-app:1545",
      "status": true,
      "response_time": 0,
      "url": "http://localhost:3001/api/v1/app/bbu30/health"
    }
  ]
}
```

#### Добавление новой информационной базы
Запрос
```shell
curl -X POST "http://localhost:3001/api/v1/app/bbu30/infobases" -H "accept: application/json" \
     -H "Content-Type: application/json" \ 
     -d "{ \"date_offset\": 0, \"db_name\": \"base\", \"db_pwd\": \"password\", \"db_server\": \"sql\", \"db_user\": \"user\", \"dbms\": \"MSSQLServer\", \"denied_from\": \"2020-10-01T08:30:00Z\", \"denied_message\": \"Выполняется обновление базы\", \"denied_parameter\": \"123\", \"denied_to\": \"2020-10-01T08:30:00Z\", \"descr\": \"Это очень хорошая база\", \"external_session_manager_connection_string\": \"http://auth2.com\", \"external_session_manager_required\": false, \"license_distribution\": 0, \"locale\": \"ru_RU\", \"name\": \"test\", \"permission_code\": \"123\", \"reserve_working_processes\": false, \"safe_mode_security_profile_name\": \"profile1\", \"scheduled_jobs_deny\": false, \"security_level\": 0, \"security_profile_name\": \"sec_profile1\", \"sessions_deny\": false}"
```
Ответ
```json
{
  "code": 0,
  "data": {
    "cluster_id": "efa3672f-947a-4d84-bd58-b21997b83561",
    "date_offset": 0,
    "db_name": "base",
    "db_pwd": "password",
    "db_server": "sql",
    "db_user": "user",
    "dbms": "MSSQLServer",
    "denied_from": "2020-10-01T08:30:00Z",
    "denied_message": "Выполняется обновление базы",
    "denied_parameter": "123",
    "denied_to": "2020-10-01T08:30:00Z",
    "descr": "Это очень хорошая база",
    "external_session_manager_connection_string": "http://auth2.com",
    "external_session_manager_required": false,
    "license_distribution": 0,
    "locale": "ru_RU",
    "name": "bbu_org",
    "permission_code": "123",
    "reserve_working_processes": false,
    "safe_mode_security_profile_name": "profile1",
    "scheduled_jobs_deny": false,
    "security_level": 0,
    "security_profile_name": "sec_profile1",
    "sessions_deny": false,
    "uuid": "efa3672f-947a-4d84-bd58-b21997b83561"
  }
}
```
#### Получение списка информационных базы
Запрос
```shell
curl -X GET "http://localhost:3001/api/v1/app/bbu30/infobases" -H "accept: application/json"
```
Ответ
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "cluster_id": "a7e54a28-13a5-4fad-9abd-7c2aecc46a01",
      "uuid": "2c0d42f3-3ae2-4d98-99a7-b30ee93449b5",
      "name": "bbu_org",
      "descr": "ББУ ООО Рога и копыта"
    }]
}
```
#### Получение списка соединений
Запрос
```shell
curl -X GET "http://localhost:3001/api/v1/app/bbu30/connections" -H "accept: application/json"
```
Ответ
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "uuid": "c0c3e30b-1c96-4ad2-8995-61fdc256c72b",
      "conn_id": 0,
      "host": "srv-app",
      "process": "af442afa-2c8f-4eea-81ee-31384c2b01aa",
      "cluster_id": "a7e54a28-13a5-4fad-9abd-7c2aecc46a01",
      "infobase_id": "00000000-0000-0000-0000-000000000000",
      "application": "RAS",
      "connected_at": "2021-02-08T16:22:53+03:00",
      "session_id": 0,
      "blocked_by_ls": 0
    },
    {
      "uuid": "db8e27b1-a0ed-4283-93e3-562ff2571515",
      "conn_id": 1440207,
      "host": "term1c",
      "process": "af442afa-2c8f-4eea-81ee-31384c2b01aa",
      "cluster_id": "a7e54a28-13a5-4fad-9abd-7c2aecc46a01",
      "infobase_id": "7a360e0e-79b9-4803-9794-0677af6eb20f",
      "application": "1CV8C",
      "connected_at": "2021-02-08T13:23:41+03:00",
      "session_id": 0,
      "blocked_by_ls": 0
    }
  ]
}
```
#### Получение списка сеансов
Запрос
```shell
curl -X GET "http://localhost:3001/api/v1/app/bbu30/sessions" -H "accept: application/json"
```
Ответ
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "uuid": "1b50d66b-f46c-4513-9379-45b03df74d74",
      "id": 986,
      "infobase_id": "7a360e0e-79b9-4803-9794-0677af6eb20f",
      "connection_id": "00000000-0000-0000-0000-000000000000",
      "process_id": "00000000-0000-0000-0000-000000000000",
      "user_name": "ФИО_4",
      "host": "",
      "app_id": "WebClient",
      "locale": "ru_RU",
      "started_at": "2021-02-08T14:23:14+03:00",
      "last_active_at": "2021-02-08T14:43:31+03:00",
      "hibernate": true,
      "passive_session_hibernate_time": 1200,
      "hibernate_dession_terminate_time": 86400,
      "blocked_by_dbms": 0,
      "blocked_by_ls": 0,
      "bytes_all": 397402,
      "bytes_last_5_min": 397402,
      "calls_all": 185,
      "calls_last_5_min": 185,
      "dbms_bytes_all": 6049699,
      "dbms_bytes_last_5_min": 6049699,
      "db_proc_info": "",
      "db_proc_took": 0,
      "db_proc_took_at": "0001-01-01T00:00:00Z",
      "duration_all": 13030,
      "duration_all_dbms": 4613,
      "duration_current": 0,
      "duration_current_dbms": 0,
      "duration_last_5_min": 13030,
      "duration_last_5_min_dbms": 4613,
      "memory_current": 0,
      "memory_last_5_min": 35252740,
      "memory_total": 35252740,
      "read_current": 0,
      "read_last_5_min": 4630914,
      "read_total": 4630914,
      "write_current": 0,
      "write_last_5_min": 4237863,
      "write_total": 4237863,
      "duration_current_service": 0,
      "duration_last_5_min_service": 0,
      "duration_all_service": 0,
      "current_service_name": "",
      "cpu_time_current": 0,
      "cpu_time_last_5_min": 0,
      "cpu_time_total": 0,
      "data_separation": "",
      "client_ip_address": "",
      "licenses": null,
      "cluster_id": "a7e54a28-13a5-4fad-9abd-7c2aecc46a01"
    },
    {
      "uuid": "3b9d538d-742a-4f93-a4a9-7657e0f136f5",
      "id": 987,
      "infobase_id": "7a360e0e-79b9-4803-9794-0677af6eb20f",
      "connection_id": "d229aab6-eed9-4790-89e4-b972baf9d25a",
      "process_id": "af442afa-2c8f-4eea-81ee-31384c2b01aa",
      "user_name": "Администратор",
      "host": "SRV-APP",
      "app_id": "BackgroundJob",
      "locale": "ru",
      "started_at": "2021-02-08T16:24:29+03:00",
      "last_active_at": "2021-02-08T16:29:43+03:00",
      "hibernate": false,
      "passive_session_hibernate_time": 0,
      "hibernate_dession_terminate_time": 0,
      "blocked_by_dbms": 0,
      "blocked_by_ls": 0,
      "bytes_all": 0,
      "bytes_last_5_min": 0,
      "calls_all": 0,
      "calls_last_5_min": 0,
      "dbms_bytes_all": 4300078,
      "dbms_bytes_last_5_min": 0,
      "db_proc_info": "",
      "db_proc_took": 0,
      "db_proc_took_at": "0001-01-01T00:00:00Z",
      "duration_all": 0,
      "duration_all_dbms": 561,
      "duration_current": 317664,
      "duration_current_dbms": 0,
      "duration_last_5_min": 0,
      "duration_last_5_min_dbms": 0,
      "memory_current": 3880244,
      "memory_last_5_min": 0,
      "memory_total": 0,
      "read_current": 0,
      "read_last_5_min": 0,
      "read_total": 0,
      "write_current": 1124,
      "write_last_5_min": 0,
      "write_total": 0,
      "duration_current_service": 315278,
      "duration_last_5_min_service": 0,
      "duration_all_service": 1450,
      "current_service_name": "ExternalDataSourceXMLAService",
      "cpu_time_current": 0,
      "cpu_time_last_5_min": 0,
      "cpu_time_total": 0,
      "data_separation": "",
      "client_ip_address": "",
      "licenses": null,
      "cluster_id": "a7e54a28-13a5-4fad-9abd-7c2aecc46a01"
    },
    {
      "uuid": "95070ba8-8727-4286-9004-00379010dd1a",
      "id": 970,
      "infobase_id": "7a360e0e-79b9-4803-9794-0677af6eb20f",
      "connection_id": "00000000-0000-0000-0000-000000000000",
      "process_id": "00000000-0000-0000-0000-000000000000",
      "user_name": "ФИО_3",
      "host": "",
      "app_id": "1CV8C",
      "locale": "ru_RU",
      "started_at": "2021-02-08T13:04:09+03:00",
      "last_active_at": "2021-02-08T16:29:46+03:00",
      "hibernate": false,
      "passive_session_hibernate_time": 1200,
      "hibernate_dession_terminate_time": 86400,
      "blocked_by_dbms": 0,
      "blocked_by_ls": 0,
      "bytes_all": 6475075,
      "bytes_last_5_min": 222091,
      "calls_all": 2409,
      "calls_last_5_min": 70,
      "dbms_bytes_all": 11713901,
      "dbms_bytes_last_5_min": 347706,
      "db_proc_info": "",
      "db_proc_took": 0,
      "db_proc_took_at": "0001-01-01T00:00:00Z",
      "duration_all": 274656,
      "duration_all_dbms": 105209,
      "duration_current": 0,
      "duration_current_dbms": 0,
      "duration_last_5_min": 14733,
      "duration_last_5_min_dbms": 8546,
      "memory_current": 0,
      "memory_last_5_min": 19998075,
      "memory_total": 236560645,
      "read_current": 0,
      "read_last_5_min": 468974,
      "read_total": 12552059,
      "write_current": 0,
      "write_last_5_min": 468974,
      "write_total": 12416671,
      "duration_current_service": 0,
      "duration_last_5_min_service": 501,
      "duration_all_service": 52298,
      "current_service_name": "",
      "cpu_time_current": 0,
      "cpu_time_last_5_min": 0,
      "cpu_time_total": 0,
      "data_separation": "",
      "client_ip_address": "",
      "licenses": [
        {
          "process_id": "00000000-0000-0000-0000-000000000000",
          "session_id": "95070ba8-8727-4286-9004-00379010dd1a",
          "user_name": "ФИО_3",
          "host": "",
          "app_id": "1CV8C",
          "full_name": "",
          "series": "ORG8B",
          "issued_by_server": true,
          "license_type": 1,
          "net": true,
          "max_users_all": 500,
          "max_users_cur": 500,
          "rmngr_address": "SRV-APP",
          "rmngr_port": 1541,
          "rmngr_pid": "3192",
          "short_presentation": "Сервер, ORG8B Сет 500",
          "full_presentation": "Сервер, 3192, SRV-APP, 1541, ORG8B Сетевой 500"
        }
      ],
      "cluster_id": "a7e54a28-13a5-4fad-9abd-7c2aecc46a01"
    },
    {
      "uuid": "b1e6b6aa-2e87-4bd7-b42a-dc362ac62a98",
      "id": 948,
      "infobase_id": "7a360e0e-79b9-4803-9794-0677af6eb20f",
      "connection_id": "00000000-0000-0000-0000-000000000000",
      "process_id": "00000000-0000-0000-0000-000000000000",
      "user_name": "ФИО_2",
      "host": "",
      "app_id": "1CV8C",
      "locale": "ru",
      "started_at": "2021-02-08T11:39:29+03:00",
      "last_active_at": "2021-02-08T16:28:57+03:00",
      "hibernate": false,
      "passive_session_hibernate_time": 1200,
      "hibernate_dession_terminate_time": 86400,
      "blocked_by_dbms": 0,
      "blocked_by_ls": 0,
      "bytes_all": 1842619,
      "bytes_last_5_min": 22052,
      "calls_all": 1256,
      "calls_last_5_min": 17,
      "dbms_bytes_all": 23340989,
      "dbms_bytes_last_5_min": 381562,
      "db_proc_info": "",
      "db_proc_took": 0,
      "db_proc_took_at": "0001-01-01T00:00:00Z",
      "duration_all": 401505,
      "duration_all_dbms": 303120,
      "duration_current": 0,
      "duration_current_dbms": 0,
      "duration_last_5_min": 7677,
      "duration_last_5_min_dbms": 5927,
      "memory_current": 0,
      "memory_last_5_min": 856611,
      "memory_total": 123673318,
      "read_current": 0,
      "read_last_5_min": 876,
      "read_total": 6554568,
      "write_current": 0,
      "write_last_5_min": 0,
      "write_total": 6900453,
      "duration_current_service": 0,
      "duration_last_5_min_service": 0,
      "duration_all_service": 642,
      "current_service_name": "",
      "cpu_time_current": 0,
      "cpu_time_last_5_min": 0,
      "cpu_time_total": 0,
      "data_separation": "",
      "client_ip_address": "",
      "licenses": [
        {
          "process_id": "00000000-0000-0000-0000-000000000000",
          "session_id": "b1e6b6aa-2e87-4bd7-b42a-dc362ac62a98",
          "user_name": "ФИО_2",
          "host": "",
          "app_id": "1CV8C",
          "full_name": "",
          "series": "ORG8A",
          "issued_by_server": false,
          "license_type": 1,
          "net": true,
          "max_users_all": 300,
          "max_users_cur": 300,
          "rmngr_address": "",
          "rmngr_port": 0,
          "rmngr_pid": "20896",
          "short_presentation": "Клиент, ORG8A Сет 300",
          "full_presentation": "Клиент, 20896, ORG8A Сетевой 300"
        }
      ],
      "cluster_id": "a7e54a28-13a5-4fad-9abd-7c2aecc46a01"
    },
    {
      "uuid": "d5a55fec-d842-42b7-9feb-b4e3fff264b1",
      "id": 968,
      "infobase_id": "7a360e0e-79b9-4803-9794-0677af6eb20f",
      "connection_id": "00000000-0000-0000-0000-000000000000",
      "process_id": "00000000-0000-0000-0000-000000000000",
      "user_name": "ФИО_1",
      "host": "",
      "app_id": "1CV8C",
      "locale": "ru",
      "started_at": "2021-02-08T12:58:41+03:00",
      "last_active_at": "2021-02-08T16:29:19+03:00",
      "hibernate": false,
      "passive_session_hibernate_time": 1200,
      "hibernate_dession_terminate_time": 86400,
      "blocked_by_dbms": 0,
      "blocked_by_ls": 0,
      "bytes_all": 14037345,
      "bytes_last_5_min": 35078,
      "calls_all": 2009,
      "calls_last_5_min": 30,
      "dbms_bytes_all": 16093457,
      "dbms_bytes_last_5_min": 117605,
      "db_proc_info": "",
      "db_proc_took": 0,
      "db_proc_took_at": "0001-01-01T00:00:00Z",
      "duration_all": 234179,
      "duration_all_dbms": 72389,
      "duration_current": 0,
      "duration_current_dbms": 0,
      "duration_last_5_min": 3152,
      "duration_last_5_min_dbms": 483,
      "memory_current": 0,
      "memory_last_5_min": 568135,
      "memory_total": 172248735,
      "read_current": 0,
      "read_last_5_min": 7002,
      "read_total": 24506665,
      "write_current": 0,
      "write_last_5_min": 24405,
      "write_total": 24673237,
      "duration_current_service": 0,
      "duration_last_5_min_service": 16,
      "duration_all_service": 2626,
      "current_service_name": "",
      "cpu_time_current": 0,
      "cpu_time_last_5_min": 0,
      "cpu_time_total": 0,
      "data_separation": "",
      "client_ip_address": "",
      "licenses": [
        {
          "process_id": "00000000-0000-0000-0000-000000000000",
          "session_id": "d5a55fec-d842-42b7-9feb-b4e3fff264b1",
          "user_name": "ФИО_1",
          "host": "",
          "app_id": "1CV8C",
          "full_name": "",
          "series": "ORGL8",
          "issued_by_server": false,
          "license_type": 1,
          "net": false,
          "max_users_all": 10,
          "max_users_cur": 10,
          "rmngr_address": "",
          "rmngr_port": 0,
          "rmngr_pid": "3484",
          "short_presentation": "Клиент, ORGL8 Лок 10",
          "full_presentation": "Клиент, 3484, ORGL8 Локальный 10"
        }
      ],
      "cluster_id": "a7e54a28-13a5-4fad-9abd-7c2aecc46a01"
    }
  ]
}
```

И другие команды подробно описанные в документации  http://localhost:3001/docs

## Лицензирование и использование

На текущий момент это анонс и предоставление демо версии ПО с ограниченями.

#### Ограничения демо версии


## Сотрудничество

Рассмотрю сотрудничество по данному ПО. (продажа, установка и автоматизация рутинных процедур)

По всем вопросам пишите в ЛС.
