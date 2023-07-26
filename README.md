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
docker push npavan/hello-go
```
