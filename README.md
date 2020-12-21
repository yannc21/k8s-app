
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

Automations executed during the pipelines :
* Create binary repositories in Artifactory
* Index repositories and build for Xray scans
* Create watches per pipeline
* Promote artifacts and assign custom properties

## How to run the demo

1. Requirements : 
* Running environment : JFrog platform (Artifactory + Xray + Pipelines)
* Create (manually) 3 integrations in Pipelines (github, artifactory, k8s)
* Create 1 policy in Xray

2. Adapt the pipelines to your environment (see next section)

3. For K8S deployment, create the API objects (see `k8s-object-api` folder) before running the pipelines
* Run the `k8s_project_init` pipeline


## How to adapt the pipelines 

in `back/CI` and `front/CI` folders, you may want to change :

* `pipelines.steps.yaml`, modify: 
    * the integration names `artifactory_eu` and  `yann_k8s`
    * all environment variables 
          
* `pipelines.resources.yaml`, modify:
    * the integration names `artifactory_eu` and  `yann_github`
    * branch name : `jfrog_ode`
    * branch path : `cyan21/k8s-app`
    
* for Xray, change the policy name in `watch.json`

> if you want to change the repo names (created during the pipelines), edit `repo.yaml` and adapt accordingly the pipelines environment variables

> naming conventions for variables : 
> * `envVar*` for environment variables
> * `runVar*` for variables generated during the pipelines

> property bags are used to pass variables between pipelines