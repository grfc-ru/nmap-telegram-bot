FROM golang:latest AS builder
WORKDIR /build
COPY app .
RUN go get main
RUN CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o monitor


FROM alpine:latest
RUN apk add --no-cache nmap
WORKDIR /app
COPY --from=builder /build/monitor .
CMD sh -c "./monitor"