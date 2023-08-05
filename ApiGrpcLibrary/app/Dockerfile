FROM golang:1.20

WORKDIR /app

COPY . .

RUN go build -o main .

CMD sh -c "sleep 30 && go run main.go"