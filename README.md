# APIRest-Golang

API construída em Golang, utilizando técnicas de containerização.

# Tecnologias utilizadas

Linguagem: Golang
Framework web: Gin Gonic
Banco de dados: PostgreSQL
Containerização: Docker e Docker Compose

# Estrutura de Pastas 

rest-api/
├── cmd/

│   └── main.go                 # Ponto de entrada (Inicia o servidor e rotas)

├── controller/

│   └── product_controller.go   # Porta de entrada das requisições (HTTP), valida os dados de entrada e envia a resposta ao cliente

├── usecase/

│   └── product_usecase.go      # Regras de negócio (Lógica principal), contém a lógica principal da aplicação, servindo de ponte entre controller e repository

├── repository/

│   └── product_repository.go   # Comunicação com o Banco de Dados (SQL)

├── model/

│   ├── product.go              # Estrutura da entidade Produto

│   └── response.go             # Estruturas de resposta da API

├── db/

│   └── conn.go                 # Configuração da conexão com Postgres

├── docker-compose.yml          # Orquestração (App + DB)

├── dockerfile                  # Receita da imagem Docker da API

├── go.mod                      # Gerenciador de dependências

├── go.sum                      # Checksum das dependências


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