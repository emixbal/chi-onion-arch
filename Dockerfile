# Building the binary of the App
FROM golang:1.20-alpine AS build

# `user-service` should be replaced with your project name
WORKDIR /go/src/user-service

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app -installsuffix cgo .

# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest

WORKDIR /app

# Copy the binary and configuration files from the build stage
COPY --from=build /go/src/user-service/app .
COPY --from=build /go/src/user-service/config.json .

# Exposes port 3000 because our program listens on that port
EXPOSE 3000

CMD ["./app"]
