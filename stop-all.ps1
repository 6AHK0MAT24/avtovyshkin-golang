# Script to stop all services

# Set window title
$Host.UI.RawUI.WindowTitle = "Stop All Services"

Write-Host "Stopping all services..." -ForegroundColor Yellow
# Stop processes on ports 8082, 8081 and 5173
$ports = @("8082", "8081", "5173")
foreach ($port in $ports) {
    $portInUse = netstat -ano | Select-String ":$port" | Select-String "LISTENING"
    if ($portInUse) {
        Write-Host "Stopping process on port $port..." -ForegroundColor Cyan
        $pid = ($portInUse -split '\s+')[-1]
        try {
            taskkill /F /PID $pid | Out-Null
            Write-Host "Process on port $port stopped." -ForegroundColor Green
        } catch {
            Write-Host "Failed to stop process on port $port." -ForegroundColor Red
        }
    } else {
        Write-Host "Port $port is not in use." -ForegroundColor Gray
    }
}

Write-Host ""
Write-Host "All services stopped." -ForegroundColor Green
