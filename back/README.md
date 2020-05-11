
## build

```
docker build -t backapp:1.0 --rm --build-arg GIT_USER=cyan21 --build-arg GIT_TOKEN=<TOKEN> .
```

## run

env.txt

```
# FQDN can be the hostname or IP 
# default 0.0.0.0
APP_FQDN=""

# default 8090 
APP_PORT=""

# if true will read data.json 
# if unset or empty or false will query  https://www.quotes.net
OFFLINE_MODE=true

# credentials for https://www.quotes.net
# only if OFFLINE_MODE=false or unset
API_UID=""
API_TOKENID=""

```
Execute
```
docker run --env-file env.txt -p 8090:8090 -d backapp:1.0
```
