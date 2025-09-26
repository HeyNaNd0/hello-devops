# Build stage
FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /app/server

# Runtime stage
FROM scratch
ENV PORT=8080
EXPOSE 8080
COPY --from=build /app/server /server
ENTRYPOINT ["/server"]
