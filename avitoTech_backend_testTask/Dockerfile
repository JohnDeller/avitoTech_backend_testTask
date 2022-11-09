FROM golang:1.14-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o avitoTech_backend_testTask ./cmd/main.go

CMD ["./avitoTech_backend_testTask"]