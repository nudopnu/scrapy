FROM golang:1.23 as builder
WORKDIR /app
COPY src .
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 go build -o scraper

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/scraper .
CMD ["./scraper"]