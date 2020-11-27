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
