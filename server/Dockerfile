FROM golang:1.19

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o /usr/local/bin/app .

EXPOSE 8000

CMD ["app", "8000"]
