
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

Automations executed during the `k8s_project_init` pipeline :
* Create binary repositories in Artifactory
* Index repositories and build for Xray scans
* Create watches per pipeline
* Create base images

> naming conventions for variables : 
> * `envVar*` for environment variables
> * `runVar*` for variables generated during the pipelines

> property bags are used to pass variables between pipelines

### Workarounds

* in `k8s_project_init`, build a docker image with `helmV3`
* in all docker pipelines, a `setup_env` injects `insecured registries` in the `/etc/docker/daemon.json`

## How to run the demo

> Requirements : 
>   Running environment : JFrog platform (Artifactory + Xray + Pipelines)

1. In JFrog Pipelines, create 3 integrations:
* github called `yann_github` pointing to this code repo
* artifactory called `artifactory_eu` configured to connect to your Artifactory 
* kubernetes called `yann_k8s` configured to connect to your your K8S cluster (see `k8s-object-api` folder)

> if you want to change the integration names see next section

2. In JFrog Xray, create 1 policy in Xray called `global` enabling `fail_build`

3. In the source repo, modify the repository URLs : 
* in `pipelines.steps.yaml` files in  `chart/CI/jfrog`, `back/CI/jfrog` and `front/CI/jfrog`
* in `values.yaml` in `chart/content` 

4. In JFrog Pipelines, start manually the `k8s_project_init` pipeline

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