FROM golang:alpine3.11 AS builder

WORKDIR /tmp/ 

ARG GIT_USER
ARG GIT_TOKEN

RUN apk update && apk add git 
RUN git clone https://$GIT_USER:$GIT_TOKEN@github.com/cyan21/k8s-app

WORKDIR /tmp/k8s-app/back

RUN go build -o backend 

FROM alpine:3.11
RUN apk update && apk add curl 

WORKDIR /root/

COPY --from=builder /tmp/k8s-app/back/backend .
COPY --from=builder /tmp/k8s-app/back/data.json .

EXPOSE $APP_PORT

CMD ["./backend"]  
