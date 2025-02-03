# liquor

<img align="right" width="159px" src="https://avatars.githubusercontent.com/u/197004919">

Liquor is a web development framework built with Golang, designed to simplify the implementation of web projects. It is primarily focused on backend development and leverages several libraries specifically tailored for web development.


- [Installation](#install-cli)
- [Usage](#usage)
    - [Create a new app](#create-a-new-app)
- [Features](#features)
- [Modules](#modules)

## Install CLI

```bash
go install github.com/go-liquor/liquor@latest
```

## Usage

### Create a new app

```
liquor app create --name <APP_NAME> --pkg <PACKAGE_NAME>
```

## Features

- Application Modular (with https://github.com/go-uber/fx)
- Config file
- Gin Framework implementation
- CORS
- Database connection
    - Sqlite
    - MySQL
    - Postgres
    - MongoDB
- Logger (with https://github.com/go-uber/zap)


## Modules

- [database/mongodb](sdk/modules/database/mongodb/README.md)
- [database/mysql](sdk/modules/database/mysql/README.md)
- [database/postgres](sdk/modules/database/postgres/README.md)
- [database/sqlite](sdk/modules/database/sqlite/README.md)

## Docs

https://go-liquor.github.io
