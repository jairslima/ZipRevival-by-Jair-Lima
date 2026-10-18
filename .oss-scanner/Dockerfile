# ZipRevival by Jair Lima: build para o OSS Scanner (https://github.com/anthropics/oss-scanner).
# A rede só existe durante este build: `go mod download` traz todas as dependências para a imagem.
FROM golang:1.26-bookworm

ENV CGO_ENABLED=0

# contrato do scanner: o checkout fica em /src
COPY . /src
WORKDIR /src

RUN go mod download

# Compila os 10 utilitários (zip, unzip, pkzip, pkunzip, unrar, untar, un7z, unlha, uniso, uncab) em /src/bin,
# com símbolos de depuração (sem -s -w) e sem otimizações que atrapalhem a análise.
RUN go build -gcflags="all=-N -l" -o /src/bin/ ./cmd/...

# Testes rápidos: não deixam o build falhar.
RUN go vet ./... || echo "WARNING: go vet reportou problemas"; \
    go test ./... || echo "WARNING: go test reportou falhas"
