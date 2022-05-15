FROM golang:alpine3.15

WORKDIR /our-code

COPY go.mod go.sum ./
COPY vendor ./vendor
COPY cmd ./cmd
COPY src ./src

CMD ["go", "run", "./cmd/web"]
