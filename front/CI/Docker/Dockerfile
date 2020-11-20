FROM nginx:1.19.4-alpine

WORKDIR /tmp/

ENV URL="morty.com"
COPY app.tar.gz .
#COPY nginx.conf /etc/nginx/nginx.conf

RUN tar -xzvf app.tar.gz -C /usr/share/nginx/
RUN ls -l && mv /usr/share/nginx/build/* /usr/share/nginx/

EXPOSE 80
ENTRYPOINT ["nginx","-g","daemon off;"]