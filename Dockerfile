FROM golang:1.25-alpine3.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o build/app cmd/main.go

FROM alpine:3.24 AS runner

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/build/app .

ARG PORT=8080

ENV GIN_MODE=release

ENV PORT=${PORT}

EXPOSE ${PORT}

CMD ["./app"]