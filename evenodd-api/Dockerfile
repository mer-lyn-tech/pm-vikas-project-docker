FROM golang:1.26

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o evenodd-api

EXPOSE 8080

CMD ["./evenodd-api"]