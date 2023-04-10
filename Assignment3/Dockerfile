FROM golang:latest

WORKDIR /assignment3

COPY . /assignment3

RUN go mod download

RUN go build -o Assignment3 ./main.go

EXPOSE 8080

CMD ["./assignment3"]