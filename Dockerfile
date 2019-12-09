FROM golang:alpine AS build-env
WORKDIR /go/src/github.com/youtangai/files_api_mock
COPY ./ ./
RUN go build -o server main.go

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/youtangai/files_api_mock/server /usr/local/bin/server
WORKDIR /var/cloud
ENV APP_WORK_DIR /var/cloud

EXPOSE 8082
CMD ["/usr/local/bin/server"]