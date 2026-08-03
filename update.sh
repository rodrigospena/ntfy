#!/bin/bash
echo "=========================================="
echo "   Atualizando o Servidor ntfy-server     "
echo "=========================================="

echo "[1/4] Baixando a versão mais recente do GitHub..."
git pull origin main

echo "[2/4] Parando processo antigo do servidor..."
sudo pkill -f ntfy-server || true

echo "[3/4] Compilando novo ntfy-server..."
go build -o ntfy-server main.go

echo "[4/4] Iniciando o servidor ntfy em segundo plano..."
nohup ./ntfy-server serve --config server.yml > ntfy.log 2>&1 &

echo "=========================================="
echo "   ✅ Servidor Atualizado com Sucesso!     "
echo "=========================================="
