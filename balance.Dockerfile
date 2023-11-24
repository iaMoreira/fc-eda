FROM golang:1.20

WORKDIR /app/

RUN apt-get update && apt-get install -y librdkafka-dev

    

CMD ["go", "run",  "cmd/balances/main.go"]
# CMD ["tail", "-f", "/dev/null"]