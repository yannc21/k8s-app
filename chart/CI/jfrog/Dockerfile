FROM kaivengers-docker.artifactory-eu-yannc4-0.soleng-emea-staging.jfrog.team:80/jfrog/pipelines-u18node:10.21.0
LABEL maintainer="yannc@jfrog.com"

ENV JFROG_CLI_OFFER_CONFIG=false

WORKDIR /tmp

RUN apt-get update && apt-get install -y curl jq

RUN curl -fL https://getcli.jfrog.io | sh &&  chmod 755 jfrog &&  mv jfrog /usr/local/bin/

CMD ["/bin/bash", "-c", "jfrog --version"]
