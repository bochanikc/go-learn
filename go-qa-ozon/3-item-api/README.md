# Item api

### api

Контракт сервиса

### cmd

Входная точка в приложение

### internal

Внутренние файлы сервиса

### migrations

Миграции Для выполнения миграции выполните `make migration`

### pkg

Клиенты и структуры для сервера и клиента, файл swagger. Для генерации из proto выполните `make generate`

### test

Интеграционные тесты


# Как запустить интеграционные тесты?

1. `make version` - проверить установку всего нужного
2. `make database` - запустить докер с локальной бд
3. `make migrate` - сделать миграции бд
4. `make run` - запустить сервис
5. `make integration-test` - запустить интеграционные тесты для поднятого сервиса