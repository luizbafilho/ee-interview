FROM golang:1.22.5-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "test", "./..."]