FROM golang:1.16-alpine

ENV GO111MODULE=on

LABEL maintainer="Zhenis Duissekov  <zduissekov@gmail.com>"

RUN ln -snf /usr/share/zoneinfo/Asia/Almaty /etc/localtime && echo Asia/Almaty > /etc/timezone

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /dummy_service

RUN chmod u+x /dummy_service

COPY . .

EXPOSE 1111

#HEALTHCHECK --interval=5m --timeout=3s CMD curl -f http://localhost:1111/healthcheck || exit 1

#CMD ["/dummy_service"]
