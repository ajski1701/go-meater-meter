FROM registry.access.redhat.com/ubi8/go-toolset:latest AS build

WORKDIR /app

COPY . .
RUN go mod download && go build

FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
ENV CONFIG_FILE=/app/config.ini
COPY --from=build /app/go-meater-meter /app/go-meater-meter
ENTRYPOINT ["/app/go-meater-meter"]