# APIRest-Golang

API robusta construída em Golang, focada em performance e escalabilidade, utilizando técnicas modernas de containerização e arquitetura em camadas.

# Tecnologias Utilizadas

Linguagem: Golang

Framework Web: Gin Gonic

Banco de Dados: PostgreSQL

Containerização: Docker & Docker Compose

# Estrutura de Pastas

A organização segue os princípios de separação de responsabilidades:

cmd/        Ponto de entrada da aplicação (main.go). Inicia o servidor e rotas.

controller/	Porta de entrada das requisições (HTTP). Valida entradas e envia respostas.

usecase/	Regras de negócio. Ponte lógica entre o controller e o repository.

repository/	Comunicação direta com o Banco de Dados (SQL).

model/	    Definição das entidades (Product) e structs de resposta.

db/	        Configurações de conexão com o Postgres.


# Inicializar o projeto:
    # Criando go.mod: 
    go mod init go-api

    # Instalando framework web utilizado para criar uma Rest API: 
    go get github.com/gin-gonic/gin

    # Comando para rodar o container:
    docker-compose up -d go_db

    # Verificar se o container está rodando:
    docker container ls

    # Gerar imagem da API com as instruções do Dockerfile:
    docker build -t go-api .

    # Verificar se a imagem foi criada:
    docker image ls

    # Subir banco de dados e aplicação juntos:
    docker compose up -d 

    # Reconstruir o container:
    docker compose up -d --build

    # Ver se a API está viva e se tem erros:
    docker logs -f go_app
