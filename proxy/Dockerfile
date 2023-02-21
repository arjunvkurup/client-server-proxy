FROM golang:1.19

WORKDIR /usr/src/app

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o /usr/local/bin/app .

EXPOSE 8080

CMD ["app", "8080"]
