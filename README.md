# MyGram - Final Project Golang (Digitalent FGA)

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Contributing](../CONTRIBUTING.md)

## About <a name = "about"></a>

MyGram, an app to save, post, and comment photos of other people. MyGram is built using Go with [Gin Web Framework](https://gin-gonic.com/), PostgreSQL for database, and [GORM](https://gorm.io/) for ORM. This project is created as a final project for Digital Talent Kominfo FGA "Scalable Web Service with Golang" Program by Hacktiv8.

## Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- What things you need to install

  1. [go](https://go.dev/), kindly choose go version that compatible with your Operating System (OS), latest version 1.19.1
  2. postgresql, set up your database first before running this project

- Set the environment variable in `.env` files

  ```
  MY_GRAM_POSTGRES_HOST=
  MY_GRAM_POSTGRES_PORT=
  MY_GRAM_POSTGRES_DATABASE=
  MY_GRAM_POSTGRES_USERNAME=
  MY_GRAM_POSTGRES_PASSWORD=

  MY_GRAM_SECRET_JWT_SIGNATURE=
  ```

## Usage <a name = "usage"></a>

Before you run this project, you must run these commands below

```golang
go mod download && go mod tidy
```

To test this project, you need to have Go in your local and run command below

```golang
go run application.go
```

### Endpoint List

TBC
