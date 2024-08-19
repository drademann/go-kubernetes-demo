# Go Kubernetes Demo

## The Greeter

The Pod includes a simple Go program starting a server on port 8080.
The server takes the first argument as its name and answers requests
to the URL /greet with

    This is <name> responding!

You can test this by running

    $ go run ./greeter/main.go Tweety
    2024/08/19 09:45:14 starting Tweety Server on 8080

and requesting a response in another shell:

    $ curl localhost:8080/greet
    This is Tweety responding!

## The Docker

### Registry

Kubernetes needs a registry to pull Docker images from.
Install a local registry by running

    $ docker run -d -p 5050:5000 --name registry registry:latest

The registry is now accessible at localhost:5050 within Docker Desktop.

### Image

The Dockerfile builds the Greeter image.

    $ docker --debug build -t greeter .

    $ docker tag greeter localhost:5050/greeter

It is pushed to the local registry by calling

    $ docker push localhost:5050/greeter

## The Kubernetes

Assuming that you have your local Kubernetes (Docker Desktop) running,
we first deploy the Kubernetes resources by hand with `kubectl`.

### Namespace

Create the namespace `demo-space` as a _folder_ for the resources to come.

    $ kubectl apply -f ops/kubernetes/namespace.yaml

### Pods

Tell Kubernetes to create the pods.

    $ kubectl apply -f ops/kubernetes/greeter-alpha-pod.yaml
    $ kubectl apply -f ops/kubernetes/greeter-bravo-pod.yaml

### Services

Then tell Kubernetes to create a service.

    $ kubectl apply -f ops/kubernetes/greeter-service.yaml

### Ingress

If not already there, install a nginx ingress controller.

    $ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml

After giving the controller some time to boot, install the Ingress.

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

## Useful additional tools

- `stern` to see logs of pods combined
- `k9s` to see pods in a more convenient way
