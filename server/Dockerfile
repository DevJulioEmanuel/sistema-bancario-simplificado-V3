FROM golang:1.26.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o servidor ./cmd

EXPOSE 8080

CMD ["./servidor"]