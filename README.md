## INSTALAÇÃO

Antes de tudo é necessário importar as dependencias do projeto.
```
$ go get
```

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
Para realizar a configuração de **usuário**, **senha**, **host** e **nome do banco**; vá até o arquivo `database.go`.
```
Caminho: panda-api/database/database.go
```
```golang
const DB_HOST = "localhost"
const DB_NAME = "panda"
const DB_USER = "pandaapi"
const DB_PASSWORD = "1234"
const DB_MAX_CONNECTION = 1
const DB_LOG_MODE = true
```

### JWT Token Auth
Para fazer qualquer alteração no `SECRET_KEY` do token, vá até o arquivo `security.go`.
```
Caminho: panda-api/security/security.go
```
```golang
const SECRET_KEY = "secret_key"
```

### Porta API
Por padrão o sistema tenta buscar a variavel ambiente na maquina pelo nome `PORT`, caso ele não encontre ele irá configurar a porta padrão `8080`.
Para realizar alteração na porta, vá até o arquivo `server.go`.
```
Caminho: panda-api/server.go
```