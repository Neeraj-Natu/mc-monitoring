# mc-monitoring

This repository is created to figure out how the istio-promethus setup can be done for multi-cluster setup.

### Setting up the clusters :

```
kind create cluster --name=istio-1 --config=backed/cluster-spec/kind-config.yaml
```

### Setting up istio :

```
istioctl install -f backend/istio-spec/istio-kind-profile.yaml
```

### Deploying echoserver, gateway and virtualservice

```
k apply -f backed/apis/manifest.yaml

k apply -f backend/apis/gateway.yaml
```

Above setup installs Istio-Ingressgateway as a NodePort service, this is needed as we install k8s on local machine without a network loadbalancer.