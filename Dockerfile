FROM node:23 AS build-front

COPY front/package.json front/package-lock.json ./
RUN npm i
COPY front/ ./
RUN npm run build


FROM golang:alpine AS build-back
RUN apk update
RUN apk add --no-cache git ca-certificates tzdata && update-ca-certificates

ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /go/app

COPY back/go.mod back/go.sum ./
RUN go mod download
RUN go mod verify

COPY back/ ./
COPY --from=build-front /dist/ ./static/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o /build/main main.go


FROM scratch

WORKDIR /go/bin/

ENV PORT=8080
EXPOSE 8080

COPY --from=build-back /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build-back /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-back /etc/passwd /etc/passwd
COPY --from=build-back /etc/group /etc/group

COPY --from=build-back  /build/main ./main

USER appuser:appuser

ENTRYPOINT ["./main"]
