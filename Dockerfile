FROM golang:latest AS build
WORKDIR /app
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
COPY . .
RUN go mod download && make buildProd

FROM gcr.io/distroless/static-debian12
LABEL org.opencontainers.image.source=https://github.com/
WORKDIR /
COPY --from=build /app/tmp/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]