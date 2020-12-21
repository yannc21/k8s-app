# PRE REQUISITES

JFrog platform enabling Artifactory, Xray, Pipelines

| Service | Type | Name | Description | 
| ----------- | ----------- |----------- | ----------- |
| Artifactory | NPM virtual repo | avengers-npm | aggregate npm dev local and remote | 
| Artifactory | NPM local repo   | avengers-dev-npm-local | | 
| Artifactory | NPM local repo   | avengers-rc-npm-local | for npm promotion | 
| Artifactory | NPM remote repo  | avengers-npmjs-remote | | 
| Artifactory | Generic virtual repo | avengers-generic | | 
| Artifactory | NPM local repo   | avengers-dev-generic-local | | 
| Artifactory | NPM local repo   | avengers-rc-generic-local | for promotion | 
| Artifactory | Docker virtual repo | avengers-dev-docker | aggregate docker dev local and remote | 
| Artifactory | Docker local repo   | avengers-dev-docker-local | | 
| Artifactory | Docker local repo   | avengers-rc-docker-local | for mvn promotion | 
| Artifactory | Docker remote repo  | avengers-dockerhub-remote | | 
| Artifactory | Helm virtual repo   | kaivengers-helm | | 
| Artifactory | Helm local repo     | kaivengers-dev-helm-local | | 
| Artifactory | Helm remote repo    | kaivengers-chartcenter-remote | | 
| Pipelines   | Github Integration | yann_github | pointing to https://github.com/cyan21 |
| Pipelines   | Artifactory Integration | artifactory_eu | |
| Pipelines   | K8S Integration | yann_k8s | see k8s-object-api folder to create it|

## Repository creation

see `chart/CI/jfrog/init.sh`

if you change the repo names, make sure to edit environment variables in **pipelines.steps.yaml**

## Integration creation

integrations have to be created manually for now on JFrog pipelines




