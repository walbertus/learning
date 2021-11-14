#!/bin/bash
set -e

kubectl apply -f details-service.yml
kubectl apply -f productpage-service.yml
kubectl apply -f ratings-service.yml
kubectl apply -f reviews-service.yml
kubectl apply -f reviews-v1-service.yml
kubectl apply -f reviews-v2-service.yml
kubectl apply -f reviews-v3-service.yml
