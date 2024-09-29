FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o websocket-server main.go

EXPOSE 8080

CMD ["./websocket-server"]

# docker build -t websocket-server .
# docker run -p 8080:8080 websocket-server

