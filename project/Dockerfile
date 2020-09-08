FROM golang:1.14-alpine AS build

RUN apk add --no-cache git gcc libc-dev curl

WORKDIR /build

ADD . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 go build -o pokemon

FROM alpine

RUN apk add --no-cache bash

WORKDIR /usr/local/bin

COPY --from=build /build/pokemon .
COPY --from=build /build/config/config.json ./config/config.json
COPY --from=build /build/migration ./migration

RUN chmod +x pokemon

EXPOSE 3000
CMD ["pokemon"]
