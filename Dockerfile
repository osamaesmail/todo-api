FROM golang:1.12

ADD . /go/src/todo-api
WORKDIR /go/src/todo-api

RUN go get .
RUN go build

EXPOSE 3000

ENTRYPOINT ["/go/bin/todo-api"]