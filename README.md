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
