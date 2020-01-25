FROM golang:1.13.6-alpine AS build

WORKDIR /go/src/github.com/analogworker/minikube-sample

RUN apk add --no-cache git make

COPY . /go/src/github.com/analogworker/minikube-sample

RUN go build -o api

FROM alpine AS app

COPY --from=build /go/src/github.com/analogworker/minikube-sample/api app/api

EXPOSE 3000

ENTRYPOINT ["/app/api"]