FROM golang:latest AS build
WORKDIR /app
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
COPY . .
RUN go mod download && make buildProd

FROM ghcr.io/soemii/distroless-healthcheck:static-debian12-latest
WORKDIR /
COPY --from=build /app/tmp/app /app
EXPOSE 8080
HEALTHCHECK --interval=10s --start-period=1s CMD ["healthcheck", "http", "http://localhost:8080/metrics"]
ENTRYPOINT ["/app"]