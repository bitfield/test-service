FROM golang:1.12-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /test-service

FROM scratch
COPY --from=build /test-service /test-service
ENTRYPOINT ["/test-service"]
