### Скачать образ postgres
`docker pull postgres`

### Запустить образ
`docker run --name=wish-list -e POSTGRES_PASSWORD="postgres" -p 5432:5432 -d --rm postgres`

### Установить утилиту миграций, если не установлена
`https://github.com/golang-migrate/migrate/tree/master/cmd/migrate`

### Запустить миграции
`migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up`