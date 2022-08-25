# Start from the latest golang base image
FROM golang:latest as builder

LABEL maintainer="Hüsamettin ARABACI <husamettin.arabaci@hsmteknoloji.com>"

WORKDIR /app

COPY ./Services/BE_Services/memcacheserver/go.mod ./
COPY ./Services/BE_Services/memcacheserver/go.sum ./

RUN go mod download

COPY ./Services/BE_Services/memcacheserver/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o memcacheserver .


######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache tzdata
ENV TZ=Europe/Istanbul
WORKDIR /root/

ENV CONTAINER_TYPE=memcacheserver

COPY --from=builder /app/memcacheserver .

CMD ["./memcacheserver"]