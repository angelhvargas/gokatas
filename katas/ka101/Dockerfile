FROM golang:alpine as build

WORKDIR ./

ENV CGO_ENABLED=0

COPY . .

RUN go get -d -v ./...

RUN go build main.go

CMD ["./main", "test.csv"]

