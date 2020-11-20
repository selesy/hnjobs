#!/bin/bash

minikube start --cpus=2 --memory=4G --addons=ingress
eval $(minikube -p minikube docker-env)

kubectl apply -f k8s/minikube/coredns.yaml
kubectl apply -f k8s/minikube/rook-cockroachdb-operator.yaml
