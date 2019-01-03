FROM golang:1.10

WORKDIR /go/src/back
COPY . .
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./main.go" -command="./main"
EXPOSE 6060
