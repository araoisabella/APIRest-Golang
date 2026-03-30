FROM golang:1.25.0

#set working directory
WORKDIR /go/scr/app

#copy the source code 
COPY . .

#expose the port 
EXPOSE 8000

#build the go app
RUN go build -o main cmd/main.go

#run the executable
CMD ["./main"]

#instrucoes para gerar imagem da aplicacao