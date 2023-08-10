FROM golang:1.20.2-bullseye

RUN mkdir -p /usr/src/app/logs
WORKDIR /usr/src/app/

COPY . /usr/src/app/

CMD ["go", "run", "newgrpc.go"]