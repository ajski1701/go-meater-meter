FROM registry.access.redhat.com/ubi8/go-toolset:latest AS build

WORKDIR /app

COPY . .
RUN go mod download && go build

FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
COPY --from=build /app/go-meater-meter /bin/go-meater-meter