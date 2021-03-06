FROM golang:1.11 as builder
ENV CGO_ENABLED=0
ENV GO111MODULE="on"
WORKDIR /go/src/github.com/smarthut/automata
COPY . .
RUN go build .

FROM alpine:latest
RUN apk --no-cache add tzdata zip ca-certificates
COPY --from=builder /go/src/github.com/smarthut/automata/automata /usr/local/bin/
EXPOSE 8080
VOLUME ["/scripts"]
ENTRYPOINT ["/bin/sh", "-c"]
CMD ["automata"]
