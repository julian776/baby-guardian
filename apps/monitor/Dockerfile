# To build the image you can run the following command from
# the root of the project baby-guardian not from monitor directory:
# DOCKER_BUILDKIT=1 docker build -t monitor --file apps/monitor/Dockerfile .

FROM golang as builder

WORKDIR /app

COPY ./go.work .

COPY ./protos ./protos

COPY apps/monitor/go.mod apps/monitor/go.sum ./apps/monitor/

RUN go mod download

COPY apps/monitor ./apps/monitor

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o monitor ./apps/monitor/cmd


FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/monitor /usr/local/bin/

RUN chmod +x /usr/local/bin/monitor

CMD ["/usr/local/bin/monitor"]
