# For demo

this simple helm chart deploys :
* 1 front end (1 pod + 1 service)
* 1 back end (1 pod + 1 service)

> using the nodePort type for each service

# Deploy Chart

```
git clone https://github.com/cyan21/k8s-app.git
cd k8s-app
helm install the-app chart/ --set services.back.ip=$(minikube ip)
```

# Publish Chart to Artifactory

* Create Helm repositories
```
curl -uadmin:chaysinh -X PATCH "http://localhost:8081/artifactory/api/system/configuration" -H "Content-Type: application/yaml" -T repo.yml
```

* Package chart
```
helm package chart/
helm install the-app chart/ --set services.back.ip=$(minikube ip)
```

* Upload chart to Artifactory
```
curl -u<USERNAME>:<PASSWORD> -T <PATH_TO_FILE> "http://localhost:8081/artifactory/kaivengers-dev-helm-local/<TARGET_FILE_PATH>" 
```