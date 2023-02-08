FROM golang:1.19

# Port to expose
ENV PORT=8080

WORKDIR /go/src/app

COPY . .

# Install dependencies
RUN go get -d -v ./...

# Build the app
RUN go build -o main .

# Expose the port
EXPOSE 8080

# Run the app
CMD ["./main"]
