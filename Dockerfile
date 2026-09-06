FROM golang:1.26 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app .

FROM debian:bookworm-slim
COPY --from=build /app/app /app
CMD ["/app"]
