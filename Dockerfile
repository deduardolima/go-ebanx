FROM golang:1.22.3 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN swag init --dir cmd/go-ebanx,internal/infra/web --output ./docs

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go-ebanx ./cmd/go-ebanx

FROM scratch

WORKDIR /app

COPY --from=builder /go-ebanx /go-ebanx

EXPOSE 8080

CMD ["/go-ebanx"]
