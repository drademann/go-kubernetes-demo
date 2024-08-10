# Go Kubernetes Demo

## The Docker

### Registry

Run a local registry accessible at localhost:5050 within Docker Desktop.

    $ docker run -d -p 5050:5000 --name registry registry:latest

### Image

Build the Greeter image.

    $ docker build --debug -t greeter .

    $ docker tag greeter localhost:5050/greeter

    $ docker push localhost:5050/greeter

## The Kubernetes

### Namespace

Create the namespace `demo-space` as a _folder_ for the resources to come.

    $ kubectl apply -f ops/kubernetes/namespace.yaml



### Pods

Tell kubernetes to create the first pod.

    $ kubectl apply -f ops/kubernetes/pods/greeter-alpha.yaml
