# syntax=docker/dockerfile:1

FROM golang:1.21-alpine AS build-stage

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-gin-finance

# Deploy the application binary into a lean image
#FROM gcr.io/distroless/base-debian11 AS build-release-stage
FROM golang:1.21-alpine

WORKDIR /

COPY --from=build-stage /go-gin-finance /go-gin-finance
#copying the config files as required input for the application
COPY --from=build-stage /app/config /config

EXPOSE 8080

# Run
CMD ["./go-gin-finance"]