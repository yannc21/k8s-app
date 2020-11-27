# Deploying the Helm Chart

Before deploying the helm chart :

* create secret for Artifactory

````
sudo kubectl create secret docker-registry art7 \
    --docker-server=avengers-docker.artifactory-eu-yannc4-0.soleng-emea-staging.jfrog.team:80 \
    --docker-username=admin --docker-password=password --docker-email=yannc@jfrog.com
````

* create namespace

````
sudo kubectl create ns ninjavengers
````

* create service account

````
sudo kubectl create sa ironman -n ninjavengers
````

* create secret for service account

````
sudo kubectl apply -f secret.yaml -n ninjavengers
````

* create role

````
sudo kubectl apply -f role.yaml
````

* create role binding

````
sudo kubectl apply -f rolebinding.yaml
````

# JFrog Pipelines configuration

to deploy a helm chart via JFrog Pipelines, you'll have to create a K8S integration which requires a kubeconfig.

It's recommended to use a service account (and not a normal user) to configure the kubeconfig.

So to craft a kubeconfig with a service account : 

1/ as a normal user, connect to the targeted K8S cluster with kubectl

2/ have a look to your kubeconfig 

````
kubectl config view --flatten --minify
// kubeconfig should be in <HOME>/.kube/config
````

3/ Extract the token of the service account via these commands

````
sudo kubectl get secret -n ninjavengers
sudo kubectl describe secret ironman-token-xxxxxx -n ninjavengers
````

4/ Inject the token into a kubeconfig 

See example of kubeconfig in the folder
