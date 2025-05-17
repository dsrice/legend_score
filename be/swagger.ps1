# Check if swag is installed
$swagInstalled = $null
try {
    $swagInstalled = Get-Command swag -ErrorAction SilentlyContinue
} catch {
    # Command not found
}

# Install swag if not already installed
if ($null -eq $swagInstalled) {
    Write-Host "Installing swag..."
    go install github.com/swaggo/swag/cmd/swag@latest
}

# Generate Swagger documentation
Write-Host "Generating Swagger documentation..."
swag init -g swagger.go -o docs

Write-Host "Swagger documentation generated successfully!"