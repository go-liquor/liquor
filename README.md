# liquor

<img align="right" width="159px" src="https://avatars.githubusercontent.com/u/197004919">

Liquor is a web development framework built with Golang, designed to simplify the implementation of web projects. It is primarily focused on backend development and leverages several libraries specifically tailored for web development.


- [Docs](https://go-liquor.github.io)
- [Installation](#install-cli)
- [Usage](#usage)
    - [Create a new app](#create-a-new-app)
    - [Create a new migration](#create-a-new-migration)
    - [Create a new service](#create-a-new-service)
    - [Create a new api](#create-a-new-api)

## Install CLI

```bash
go install github.com/go-liquor/liquor/v2@latest
```

## Usage

### Create a new app

```bash
liquor create app --name <APP_NAME> --pkg <PACKAGE_NAME>
```

### Create a new migration
```bash
liquor create migration --name <MIGRATION_NAME>
```

### Create a new service

```bash
liquor create service --name <SERVICE_NAME>
```

### Create a new api

```bash
liquor create api --name <API_NAME>
```