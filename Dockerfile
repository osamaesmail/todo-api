FROM golang:1.12 as build-env

RUN mkdir /api
WORKDIR /api
ADD go.mod go.sum ./

RUN go mod download

ADD . .

FROM scratch
COPY --from=build-env /go/bin/api /go/bin/api
ENTRYPOINT ["/go/bin/api"]