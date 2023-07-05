# Server builder
FROM golang:1.21-rc-bullseye AS goBuilder

WORKDIR /app
COPY . .

RUN make build-server

# Web builder
FROM node:18.16.0-bullseye-slim AS webBuilder

WORKDIR /app/web

COPY ./web .

RUN npx tailwindcss -i styles/tailwind.css -o public/style.css

# Server
FROM debian:bullseye-20230703-slim

ARG MODE=release

ENV GIN_MODE=$MODE

RUN useradd -ms /bin/bash appuser

USER appuser
WORKDIR /home/appuser/app

COPY --from=goBuilder /app/bin/main .
COPY --from=webBuilder /app/web web/ 

EXPOSE 8080

ENTRYPOINT ["./main"]
