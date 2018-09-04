FROM golang:alpine AS build
ADD . /go/src/github.com/srizzling/whoami-aws
RUN go install github.com/srizzling/whoami-aws

FROM alpine
RUN apk add --update \
    ca-certificates
RUN mkdir -p /opt/srizzling
WORKDIR /opt/srizzling
COPY --from=build /go/bin/whoami-aws /opt/srizzling/whoami-aws
EXPOSE 8081
ENTRYPOINT ["/opt/srizzling/whoami-aws"]
