#!/bin/bash
set -e

kubectl apply -f reviews-traffic-split.yml
kubectl rollout restart deployment productpage-v1
