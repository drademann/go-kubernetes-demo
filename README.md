# Go Kubernetes Demo

## The Docker

### Registry

Run a local registry accessible at localhost:5050 within Docker Desktop.

    $ docker run -d -p 5050:5000 --name registry registry:latest

### Image

Build the Greeter image.

    $ docker --debug build -t greeter .

    $ docker tag greeter localhost:5050/greeter

    $ docker push localhost:5050/greeter

## The Kubernetes

### Namespace

Create the namespace `demo-space` as a _folder_ for the resources to come.

    $ kubectl apply -f ops/kubernetes/namespace.yaml

### Pods

Tell Kubernetes to create the pods.

    $ kubectl apply -f ops/kubernetes/greeter-alpha-pod.yaml
    $ kubectl apply -f ops/kubernetes/greeter-bravo-pod.yaml

### Services

Or tell Kubernetes to create a service.

    $ kubectl apply -f ops/kubernetes/greeter-service.yaml

### Ingress

If not already there, install a nginx ingress controller.

    $ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml

Then install the Ingress.

    $ kubectl apply -f ops/kubernetes/ingress.yaml

Add demo.net as test host to `/etc/hosts` reroute traffic to demo.net to localhost instead:

```
...
127.0.0.1   demo.net
...
```

## Run

Call the pods:

```
    $ curl demo.net/greet
```

Either the ALPHA or BRAVO pod should greet you.