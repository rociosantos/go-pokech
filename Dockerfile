# First stage: build the executable.
FROM golang:buster AS base

# Dev dependencies
# RUN GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.21

# Copy code
COPY . /app
WORKDIR /app

# Install dependencies
RUN go mod tidy
# Build app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Run app
CMD ["./app"]

# Expose port
EXPOSE 80

# docker build --tag pokech:1.0 .
# docker run --name pokech -p 80:8085 pokech:1.0