FROM golang:1.21-rc-bullseye AS builder

WORKDIR /app
COPY . .

RUN apk add npm
RUN make build 

FROM debian:bullseye-20230703-slim

ARG MODE=release

ENV GIN_MODE=$MODE

RUN useradd -ms /bin/bash appuser

USER appuser
WORKDIR /home/appuser/app

COPY --from=builder /app/bin/main .
COPY --from=builder /app/web web/ 

EXPOSE 8080

ENTRYPOINT ["./main"]
