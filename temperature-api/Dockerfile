FROM golang:1.26

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o temperature-api

EXPOSE 8080

CMD ["./temperature-api"]