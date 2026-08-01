@echo off
setlocal
set PATH=C:\Program Files\w64devkit\bin;%PATH%

set PORT=8080
if not "%~1"=="" set PORT=%~1

echo ========================================================
echo Iniciando Servidor ntfy na porta %PORT%...
echo Interface Web disponível em: http://localhost:%PORT%/app.html
echo API de Notificação em:       http://localhost:%PORT%/
echo ========================================================

start http://localhost:%PORT%/app.html
".\ntfy-app.exe" serve --listen-http ":%PORT%"
