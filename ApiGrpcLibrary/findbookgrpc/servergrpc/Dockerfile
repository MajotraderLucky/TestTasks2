FROM golang:1.20

WORKDIR /app

RUN mkdir /app/logs

COPY . .

RUN go build -o main .

CMD sh -c "sleep 20 && go run server.go"