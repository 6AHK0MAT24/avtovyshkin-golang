# Generate driver photos script
$ErrorActionPreference = "Stop"

$photoDir = "uploads/photo"
$dbName = "avtovyshkin_db_go"

Write-Host "=== Generating driver photos ===" -ForegroundColor Green

if (-not (Test-Path $photoDir)) {
    New-Item -ItemType Directory -Path $photoDir -Force | Out-Null
    Write-Host "Created directory: $photoDir" -ForegroundColor Yellow
}

Write-Host "`nFetching drivers from database..." -ForegroundColor Yellow

$query = "SELECT fldid, fldfirstname, fldlastname FROM drivers ORDER BY fldlastname"

$drivers = docker exec -i avtovyshkin-postgres psql -U postgres -d $dbName -t -A -F "|" -c $query

if (-not $drivers) {
    Write-Host "No drivers found!" -ForegroundColor Red
    exit 1
}

$driverList = $drivers -split "`n" | Where-Object { $_ -ne "" }
Write-Host "Found drivers: $($driverList.Count)" -ForegroundColor Green

$counter = 1
foreach ($driver in $driverList) {
    $parts = $driver -split "\|"
    $driverId = $parts[0]
    $firstName = $parts[1]
    $lastName = $parts[2]

    Write-Host "`n[$counter/$($driverList.Count)] Processing: $lastName $firstName (ID: $driverId)" -ForegroundColor Cyan

    $fileName = "driver_${driverId}.jpg"
    $filePath = Join-Path $photoDir $fileName
    $photoUrl = "/uploads/photo/$fileName"

    if (Test-Path $filePath) {
        Write-Host "  Photo already exists: $fileName" -ForegroundColor Gray
    } else {
        # Generate random color
        $randomColor = Get-Random -Minimum 0 -Maximum 16777215
        $colorHex = "{0:X6}" -f $randomColor

        # Create simple colored image using PowerShell
        try {
            Write-Host "  Creating image..." -ForegroundColor Gray

            # Create a simple 200x200 bitmap with random color
            $bitmap = New-Object System.Drawing.Bitmap 200, 200
            $graphics = [System.Drawing.Graphics]::FromImage($bitmap)
            $color = [System.Drawing.Color]::FromArgb([Convert]::ToInt32($colorHex, 16))
            $graphics.Clear($color)

            # Add initials
            $font = New-Object System.Drawing.Font("Arial", 60)
            $brush = New-Object System.Drawing.SolidBrush([System.Drawing.Color]::White)
            $initials = "$($firstName[0])$($lastName[0])"
            $stringSize = $graphics.MeasureString($initials, $font)
            $x = (200 - $stringSize.Width) / 2
            $y = (200 - $stringSize.Height) / 2
            $graphics.DrawString($initials, $font, $brush, $x, $y)

            $bitmap.Save($filePath, [System.Drawing.Imaging.ImageFormat]::Jpeg)
            $graphics.Dispose()
            $bitmap.Dispose()

            Write-Host "  Saved: $fileName" -ForegroundColor Green
        } catch {
            Write-Host "  Error creating image: $_" -ForegroundColor Red
            continue
        }
    }

    $updateQuery = "UPDATE drivers SET fldphoto = '$photoUrl' WHERE fldid = '$driverId'"

    try {
        docker exec -i avtovyshkin-postgres psql -U postgres -d $dbName -c $updateQuery | Out-Null
        Write-Host "  Database updated" -ForegroundColor Green
    } catch {
        Write-Host "  DB update error: $_" -ForegroundColor Red
    }

    $counter++
}

Write-Host "`n=== Generation complete! ===" -ForegroundColor Green
Write-Host "Total drivers processed: $($driverList.Count)" -ForegroundColor Green
Write-Host "Photos saved in: $photoDir" -ForegroundColor Green
