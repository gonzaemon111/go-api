FROM golang:1.8.3

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

RUN apt-get update && apt-get install -y postgresql-client
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm/dialects/postgres
