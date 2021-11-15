#!/bin/bash
set -e

kubectl apply -f reviews-traffic-split.yml
kubectl apply -f productpage-cm-traefik.yml
kubectl rollout restart deployment productpage-v1
