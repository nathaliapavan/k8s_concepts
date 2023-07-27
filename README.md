# k8s_concepts

### executar a aplicação sem o docker

```bash
go run server.go
```

### criar docker build

```bash
docker build -t npavans/hello-go .
```

### executar imagem

```bash
docker run --rm -p 8080:8080 npavans/hello-go
```

### publicar imagem

```bash
docker push npavans/hello-go
```

### criar um cluster genérico

```bash
kind create cluster
```

### criar um cluster a partir de um arquivo de configuração yaml

```bash
kind create cluster --config=kind.yaml --name=CLUSTER_NAME
```

### listar configuração dos clusters disponíveis

```bash
kubectl config get-clusters
```

### listar os clusters disponíveis

```bash
kind get clusters
```

### deletar um cluster pelo nome

```bash
kind delete clusters CLUSTER_NAME
```

### acessar contexto do cluster

```bash
kubectl cluster-info --context kind-CLUSTER_NAME
```

### listar nodes do cluster atual

```bash
kubectl get nodes
```

### aplicar configuração a partir de um arquivo (-f file) yaml

```bash
kubectl apply -f pod.yaml
```

### ver detalhes do pod

```bash
kubectl describe pod POD_NAME
```

### listar pods

```bash
kubectl get pods
```

### deletar pod pelo nome

```bash
kubectl delete pod POD_NAME --force
```

### redirecionar porta do pod para acessar

```bash
kubectl port-forward pod/goserver 8000:8080
# acessar em http://localhost:8000/hello
```

### listar replicasets

```bash
kubectl get replicasets
```

### listar replicasets

```bash
kubectl delete replicaset REPLICASET_NAME
```

### listar deployments

```bash
kubectl get deployments
```

### ver detalhes do deployment

```bash
kubectl describe deployment DEPLOYMENT_NAME
```

### visualizar o histórico de deployments

```bash
kubectl rollout history deployment DEPLOYMENT_NAME
# output
# REVISION  CHANGE-CAUSE
# 1         <none>
```

### voltar para a versão anterior do deployment

```bash
kubectl rollout undo deployment DEPLOYMENT_NAME
```

### voltar para a revisão específica do deployment

```bash
kubectl rollout undo deployment DEPLOYMENT_NAME --to-revision=REVISION_NUMBER
```

### listar services

```bash
kubectl get svc
```

### redirecionar porta do service para acessar

```bash
kubectl port-forward svc/SERVICE_NAME 8000:8080
# acessar em http://localhost:8000/hello
```

### acessar api kubernetes

```bash
kubectl proxy --port=8001
# acessar em http://localhost:8001/api/v1/pods
```