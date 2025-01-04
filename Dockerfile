FROM node:23 AS build-front

COPY front/package.json front/package-lock.json ./
RUN npm i
COPY front/ ./
RUN npm run build


FROM golang:1.23 AS build-back

COPY back/ ./
COPY --from=build-front /dist/ /static/
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o /build/main main.go


FROM scratch

ENV PORT=8080
EXPOSE 8080

WORKDIR /go/bin/
COPY --from=build-back  /build/main ./main

ENTRYPOINT ["./main"]
