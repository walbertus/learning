#!/bin/bash
set -e

helm install traefik-mesh traefik-mesh/traefik-mesh --values values.yaml -n traefik-mesh
