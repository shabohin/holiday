# Holiday database

## Что бы запустить локально

`docker compose up`

При этом запускается баз данных и админка pgAdmin

В файле docker-compose.yml в разделе environment лежат

-   имя базы (POSTGRES_DB)
-   пользователь (POSTGRES_USER)
-   пароль (POSTGRES_PASSWORD)

После запуска, появится папка pgdata, в ней хранится база данных между запусками.
Если нужно очистить базу до начального состояния, просто удали эту папку.
Так же добавится папка pgadmin которую можешь не трогать.

## Что бы завести админку

-   Перейди на http://localhost:5050/browser/
-   Введи пароль adminpwd (PGADMIN_DEFAULT_PASSWORD)
-   Нажми Add New Server
-   Введи любое имя в Name
-   Перейди во вторую вкладку Connection
-   Host name/address: postgres_container (postgres.container_name)
-   Maintenance database: holidaydb (POSTGRES_DB)
-   Username: holidayusr (POSTGRES_USER)
-   Password: holidaypwd (POSTGRES_PASSWORD)
-   Не забудь сохранить паспорт

## Иницилизация базы данных

В папке init хранятся скрипты, которые зальются при старте контейнера
