# Start from the latest golang base image
FROM golang:latest as builder

LABEL maintainer="Hüsamettin ARABACI <husamettin.arabaci@hsmteknoloji.com>"

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o webserver .


######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache tzdata
ENV TZ=Europe/Istanbul
WORKDIR /root/

ENV CONTAINER_TYPE=webserver

COPY --from=builder /app/webserver .
RUN mkdir ./views
COPY ./views/* ./views/

EXPOSE 80
CMD ["./webserver"]