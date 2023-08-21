FROM golang:1.21 as builder

# Set the working directory within the container
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project directory into the container's working directory
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM alpine:latest

# Set the working directory
WORKDIR /root/

COPY --from=builder /app/main .

ARG PORT
ENV PORT=8080
EXPOSE $PORT

CMD ["./main"]
