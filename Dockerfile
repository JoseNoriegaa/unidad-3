FROM golang:1.14.2

ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8
ENV DEBIAN_FRONTEND=noninteractive
ENV DEBIAN_FRONTEND=teletype

RUN apt update
RUN apt install git -y

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/jinzhu/gorm/dialects/mysql

WORKDIR /go/src/github.com/josenoriegaa/unidad-3

COPY . .

EXPOSE 3000
