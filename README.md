
# K8S DEMO

## What is it

* Build a containerized web app composed of a front and back end (Gradle + NPM)
* Deploy it to K8S cluster via a helm chart

After loading the YAML files into JFrog Pipelines, the following pipelines will be created

| Pipeline Name | Description | 
| ----------- | ----------- |
| k8s_project_init | Create Helm Chart and Docker repository  |
| k8s_backapp_gradle | Build and promote Gradle app|
| k8s_backapp_gradle_docker | Containerize Gradle app|
| k8s_backapp_gradle_deployment | Deploy Gradle app on K8S cluster via Helm Chart |
| k8s_frontapp_npm | Build and promote ReactJS app|
| k8s_frontapp_js_docker | Containerize ReactJS app|
| k8s_frontapp_js_deployment | Deploy ReactJS app on K8S cluster via Helm Chart |


### Workarounds

* in `k8s_project_init`, build a docker image with `helmV3`
* in all docker pipelines, a `setup_env` injects `insecured registries` in the `/etc/docker/daemon.json`

Automations executed during the pipelines :
* Create binary repositories in Artifactory
* Index repositories and build for Xray scans
* Create watches per pipeline
* Promote artifacts and assign custom properties

> naming conventions for variables : 
> * `envVar*` for environment variables
> * `runVar*` for variables generated during the pipelines

> property bags are used to pass variables between pipelines

## How to run the demo

1. Requirements : 
* Running environment : JFrog platform (Artifactory + Xray + Pipelines)
* Create 3 integrations in JFrog Pipelines:
** github called `yann_github` 
** artifactory called `artifactory_eu`
** kubernetes called `yann_k8s`
* Create 1 policy in Xray called `global` enabling `fail_build`

2. For K8S deployments, create the API objects (see `k8s-object-api` folder) 
3. Run the `k8s_project_init` pipeline


## How to adapt the pipelines

### Names of the integrations

in `chart/CI/jfrog`, `back/CI/jfrog` and `front/CI/jfrog` folders, change :

* `pipelines.steps.yaml`, modify: 
    * the integration names `artifactory_eu` and  `yann_k8s`
    * all environment variables 
          
* `pipelines.resources.yaml`, modify:
    * the integration names `artifactory_eu` and  `yann_github`
    * branch name : `jfrog_ode`
    * branch path : `cyan21/k8s-app`

### Names of the repositories

* Edit the `chart/CI/jfrog/repo.yaml` and adapt accordingly the pipelines `environment variables` in *ALL* `yaml` files in `chart/CI/jfrog`, `back/CI/jfrog` and `front/CI/jfrog` folders
* the `FROM` instruction in the `Dockerfile`

### Configurations of the Watches

* Edit the `json` files and `init.sh` in  `chart/CI/jfrog`