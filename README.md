# Boatswain - Captain assistant
[Captain](https://github.com/alauda/captain) is a Helm 3 controller. It allows installing Helm charts from withing K8s clusters.

Boatswain is an external client for Captain, which exposes Rest API for charts CRUD operations.

## Design
- uses [go-client](https://github.com/kubernetes/client-go) for k8s
- uses [captain-crds](https://github.com/alauda/helm-crds) module 
- Exposes POST/DELETE Rest API endpoints to install/delete charts
- PoC scope

## Motivation
Experimental DIY for programmatic Helm charts administration. 
