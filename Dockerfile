FROM golang:1.22.5-bookworm AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final image
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/main .

RUN apt-get update && \
	apt-get install -y ca-certificates && \
	update-ca-certificates && \
	adduser --disabled-password --gecos "" appuser

# Switch to the non-root user
USER appuser

EXPOSE 8080

# Command to run the executable
CMD ["./main"]
