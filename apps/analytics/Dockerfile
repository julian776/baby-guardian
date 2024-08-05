# To build the image you can run the following command from
# the root of the project baby-guardian not from analytics directory:
# DOCKER_BUILDKIT=1 docker build -t analytics --file apps/analytics/Dockerfile .

FROM golang as builder

WORKDIR /app

COPY ./go.work .

COPY ./protos ./protos

COPY apps/analytics/go.mod apps/analytics/go.sum ./apps/analytics/

RUN go mod download

COPY apps/analytics ./apps/analytics

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o analytics ./apps/analytics/cmd


FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/analytics /usr/local/bin/

RUN chmod +x /usr/local/bin/analytics

CMD ["/usr/local/bin/analytics"]
