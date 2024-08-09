# Go Kubernetes Demo

## The Registry

    $ docker run -d -p 5050:5000 --name registry registry:latest

Local registry accessible at localhost:5050 within Docker Desktop.

## The Image

    $ docker build --debug -t greeter .

    $ docker tag greeter localhost:5050/greeter

    $ docker push localhost:5050/greeter