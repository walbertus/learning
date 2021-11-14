#!/bin/bash
set -e

kubectl delete -f details-service.yml
kubectl delete -f productpage-service.yml
kubectl delete -f ratings-service.yml
kubectl delete -f reviews-service.yml
kubectl delete -f reviews-v1-service.yml
kubectl delete -f reviews-v2-service.yml
kubectl delete -f reviews-v3-service.yml
