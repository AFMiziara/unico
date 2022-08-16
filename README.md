# Project Name
> Unico

## Indíce
- [Project Name](#project-name)
  - [Indíce](#indíce)
  - [Informação](#informação)
  - [Tecnologias](#tecnologias)
  - [Instalação](#instalação)
    - [Considerações](#considerações)
  - [Ambiente](#ambiente)
  - [Arquitetura de pastas](#arquitetura-de-pastas)
    - [Diretórios](#diretórios)
  - [Iniciando](#iniciando)
  - [Testes](#testes)

## Informação
Projeto de teste da entrevista 

## Tecnologias
* [GoLang](https://golang.org/) - Compilador da linguagem Go
* [Go Mod](https://github.com/golang/mod) - Gerenciador de dependencias
* [FIBER](https://gofiber.io/) - Web Framework Go
* [GORM](https://gorm.io/gorm) - Framework ORM Go

## Instalação
Clonando o projeto
``` bash
cd $PROJECT_HOME
git clone https://github.com/fsvxavier/unico.git
```
Instalando dependências
```
$ go mod init
```
Removendo dependencias indesejadas
``` bash
$ go mod tidy
```
Baixando as dependencias para a vendor local
``` bash
$ go mod vendor
```

Iniciando o Banco de dados postgresql replicado pelo Docker (Já com carga inicial de dados)
``` bash
$ sudo docker-compose -f docker/postgresql_master_slave/docker-compose.yaml up --build --abort-on-container-exit
```

### Considerações
``` bash
#
# Ambiente testado foi Ubuntu 20.04 LTS 64 Bits
# Certifique-se que efetuou a instalação correta do go
# Certifique-se que você tenha um compilador C/C++ instalado e acessível globalmente
# IDEs recomendadas: visual studio code
# Existem algumas funções Make no projeto para facilitar algumas ações
# Os exemplos de consumo da API estão disponíveis no Swagger da aplicação
#

```

## Ambiente
Configurando as variáveis de ambiente

| Nome | Descrição | Valor Padrão | Obrigatório |
| -- | -- | -- | -- |
| PORT | Porta padrão que a API irá subir | 5000 | :white_check_mark: |
| HTTP_READ_TIMEOUT | Timeout de leitura das para o httpserver | 60 | :white_check_mark: |
| HTTP_WRITE_TIMEOUT | Timeout de escrita das para o httpserver | 60 | :white_check_mark: |
| DB_HOST | Host da base de dados | | :white_check_mark: |
| DB_NAME | Nome da base de dados | | :white_check_mark: |
| DB_PORT | Porta da base de dados | | :white_check_mark: |
| DB_USER  | Usuário de aplicação da base de dados | | :white_check_mark: |
| DB_PASSWORD  | Senha do usuário da aplicação da base de dados | | :white_check_mark: |
| DB_SL_HOST | Host da base de dados replica | | :white_check_mark: |
| DB_SL_NAME | Nome da base de dados replica | | :white_check_mark: |
| DB_SL_PORT | Porta da base de dados replica | | :white_check_mark: |
| DB_SL_USER  | Usuário de aplicação da base de dados replica | | :white_check_mark: |
| DB_SL_PASSWORD  | Senha do usuário da aplicação da base de dados replica | | :white_check_mark: |
| EXECUTE_MIGRATION  | Ativar ou desativar o migrations (criação ou atualização de tabelas baseado nas models) | | :white_check_mark: |
| APP | Nome do app | dock | :white_check_mark: |
| LOGRUS_LOG_LEVEL | Nível de severidade do log a ser impresso | INFO | :white_check_mark: |
| VERSION_APP | versão da aplicação | INFO | :white_check_mark: |
| VERSION_API | versão da API | INFO | :white_check_mark: |

## Arquitetura de pastas
### Diretórios
```bash
unico
       |-- api
       |-- config
       |-- database
       |-- docker
           |-- postgresql_master_slave
       |-- docs
       |-- internal
           |-- interfaces
           |-- models
           |-- repositories
           |-- routes
           |-- usecases
       |-- logs
       |-- migrations
       |-- pkg
       |-- utils
       |-- .gitignore
       |-- README.md
```

## Iniciando
Buildando o projeto
``` bash
# execute o comando abaixo para buildar a aplicação e garantir que nada está quebrado
$ go build
```
Executando o projeto
``` bash
$ go run main.app or ./dock
```
## Testes
```bash
# Para execução dos testes automatizados executar o comando abaixo no terminal dentro da pasta da aplicação
$ go test -v -cover ./...

# Para gerar a interface mostrando todos os arquivos e as linhas "Covered", "Not Covered" e "Not Tracked":
$ go test ./... -coverprofile cobertura/fmtcoverage.html fmt
$ go test ./... -coverprofile cobertura/cover.out
$ go tool cover -html=cobertura/cover.out # em ambiente windows tirar o = após -html
$ go tool cover -html=cobertura/cover.out -o cobertura/cover.html # em ambiente windows tirar o = após -html
$ open 'cobertura/cover.html' file # em ambiente windows abrir externamente
```
## Gerando Swagger para commit
```bash
# Caso tenha alteração nas definições do swagger é necessário executar o comando abaixo assim alterando a pasta /docs e realizar o commit da mesma
$ swag init 

# caso seja exibido erro ao de comando não encontrado para o swag executar os comandos abaixo
$ export GOPATH="$HOME/go"
$ export PATH="$GOPATH/bin:$PATH"

```

## Link swagger 
```bash

# url para o swagger local, se atentar para a porta configurada para a palicação
http://localhost:5000/swagger/index.html#/

```
