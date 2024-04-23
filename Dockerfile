FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go get github.com/githubnemo/CompileDaemon
RUN go mod tidy

ENTRYPOINT ComileDaemon --build="go build -o app" --command="./app/app"

