FROM golang:1.20 AS build-stage

WORKDIR /app

COPY src/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /httpd-test

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /httpd-test /httpd-test

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/httpd-test"]
