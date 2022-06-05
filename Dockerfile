FROM registry.access.redhat.com/ubi8/go-toolset:latest AS build

#WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /bin/meater-meter

FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
COPY --from=build /bin/meater-meter /bin/meater-meter