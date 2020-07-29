FROM golang:1.14.6-alpine3.12

WORKDIR /go/src/app

COPY ./src/soma .
RUN go build main.go

CMD [ "./main" ]