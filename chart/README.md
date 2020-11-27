# For demo

> only tested on minikube  

this simple helm chart deploys :
* 1 front end in a pod and  1 service
* 1 back end in a pod and 1 service

relies on https://github.com/cyan21/k8s-app.git

> using the nodePort type for each service

### HOW TO USE

```
helm install <RELEASE_NAME> <PATH_TO_SRC>  --set services.back.ip=<MINIKUBE_IP>

```
Example:
```
git clone https://github.com/cyan21/k8s-helm.git

helm install the-app k8s-helm --set services.back.ip=$(minikube ip)
```

