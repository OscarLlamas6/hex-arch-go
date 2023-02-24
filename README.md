# Hexagonal Architecture in Go using Postgres and Firestore

Repositorio para practicar Arquitectura Hexagonal en Go


## Preparando emulador Firestore
```bash

# Running gcloud auth login
> gcloud auth application-default login

# Default json auth config loc
# Linux
> $HOME/.config/gcloud/application_default_credentials.json
# Windows
> %APPDATA%\gcloud\application_default_credentials.json

# Setting GOOGLE_APPLICATION_CREDENTIALS environment variable (powershell)
> $env:GOOGLE_APPLICATION_CREDENTIALS='C:\Users\oscar\AppData\Roaming\gcloud\application_default_credentials.json'

# Check if environment variable exists (powershell)
>  $env:GOOGLE_APPLICATION_CREDENTIALS
```

## Iniciar servicios 

```bash

> docker-compose -f ./docker-compose.yaml build

> docker-compose -f ./docker-compose.yaml down --remove-orphans

> docker-compose -f ./docker-compose.yaml up -d

```

## Start Air for hot reload (development)
```bash

# Install Air if needed
> go install github.com/cosmtrek/air@latest

# Start dev server
> air -c ./.air.toml

```

## CouchDB

```bash
# DB at
> http://localhost:5984/

# Admin panel at
> http://localhost:5984/_utils/


```