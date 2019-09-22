FROM golang:1.12 as build-env

RUN mkdir /api
WORKDIR /api
ADD go.mod go.sum ./

RUN go mod download

ADD . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/api
FROM scratch
COPY --from=build-env /go/bin/api /go/bin/api
ENTRYPOINT ["/go/bin/api"]