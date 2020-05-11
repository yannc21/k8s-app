FROM mhart/alpine-node:14.2 AS builder

WORKDIR /tmp/ 

ARG GIT_USER
ARG GIT_TOKEN

RUN apk update && apk add git 
RUN git clone https://$GIT_USER:$GIT_TOKEN@github.com/cyan21/k8s-app

WORKDIR /tmp/k8s-app/front

RUN npm i 
RUN npm run build 

FROM nginx:1.18-alpine
RUN apk update && apk add curl 

COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /tmp/k8s-app/front/build /usr/share/nginx/html 

EXPOSE 9000 

CMD ["nginx", "-g", "daemon off;"]
