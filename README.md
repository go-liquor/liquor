# liquor

<img align="right" width="159px" src="https://avatars.githubusercontent.com/u/197004919">

Liquor is a web development framework built with Golang, designed to simplify the implementation of web projects. It is primarily focused on backend development and leverages several libraries specifically tailored for web development.


- [Docs](https://go-liquor.github.io)
- [Installation](#install-cli)
- [Usage](#usage)
    - [Create a new app](#create-a-new-app)
    - [Run application](#run-application)
    - [Enable module](#enable-module)
    - [Create a new route group](#create-a-new-route-group)
    - [Create a new entity](#create-a-new-entity)
    - [Create a new service](#create-a-new-service)


## Install CLI

```bash
go install github.com/go-liquor/liquor@latest
```

## Usage

### Create a new app

```bash
liquor app create --name <APP_NAME> --pkg <PACKAGE_NAME>
```

### Run application

```bash
liquor run
```

### Enable module

```bash
liquor app enable <MODULE_NAME>
```


### Create a new route group


```bash
liquor create route --name <GROUP_NAME> --group /api/<GROUP_NAME> <--crud>
```

### Create a new entity

```bash
liquor create entity --name <ENTITY_NAME>
```

### Create a new service

```bash
liquor create service --name <SERVICE_NAME>
```

### Create a new repository

```bash
liquor create repository --name <REPOSITORY_NAME>
```
