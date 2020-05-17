FROM node:14-alpine3.10

WORKDIR /tmp/

ENV URL="morty.com"

COPY temp/ app/

COPY updateAndRun.sh .
RUN chmod u+x updateAndRun.sh

RUN npm i -g serve

EXPOSE 80

ENTRYPOINT ["/bin/sh", "-c", "./updateAndRun.sh $URL"]
