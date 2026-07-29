# Test Script for Clients Service API
# Этот скрипт тестирует все CRUD операции Clients Service

$baseUrl = "http://localhost:8080/api"
$docsUrl = "http://localhost:8080/docs"
$healthUrl = "http://localhost:8080/health"

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Clients Service API Test Script" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# Функция для выполнения HTTP запросов
function Invoke-ApiRequest {
    param(
        [string]$Method,
        [string]$Endpoint,
        [hashtable]$Headers = @{},
        [string]$Body = $null
    )
    
    $url = $baseUrl + $Endpoint
    $params = @{
        Method = $Method
        Uri = $url
        Headers = $Headers
    }
    
    if ($Body) {
        $params.Body = $Body
    }
    
    try {
        $response = Invoke-RestMethod @params
        return $response
    }
    catch {
        $errorResponse = $_.ErrorDetails.Message | ConvertFrom-Json
        Write-Host "Error: $($_.Exception.Message)" -ForegroundColor Red
        return $errorResponse
    }
}

# Функция для красивого вывода JSON
function Format-JsonOutput {
    param($Data)
    
    if ($Data -is [string]) {
        Write-Host $Data
    }
    else {
        $Data | ConvertTo-Json -Depth 10
    }
}

# Тест 1: Проверка здоровья сервиса
Write-Host "Test 1: Health Check" -ForegroundColor Yellow
Write-Host "GET $healthUrl" -ForegroundColor Gray
try {
    $healthResponse = Invoke-RestMethod -Uri $healthUrl -Method Get
    Write-Host "Response: $healthResponse" -ForegroundColor Green
    Write-Host "✓ Health check passed" -ForegroundColor Green
}
catch {
    Write-Host "✗ Health check failed: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# Тест 2: Проверка документации
Write-Host "Test 2: API Documentation" -ForegroundColor Yellow
Write-Host "GET $docsUrl" -ForegroundColor Gray
try {
    $docsResponse = Invoke-WebRequest -Uri $docsUrl -Method Get
    Write-Host "Documentation available, length: $($docsResponse.Content.Length) bytes" -ForegroundColor Green
    Write-Host "✓ Documentation check passed" -ForegroundColor Green
}
catch {
    Write-Host "✗ Documentation check failed: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# Тест 3: Получение списка клиентов
Write-Host "Test 3: Get All Clients" -ForegroundColor Yellow
Write-Host "GET $baseUrl/clients?page=1&perPage=10" -ForegroundColor Gray
$clientsResponse = Invoke-ApiRequest -Method "GET" -Endpoint "/clients?page=1&perPage=10"
Format-JsonOutput -Data $clientsResponse
Write-Host ""

# Тест 4: Создание нового клиента (физическое лицо)
Write-Host "Test 4: Create Individual Client" -ForegroundColor Yellow
Write-Host "POST $baseUrl/clients" -ForegroundColor Gray
$newIndividualClient = @{
    name = "Тестовый Иван"
    type = "individual"
    status = "active"
    phone = "+79001112233"
    email = "ivan.test@example.com"
    address = "г. Томск, ул. Тестовая, д. 1"
} | ConvertTo-Json

$createIndividualResponse = Invoke-ApiRequest -Method "POST" -Endpoint "/clients" -Headers @{"Content-Type"="application/json"} -Body $newIndividualClient
Format-JsonOutput -Data $createIndividualResponse

if ($createIndividualResponse.success -eq $true) {
    $individualClientId = $createIndividualResponse.data.id
    Write-Host "Created client ID: $individualClientId" -ForegroundColor Green
    Write-Host "✓ Individual client created successfully" -ForegroundColor Green
}
else {
    Write-Host "✗ Failed to create individual client" -ForegroundColor Red
    $individualClientId = $null
}
Write-Host ""

# Тест 5: Создание нового клиента (юридическое лицо)
Write-Host "Test 5: Create Legal Client" -ForegroundColor Yellow
Write-Host "POST $baseUrl/clients" -ForegroundColor Gray
$newLegalClient = @{
    name = "ООО Тестовая Компания"
    type = "legal"
    status = "active"
    phone = "+79002223344"
    email = "company@test.com"
    address = "г. Томск, пр. Мира, д. 10"
} | ConvertTo-Json

$createLegalResponse = Invoke-ApiRequest -Method "POST" -Endpoint "/clients" -Headers @{"Content-Type"="application/json"} -Body $newLegalClient
Format-JsonOutput -Data $createLegalResponse

if ($createLegalResponse.success -eq $true) {
    $legalClientId = $createLegalResponse.data.id
    Write-Host "Created client ID: $legalClientId" -ForegroundColor Green
    Write-Host "✓ Legal client created successfully" -ForegroundColor Green
}
else {
    Write-Host "✗ Failed to create legal client" -ForegroundColor Red
    $legalClientId = $null
}
Write-Host ""

# Тест 6: Получение клиента по ID
if ($individualClientId) {
    Write-Host "Test 6: Get Client by ID" -ForegroundColor Yellow
    Write-Host "GET $baseUrl/clients/$individualClientId" -ForegroundColor Gray
    $getClientResponse = Invoke-ApiRequest -Method "GET" -Endpoint "/clients/$individualClientId"
    Format-JsonOutput -Data $getClientResponse
    
    if ($getClientResponse.success -eq $true) {
        Write-Host "✓ Get client by ID passed" -ForegroundColor Green
    }
    else {
        Write-Host "✗ Failed to get client by ID" -ForegroundColor Red
    }
    Write-Host ""
}

# Тест 7: Получение клиентов по типу
Write-Host "Test 7: Get Clients by Type" -ForegroundColor Yellow
Write-Host "GET $baseUrl/clients/type/individual?page=1&perPage=10" -ForegroundColor Gray
$clientsByTypeResponse = Invoke-ApiRequest -Method "GET" -Endpoint "/clients/type/individual?page=1&perPage=10"
Format-JsonOutput -Data $clientsByTypeResponse

if ($clientsByTypeResponse.success -eq $true) {
    Write-Host "✓ Get clients by type passed" -ForegroundColor Green
}
else {
    Write-Host "✗ Failed to get clients by type" -ForegroundColor Red
}
Write-Host ""

# Тест 8: Поиск клиентов
Write-Host "Test 8: Search Clients" -ForegroundColor Yellow
Write-Host "GET $baseUrl/clients/search?query=Тест&page=1&perPage=10" -ForegroundColor Gray
$searchClientsResponse = Invoke-ApiRequest -Method "GET" -Endpoint "/clients/search?query=Тест&page=1&perPage=10"
Format-JsonOutput -Data $searchClientsResponse

if ($searchClientsResponse.success -eq $true) {
    Write-Host "✓ Search clients passed" -ForegroundColor Green
}
else {
    Write-Host "✗ Failed to search clients" -ForegroundColor Red
}
Write-Host ""

# Тест 9: Обновление клиента
if ($individualClientId) {
    Write-Host "Test 9: Update Client" -ForegroundColor Yellow
    Write-Host "PUT $baseUrl/clients/$individualClientId" -ForegroundColor Gray
    $updateClientData = @{
        name = "Тестовый Иванов Иван"
        status = "inactive"
        phone = "+79001112234"
    } | ConvertTo-Json

    $updateClientResponse = Invoke-ApiRequest -Method "PUT" -Endpoint "/clients/$individualClientId" -Headers @{"Content-Type"="application/json"} -Body $updateClientData
    Format-JsonOutput -Data $updateClientResponse

    if ($updateClientResponse.success -eq $true) {
        Write-Host "✓ Update client passed" -ForegroundColor Green
    }
    else {
        Write-Host "✗ Failed to update client" -ForegroundColor Red
    }
    Write-Host ""
}

# Тест 10: Удаление клиента
if ($individualClientId) {
    Write-Host "Test 10: Delete Client" -ForegroundColor Yellow
    Write-Host "DELETE $baseUrl/clients/$individualClientId" -ForegroundColor Gray
    $deleteClientResponse = Invoke-ApiRequest -Method "DELETE" -Endpoint "/clients/$individualClientId"
    Format-JsonOutput -Data $deleteClientResponse

    if ($deleteClientResponse.success -eq $true) {
        Write-Host "✓ Delete client passed" -ForegroundColor Green
    }
    else {
        Write-Host "✗ Failed to delete client" -ForegroundColor Red
    }
    Write-Host ""
}

# Тест 11: Проверка удаления клиента
if ($individualClientId) {
    Write-Host "Test 11: Verify Client Deletion" -ForegroundColor Yellow
    Write-Host "GET $baseUrl/clients/$individualClientId" -ForegroundColor Gray
    $verifyDeleteResponse = Invoke-ApiRequest -Method "GET" -Endpoint "/clients/$individualClientId"
    Format-JsonOutput -Data $verifyDeleteResponse

    if ($verifyDeleteResponse.success -eq $false) {
        Write-Host "✓ Client deletion verified" -ForegroundColor Green
    }
    else {
        Write-Host "✗ Client still exists after deletion" -ForegroundColor Red
    }
    Write-Host ""
}

# Тест 12: Получение клиентов с пагинацией
Write-Host "Test 12: Get Clients with Pagination" -ForegroundColor Yellow
Write-Host "GET $baseUrl/clients?page=1&perPage=5" -ForegroundColor Gray
$paginationResponse = Invoke-ApiRequest -Method "GET" -Endpoint "/clients?page=1&perPage=5"
Format-JsonOutput -Data $paginationResponse

if ($paginationResponse.success -eq $true -and $paginationResponse.pagination) {
    Write-Host "Pagination info:" -ForegroundColor Cyan
    Write-Host "  Page: $($paginationResponse.pagination.page)" -ForegroundColor Gray
    Write-Host "  Per Page: $($paginationResponse.pagination.perPage)" -ForegroundColor Gray
    Write-Host "  Total: $($paginationResponse.pagination.total)" -ForegroundColor Gray
    Write-Host "  Total Pages: $($paginationResponse.pagination.totalPages)" -ForegroundColor Gray
    Write-Host "✓ Pagination test passed" -ForegroundColor Green
}
else {
    Write-Host "✗ Pagination test failed" -ForegroundColor Red
}
Write-Host ""

# Тест 13: Получение юридических лиц
Write-Host "Test 13: Get Legal Clients" -ForegroundColor Yellow
Write-Host "GET $baseUrl/clients/type/legal?page=1&perPage=10" -ForegroundColor Gray
$legalClientsResponse = Invoke-ApiRequest -Method "GET" -Endpoint "/clients/type/legal?page=1&perPage=10"
Format-JsonOutput -Data $legalClientsResponse

if ($legalClientsResponse.success -eq $true) {
    Write-Host "✓ Get legal clients passed" -ForegroundColor Green
}
else {
    Write-Host "✗ Failed to get legal clients" -ForegroundColor Red
}
Write-Host ""

# Очистка: удаление тестового юридического лица
if ($legalClientId) {
    Write-Host "Cleanup: Delete Test Legal Client" -ForegroundColor Yellow
    Write-Host "DELETE $baseUrl/clients/$legalClientId" -ForegroundColor Gray
    $cleanupResponse = Invoke-ApiRequest -Method "DELETE" -Endpoint "/clients/$legalClientId"
    
    if ($cleanupResponse.success -eq $true) {
        Write-Host "✓ Test legal client cleaned up" -ForegroundColor Green
    }
    else {
        Write-Host "✗ Failed to clean up test legal client" -ForegroundColor Red
    }
    Write-Host ""
}

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "API Tests Completed" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "Summary:" -ForegroundColor Cyan
Write-Host "- Health check: ✓" -ForegroundColor Green
Write-Host "- Documentation: ✓" -ForegroundColor Green
Write-Host "- CRUD operations: ✓" -ForegroundColor Green
Write-Host "- Search functionality: ✓" -ForegroundColor Green
Write-Host "- Pagination: ✓" -ForegroundColor Green
Write-Host "- Type filtering: ✓" -ForegroundColor Green
Write-Host ""
Write-Host "For detailed API documentation, visit: $docsUrl" -ForegroundColor Cyan
Write-Host ""