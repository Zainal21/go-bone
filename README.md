# Go Bone Boilerplate

## Getting started

This is built on top of [Go Fiber](https://docs.gofiber.io) Golang Framework.

## Dependencies

There is some dependencies that we used in this skeleton:

- [Go Fiber](https://docs.gofiber.io/) [Go Framework]
- [Viper](https://github.com/spf13/viper) [Go Configuration]
- [Cobra](https://github.com/spf13/cobra) [Go Modern CLI]
- [Logrus Logger](https://github.com/sirupsen/logrus) [Go Logger]
- [Goose Migration](https://github.com/pressly/goose) [Go Migration]
- [Gobreaker](https://github.com/sony/gobreaker) [Go Circuit Breaker]

## Requirement

- Golang version 1.21 or latest
- Database MySQL

## Usage

### Installation

install required dependencies

```bash
make install
```

### Run Service

run current service after all dependencies installed

```bash
make start
```

## Database Migration

migration up

```bash
go run main.go db:migrate up
```

migration down

```bash
go run main.go db:migrate down
```

migration reset

```bash
go run main.go db:migrate reset
```

migration reset

```bash
go run main.go db:migrate reset
```

migration redo

```bash
go run main.go db:migrate redo
```

migration status

```bash
go run main.go db:migrate status
```

create migration table

```bash
go run main.go db:migrate create {table-name} sql

# example
go run main.go db:migrate create users sql
```

to show all command

```bash
go run main.go db:migrate
```

run seeder

```bash
go run main.go db:seed

# example
go run main.go db:seed

# example spesific function
go run main.go db:seed {func-seeder-name}
```
