# 2. The Helm Demo

Instead of handling all the YAML files describing the resources for
Kubernetes, we instead use a so-called Helm Chart.
Before doing so, delete all currently deployed resources: deployments, service, ingress.
Deletion of the deployments will delete the replica sets and their pods as well.

## Introduction

Once set up, helm lets you deploy many of the resources previously deployed one by one
with `kubectl`, now with a single `helm install`.
The core of Helm is a huge template engine.
And at its heart, it is Go's template engine.

## The Namespace

The namespace usually has a different lifecycle than the other Kubernetes resources.
Therefore, we use the existing namespace 'greeter-space' for the Helm deployments.

## The Deployments

First, have a look at `ops/helm/templates/deployment.yaml`.
When called by Helm, this template will loop over the _greeters_ array defined in `values.yaml`.
Each iteration creates a deployment.
Try it by calling

    $ helm install greeter ./ops/helm --debug --dry-run

It won't install anything yet, it is a dry run.
The output shows the resource YAML files content, now created by the template engine.
If we want another `gamma` deployment, we can add it to the array in `values.yaml`.

## The Service

The service template `ops/helm/templates/service.yaml` is even simpler:
It uses the pod's port definition from the `values.yaml`,
and finds its pods by their role, as before.

## The Ingress

Same as with the service. The template engine will replace the values so Helm can apply
the Ingress resource.

## Installing and Uninstalling

Install the Helm chart into your local Kubernetes with

    $ helm install greeter ./ops/helm

and uninstall it using

    $ helm uninstall greeter

---

## Summary

Now, this is a glimpse only of what Helm can do within the templates.
It provides many functions to form the output for the kubernetes resources as you need them.

See the good documentation of Helm at
the [Chart Template Guide](https://helm.sh/docs/chart_template_guide/getting_started/)

Last tipp: Another useful tool for viewing the current state of your Kubernetes cluster is `k9s`.