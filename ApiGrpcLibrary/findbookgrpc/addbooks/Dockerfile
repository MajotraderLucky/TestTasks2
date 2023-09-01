FROM golang:1.20

WORKDIR /app

COPY . .

RUN go build -o addbooks .

CMD sh -c "sleep 10 && go run addbooks.go"