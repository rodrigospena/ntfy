# Script de Inicialização do ntfy
param (
    [int]$Port = 8080
)

$env:PATH = "C:\Program Files\w64devkit\bin;" + $env:PATH

Write-Host "========================================================" -ForegroundColor Cyan
Write-Host "Iniciando Servidor ntfy na porta $Port..." -ForegroundColor Green
Write-Host "Interface Web: http://localhost:$Port/app.html" -ForegroundColor Yellow
Write-Host "API Endpoint:  http://localhost:$Port/" -ForegroundColor Yellow
Write-Host "========================================================" -ForegroundColor Cyan

Start-Process "http://localhost:$Port/app.html"
& ".\ntfy-app.exe" serve --listen-http ":$Port"
