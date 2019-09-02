FROM golang:alpine


ADD src src

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main src/main.go
RUN rm -rf src


CMD ["./main"]


EXPOSE 8234
EXPOSE 80
EXPOSE 8080
EXPOSE 8000
EXPOSE 8111
