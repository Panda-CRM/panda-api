# PANDA-API

[![Build Status](https://travis-ci.org/wilsontamarozzi/panda-api.svg?branch=master)](https://travis-ci.org/wilsontamarozzi/panda-api)

## ENDPOINTS

Todos os pedidos são servidos através de HTTPS. A versão atual é v1.
* [Pessoa](full_format.md#pessoa)
* [Tarefa](full_format.md#tarefa)
* [Categoria de Tarefa](full_format.md#categoria-de-tarefa)
* [Histórico de Tarefa](full_format.md#histórico-de-tarefa)

## OBJETOS

Esses são todos os objetos oferecidos pela API no formato JSON.
* [Pessoa](full_format.md#pessoa)
* [Tarefa](full_format.md#tarefa)
* [Categoria de Tarefa](full_format.md#categoria-de-tarefa)
* [Histórico de Tarefa](full_format.md#histórico-de-tarefa)

## CONFIGURAÇÕES

### Banco de Dados
Nesse projeto é utilizado o ORM [Gorm](http://jinzhu.me/gorm/) para fazer todos os relacionamentos do banco de dados.
Por padrão o sistema está configurado para fazer `AutoMigrate` de todas as entidades.
Para realizar a configuração de **usuário**, **senha**, **host** e **nome do banco**; sete as variaveis ambientes.

```golang
const(
    ENV_DB_DRIVER = "DB_DRIVER"
    ENV_DB_HOST = "DB_HOST"
    ENV_DB_NAME = "DB_NAME"
    ENV_DB_USER = "DB_USER"
    ENV_DB_PASSWORD = "DB_PASSWORD"
    ENV_DB_SSL_MODE = "DB_SSL_MODE"
    ENV_DB_MAX_CONNECTION = "DB_MAX_CONNECTION"
    ENV_DB_LOG_MODE = "DB_LOG_MODE"
)
```

Caso o sistema não encontre as variáveis ambientes, ele irá rodar localmente com essas configurações:

```golang
var(
    DB_DRIVER string = "postgres"
    DB_HOST string = "localhost"
    DB_NAME string = "panda"
    DB_USER string = "pandaapi"
    DB_PASSWORD string = "1234"
    DB_SSL_MODE string = "disable" // disable | require
    DB_MAX_CONNECTION int = 1
    DB_LOG_MODE bool = true
)
```

### JWT Token Auth
Para fazer qualquer alteração no `SECRET_KEY` do token, sete a variável ambiente abaixo:

```golang
const ENV_JWT_SECRET_KEY = "JWT_SECRET_KEY" // default = panda
```

### Porta API
A porta da API está setada para `8080`. Para realizar alteração na porta, sete a seguinte variável ambiente.

```golang
const ENV_RUN_PORT = "PORT" // default = 8080
```

### Integração de Logs com Bugsnag
Sistema é integrado com o Bugsnag para analise de logs e problemas. Para ativar a integração, basta configurar a seguinte variável ambiente:

```golang
const ENV_API_KEY_BUGSNAG = "API_KEY_BUGSNAG"
```