# mc-monitoring

This repository is created to figure out how the istio-promethus setup can be done for multi-cluster setup.

### Setting up the clusters (cluster-0, cluster-1 and cluster-2):

```
cd <cluster-name>
```


```
kind create cluster --name=<cluster-name> --config=backend/cluster-spec/kind-config.yaml
istioctl install -f backend/istio-spec/istio-kind-profile.yaml
k apply -f backend/apis/namespace.yaml
k apply -f backend/apis/gateway.yaml
k apply -f backend/apis/manifest.yaml
```


Above setup installs Istio-Ingressgateway as a NodePort service, this is needed as we install k8s on local machine without a network loadbalancer.

## API

the code for api that is deployed in all the three clusters is within the api folder.
Details about it are in readme which is within the folder.