# Rodando o Servidor com Docker

## Build da imagem

```bash
docker build -t banco-api .
```

---

## Executar o container

```bash
docker run -p 8080:8080 banco-api
```

A API ficará disponível em:

```text
http://localhost:8080
```
