FROM golang:1.18.4-alpine3.15 as build-env
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download -x
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static
COPY --from=build-env /go/bin/app /
COPY ./.env.example ./.env
CMD ["/app"]
