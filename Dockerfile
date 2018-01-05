FROM golang:alpine AS buildenv
WORKDIR /go/src/github.com/dhawton/docker-viz-monitor/
RUN apk --no-cache add ca-certificates git && \
    update-ca-certificates
COPY ./main.go .
RUN go get -v ./... && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=buildenv /go/src/github.com/dhawton/docker-viz-monitor/app .
CMD ["./app"]
