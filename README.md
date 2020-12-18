
# K8S DEMO

## What is it

* Build a containerized web app composed of a front and back end (Gradle + NPM)
* Deploy to K8S via a helm chart

Automation in the pipelines : 
* Create binary repositories in Artifactory
* Index repositories and build for Xray scans
* Promote artifacts and assign custom properties

## How to run the demo

* Environment : JFrog platform (Artifactory + Xray + Pipelines)
* Create (manually) 3 integrations in Pipelines (github, artifactory, k8s)
* Create 1 policy in Xray
* Adapt the pipelines to your environment
* Create the API objects (see `k8s-object-api` folder) before running the pipelines

> the repositories in Artifactory and the watches are created during the pipelines


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

> if you want to change the repo names (created during the pipelines), edit `repo.yaml`

> naming conventions for variables : 
> * `envVar*` for environment variables
> * `runVar*` for variables generated during the pipelines

> property bags are used to pass variables between pipelines