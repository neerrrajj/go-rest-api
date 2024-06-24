# Build the application from source
FROM golang:1.23rc1-bookworm AS build-stage
  WORKDIR /app

  COPY go.mod go.sum ./
  RUN go mod download

  COPY *.go ./

  RUN CGO_ENABLED=0 GOOS=linux go build -o /api

# Deploy the application binary into a lean image
FROM alpine AS build-release-stage
  WORKDIR /app

  COPY --from=build-stage /api /api

  EXPOSE 8080

  ENTRYPOINT ["/api"]