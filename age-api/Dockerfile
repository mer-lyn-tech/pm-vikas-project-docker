FROM golang:1.26

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o age-api

EXPOSE 8080

CMD ["./age-api"]