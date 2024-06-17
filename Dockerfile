FROM golang:1.22

WORKDIR /go/src/app

# copying source code
COPY . . 

EXPOSE 8000

RUN go build -o main cmd/main.go

CMD ["./main"]