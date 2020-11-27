# PRE REQUISITES

JFrog platform enabling Artifactory, Xray, Pipelines

| Service | Type | Name | Description | 
| ----------- | ----------- |----------- | ----------- |
| Artifactory | NPM virtual repo | avengers-npm | aggregate npm dev local and remote | 
| Artifactory | NPM local repo   | avengers-dev-npm-local | | 
| Artifactory | NPM local repo   | avengers-rc-npm-local | for npm promotion | 
| Artifactory | NPM remote repo  | avengers-npmjs-remote | | 
| Artifactory | Docker virtual repo | avengers-dev-docker | aggregate docker dev local and remote | 
| Artifactory | Docker local repo   | avengers-dev-docker-local | | 
| Artifactory | Docker local repo   | avengers-rc-docker-local | for mvn promotion | 
| Artifactory | Docker remote repo  | avengers-dockerhub-remote | | 
| Pipelines   | Github Integration | yann_github | pointing to https://github.com/cyan21 |
| Pipelines   | Artifactory Integration | artifactory_eu | |

## Repository creation

````
curl -uadmin:chaysinh -X PATCH "http://localhost:8081/artifactory/api/system/configuration" -H "Content-Type: application/yaml" -T repo.yml
````

if you change the repo names, make sure to edit : 
* repo name in the pipelines.steps.yaml (pipeline variables)


## Integration creation

integrations have to be created manually for now on JFrog pipelines




