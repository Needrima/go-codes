FROM golang:1.19-alpine3.16

RUN mkdir -p /app

WORKDIR /app

ENV DB_USERNAME=admin \
    DB_PASSWORD=admin 

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

CMD ["/app/main"]