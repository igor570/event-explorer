FROM golang:1.22.5

WORKDIR /app

#cloning mod and sum files to container directory
COPY server/go.mod server/go.sum ./

RUN go mod download

COPY server/ ./

RUN go build -o main .

EXPOSE 3100

CMD ["./main"]
