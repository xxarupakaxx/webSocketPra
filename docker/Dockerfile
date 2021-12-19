FROM golang:latest

WORKDIR /server

COPY go.mod go.sum ./
RUN go mod download


COPY /server .
COPY /server/chat .
RUN go build -o app
CMD ./app

