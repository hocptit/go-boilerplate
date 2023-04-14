# Builder
FROM golang:1.19.3-alpine3.15 as builder

RUN apk update && apk upgrade && \
    apk --update add git make bash build-base
# sudo apt install build-essential

WORKDIR /app

COPY . .

RUN make build

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app

WORKDIR /app

EXPOSE 5000

COPY --from=builder /app/engine /app/
COPY --from=builder /app/.env /app/

CMD /app/engine