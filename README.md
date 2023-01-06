# kubectl-cluster
kubectl-cluster is a tool to display/switch between Kubernetes clusters (and configure them for kubectl) easily.

## Examples
```
# list clusters
kubectl cluster
  or
kubectl cluster list

# show current cluster
kubectl cluster current

# swith context with cluster name
kubectl cluster use <cluster>
  or
kubectl cluster switch <cluster>

# show usages
kubectl cluster help

# completion
kubectl cluster completion <shell>
```

## Installation
### with go install
```
go install github.com/sp-yduck/kubectl-cluster@latest
```

## How it works
The relationship between cluster and context is "one to many". kubectl-cluster uses the newest context which using specified cluster.

for example, if your kubeconfig looks like following and you chose `cluster==docker-desktop` , kubectl-cluster will switch your current-context to `kubernetes-admin@docker-desktop` as this is the newest context in this kubeconfig.
```
clusters:
- cluster:
  name: docker-desktop
- cluster:
  name: kubernetes
- cluster:
  name: rancher-desktop
contexts:
- context:                      # here is first context using cluster:docker-desktop
    cluster: docker-desktop
  name: docker-desktop
- context:
    cluster: kubernetes
  name: kubernetes-admin@kubernetes
- context:
    cluster: rancher-desktop
  name: rancher-desktop
- context:                      # here is newest context using cluster:docker-desktop
    cluster: docker-desktop
    user: admin
  name: kubernetes-admin@docker-desktop
current-context: docker-desktop
```

## Roadmap
- [ ] krew installation
- [ ] detailed information for list command
- [ ] fuzzy switching
