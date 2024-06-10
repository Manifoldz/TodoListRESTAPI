FROM golang:1.22

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait_postgres.sh

RUN go mod download
RUN go build -o todosrv ./cmd/main.go

CMD ["./todosrv"]