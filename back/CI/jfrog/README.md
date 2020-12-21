# PRE REQUISITES

JFrog platform enabling Artifactory, Xray, Pipelines

| Service | Type | Name | Description | 
| ----------- | ----------- |----------- | ----------- |
| Artifactory | Gradle virtual repo | kaizoku-grdl | aggregate nuget dev local and remote | 
| Artifactory | Gradle local repo   | kaizoku-dev-grdl-local | | 
| Artifactory | Gradle local repo   | kaizoku-rc-grdl-local | for Gradle promotion | 
| Artifactory | Gradle remote repo  | kaizoku-jcenter-grdl-remote | | 
| Artifactory | Docker virtual repo | kaizoku-docker | | 
| Artifactory | Docker local repo   | kaizoku-dev-docker-local | | 
| Artifactory | Docker local repo   | kaizoku-release-docker-local | | 
| Artifactory | Docker remote repo  | kaizoku-dockerhub-remote | | 
| Artifactory | Docker remote repo  | kaizoku-bintray-remote | | 
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