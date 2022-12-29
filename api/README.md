# API

This is a basic api that returns the hostname, cluster and service names where the api is running.
cluster-name and service-name have to be added to the build as command-line arguments.

## Debugging / Getting Started

To run the api locally please follow below commands :

```
go run main.go cluster-1 service-1
```

To check if the service builds correctly :

```
go build -o main .
```

To build the docker image and run it :

```
docker run -p 9090:9090  -d mc-api
```
This should expose the api on port `9090` and have the command-line arguments added as `cluster-1` and `service-1` for the api.


expected response should be (Hostmane would change everytime):

```
{
    "HostName": "ed63dd1d39a9",
    "Cluster": "cluster-1",
    "Service": "service-1"
}

```


You could change the command line args with below command :

```
docker run -p 9090:9090 -d mc-api cluster-2 service-2
```

---

##### Push the docker image onto `dockerhub` registry:


My Private Registry : neerajnatu/mc-api


```
docker login
Username: neerajnatu
Password:

```

---


##### Build the image for your docker-hub registry and then push it.

```
docker build -t neerajnatu/mc-api .
docker push neerajnatu/mc-api
```

