FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /go/src/user-service

COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /go/bin/user-service

FROM scratch

COPY --from=builder /go/bin/user-service /go/bin/user-service

EXPOSE 8082

ENTRYPOINT [ "/go/bin/user-service" ]