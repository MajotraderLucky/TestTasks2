FROM golang:1.20

WORKDIR /app

COPY . .

RUN go build -o main .

CMD sh -c "sleep 15 && go run main.go"