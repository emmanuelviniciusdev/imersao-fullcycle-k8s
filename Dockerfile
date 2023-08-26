FROM golang:1.20.3

RUN mkdir src/app

WORKDIR src/app

COPY . .

RUN go build -o main .

CMD ["./main"]
