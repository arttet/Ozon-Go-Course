# Ozon Marketplace Project

![schema](images/schema.png)

Дальше везде используются **placeholder**-ы:
- `{domain}`,`{Domain}`
- `{subdomain}`,`{Subdomain}`

Например, для поддомена `package` из домена `logistic` значение **placeholder**-ов будет:
- `{domain}`,`{Domain}` = `logistic`,`Logistic`
- `{subdomain}`,`{Subdomain}` = `package`,`Package`
- `{domain}`/`{subdomain}` = `logistic`/`package`
---


### Задание 1

1. Сделать форк **ozonmp/omp-bot** репозитория в свой профиль
2. Запросить у своего тьютора свой домен/поддомен: **{domain}/{subdomain}**
3. Добавить в ветку `feature/task-1` своего форка поддержку следующих команд:
```
/help__{domain}__{subdomain} — print list of commands
/get__{domain}__{subdomain} — get a entity
/list__{domain}__{subdomain} — get a list of your entity (💎: with pagination via telegram keyboard)
/delete__{domain}__{subdomain} — delete an existing entity

/new__{domain}__{subdomain} — create a new entity // not implemented (💎: implement list fields via arguments)
/edit__{domain}__{subdomain} — edit a entity      // not implemented
```
4. Сделать PR из ветки `feature/task-1` своего форка в ветку `master` своего форка
5. Отправить ссылку на PR личным сообщением своему тьютору до конда дедлайна сдачи (см. таблицу прогресса)

#### Рецепт

Для добавления поддержки команд в рамках своего поддомена:

1. Написать структуру `{Subdomain}` с методом `String()`
2. Написать интерфейс `{Subdomain}Service` и **dummy** имплементацию
3. Написать интерфейс `{Subdomain}Commander` по обработке команд

---

2. Реализовать `{Subdomain}Service` в **internal/service/{domain}/{subdomain}/**

```go
package {subdomain}

import "github.com/ozonmp/omp-bot/internal/model/{domain}"

type {Subdomain}Service interface {
  Describe({subdomain}ID uint64) (*{domain}.{Subdomain}, error)
  List(cursor uint64, limit uint64) ([]{domain}.{Subdomain}, error)
  Create({domain}.{Subdomain}) (uint64, error)
  Update({subdomain}ID uint64, {subdomain} {domain}.{Subdomain}) error
  Remove({subdomain}ID uint64) (bool, error)
}

type Dummy{Subdomain}Service struct {}

func NewDummy{Subdomain}Service() *Dummy{Subdomain}Service {
  return &Dummy{Subdomain}Service{}
}

// ...
```

---

3. Реализовать `{Subdomain}Commander` по обработке команд в **internal/app/commands/{domain}/{subdomain}/**

```go
package {subdomain}

import (
  model "github.com/ozonmp/omp-bot/internal/model/{domain}"
  service "github.com/ozonmp/omp-bot/internal/service/{domain}/{subdomain}"
)

type {Subdomain}Commander interface {
  Help(inputMsg *tgbotapi.Message)
  Get(inputMsg *tgbotapi.Message)
  List(inputMsg *tgbotapi.Message)
  Delete(inputMsg *tgbotapi.Message)

  New(inputMsg *tgbotapi.Message)    // return error not implemented
  Edit(inputMsg *tgbotapi.Message)   // return error not implemented
}

func New{Subdomain}Commander(bot *tgbotapi.BotAPI, service service.{Subdomain}Service) {Subdomain}Commander {
  // ...
}
```
