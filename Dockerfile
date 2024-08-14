FROM golang:1.22.3 as builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go-ebanx ./cmd/go-ebanx

FROM scratch

WORKDIR /app

COPY --from=builder /go-ebanx /go-ebanx

EXPOSE 8080

CMD ["/go-ebanx"]
