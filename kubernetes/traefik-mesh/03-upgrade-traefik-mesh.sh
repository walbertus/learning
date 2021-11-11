#!/bin/bash
set -e

helm upgrade traefik-mesh traefik-mesh/traefik-mesh --values values.yaml -n traefik-mesh
