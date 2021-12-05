# week-5-workshop/category-service

Это максимально легковесный шаблон для gRPC.

## Запуск

Для запуска необходимо предварительно выполнить `docker-compose up` из папки docker

```sh
make run
```

## Прогон миграций

```sh
make migrate
```

## Добавление сущностей

```sh
pgcli -h localhost -p 4432 -U user -d db
user@localhost:db> INSERT INTO category (name, created_at) VALUES ('Auto', NOW()), ('Electronics', NOW()), ('Toys', NOW());
```
