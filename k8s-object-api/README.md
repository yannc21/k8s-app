# Deploying the Helm Chart

Before deploying the helm chart :

* create namespace

````
sudo kubectl create ns kaivengers
````

* create service account

````
sudo kubectl create sa ironman -n kaivengers
````

* create secrets to use Artifactory as a private container registry (there are referenced in the helm chart)

````
sudo kubectl create secret docker-registry avengers-registry \
    --docker-server=avengers-docker.artifactory-eu-yannc4-0.soleng-emea-staging.jfrog.team:80 \
    --docker-username=admin --docker-password=password --docker-email=yannc@jfrog.com -n kaivengers 

sudo kubectl create secret docker-registry kaizoku-registry \
    --docker-server=kaizoku-docker.artifactory-eu-yannc4-0.soleng-emea-staging.jfrog.team:80 \
    --docker-username=admin --docker-password=password --docker-email=yannc@jfrog.com -n kaivengers   
````

* create secret for service account + role + rolebinding

````
sudo kubectl apply -f setup.yaml -n kaivengers
````

# JFrog Pipelines configuration

to deploy a helm chart via JFrog Pipelines, you'll have to create a K8S integration which requires a kubeconfig.

It's recommended to use a service account (and not a normal user) to configure the kubeconfig.

So to craft a kubeconfig with a service account : 

1/ as a normal user, connect to the targeted K8S cluster with kubectl

2/ have a look to your kubeconfig 

````
sudo kubectl config view --flatten --minify > myKubeConfig
// kubeconfig should be in <HOME>/.kube/config
````

3/ Extract the token of the service account 

````
sudo kubectl describe secrets $(sudo kubectl get secret -n kaivengers -o custom-columns=NAME:.metadata.name | grep ironman)  -n kaivengers | grep "^token" | tr -d [[:space:]] | cut -d: -f2
````

4/ Inject the token into a kubeconfig 

See example of kubeconfig in the folder
