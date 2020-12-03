This project includes two ways to get informations from kubernetes through outside or inside you
cluster.

First compile the application for Linux:
```go build -o ./app .```

Then package it to a docker image using the provided Dockerfile to run it on Kubernetes.

If you are running a Minikube cluster, you can build this image directly on the Docker engine of the Minikube node without pushing it to a registry. To build the image on Minikube:
```docker build -t in-cluster .```

If you are not using Minikube, you should build this image and push it to a registry that your Kubernetes cluster can pull from.

If you have RBAC enabled on your cluster, use the following snippet to create role binding which will grant the default service account view permissions.
```kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default```

Then, run the image in a Pod with a single instance Deployment:
```kubectl run --rm -i demo --image=in-cluster --image-pull-policy=IfNotPresent```
