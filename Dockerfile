FROM golang:1.21.5 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates

RUN apk add --no-cache bash
COPY ./docker-entrypoint.sh /
RUN chmod +x /docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]

WORKDIR /root/

COPY --from=builder /app/build .
COPY --from=builder /app/config.json .

CMD ["./build"]
