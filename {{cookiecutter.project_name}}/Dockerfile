# Building the binary of the App
FROM docker.io/golang:1.20 AS build

WORKDIR /go/src/{{cookiecutter.project_name}}

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app cmd/{{cookiecutter.project_name}}/main.go


# Moving the binary to the 'final Image' to make it smaller
FROM docker.io/alpine:latest AS release

WORKDIR /app

COPY --from=build /go/src/{{cookiecutter.project_name}}/app .

# Add packages
RUN apk -U upgrade \
    && apk add --no-cache dumb-init ca-certificates \
    && chmod +x /app/app

# Exposes port 3000 because our program listens on that port
EXPOSE 3000

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
