
# K8S DEMO

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

> if you want to change the repo name (created during the pipelines), edit `repo.yaml`