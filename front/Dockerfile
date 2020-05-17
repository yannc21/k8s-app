FROM node:14-alpine3.10 

WORKDIR /tmp/ 

ARG GIT_USER
ARG GIT_TOKEN
ENV URL="morty.net"

RUN apk update && apk add git 
RUN git clone https://$GIT_USER:$GIT_TOKEN@github.com/cyan21/k8s-app

WORKDIR /tmp/k8s-app/front

RUN npm i 
RUN npm run build 
RUN npm i -g serve 

COPY updateAndRun.sh . 
RUN chmod u+x updateAndRun.sh

EXPOSE 80 

ENTRYPOINT ["/bin/sh", "-c", "./updateAndRun.sh $URL"]
