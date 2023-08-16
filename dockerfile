FROM golang:latest

# Set the working directory
WORKDIR /app
COPY . .

# Install necessary system dependencies for CGo
RUN apt-get update && apt-get install -y gcc


RUN go build -o gojob

CMD ["./gojob"]
