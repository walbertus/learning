# Traefik Mesh Installation

## Requirements

- Helm v3 ([install page](https://helm.sh/docs/intro/install/))
  <br>I recommend using [asdf](http://asdf-vm.com/guide/getting-started.html) for installation
  <br>asdf plugin for helm can be found [here](https://github.com/Antiarchitect/asdf-helm)
- Execute `./01-prepare-helm-chart.sh` to add traefik-mesh helm repository

## Install
1. Create `traefik-mesh` namespace using this command
   <br>`kubectl create ns traefik-mesh`
1. Execute `./02-install-traefik-mesh.sh` to install traefik mesh to your Kubernetes cluster.
   <br>Use [values.yaml](./values.yaml) to customize the installation

## Check installation
Use this command to view the generated resources
```shell
kubectl get all -n traefik-mesh
```

## Onboard application to use traefik
You can config your application to use traefik mesh by changing the communication between service from using `.svc.cluster.local` to using `.traefik.mesh`

## Upgrade installation
Change configuration in [values.yaml](./values.yaml) then run this command to upgrade the installation.
`./03-upgrade-traefik-mesh.yml`

## Uninstall
Execute `./04-remove-traefik-mesh.sh` to remove traefik-mesh installation.
