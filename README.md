This project includes two ways to get informations from kubernetes through outside or inside you
cluster.

First compile the application for Linux:
```
go build -o ./app .
```

Then package it to a docker image using the provided Dockerfile to run it on Kubernetes.

If you are running a Minikube cluster, you can build this image directly on the Docker engine of the Minikube node without pushing it to a registry. To build the image on Minikube: 
```
docker build -t in-cluster .
```

If you are not using Minikube, you should build this image and push it to a registry that your Kubernetes cluster can pull from.

If you have RBAC enabled on your cluster, use the following snippet to create role binding which will grant the default service account view permissions.
```
kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default
```

Then, run the image in a Pod with a single instance Deployment:
```
kubectl run --rm -i demo --image=in-cluster --image-pull-policy=IfNotPresent
```

You should bind service account system:serviceaccount:default:default (which is the default account bound to Pod) with 
role cluster-admin when it occurs problems like this:
```
GetK8sInfoFromIncluster err:[pods is forbidden: User "system:serviceaccount:default:default" cannot list resource "pods" in API group "" at the cluster scope]
```

Just create a yaml(with any name you like) with following contents:
```
# NOTE: The service account `default:default` already exists in k8s cluster.
# You can create a new account following like this:
#---
#apiVersion: v1
#kind: ServiceAccount
#metadata:
#  name: <new-account-name>
#  namespace: <namespace>

apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: clientgo
subjects:
  - kind: ServiceAccount
    # Reference to upper's `metadata.name`
    name: default
    # Reference to upper's `metadata.namespace`
    namespace: default
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io
```

Then, apply it by running the following command:
```
kubectl apply -f clientgo.yaml
```

If you wnat unbind them, just run:
```
kubectl delete -f clientgo.yaml
```
