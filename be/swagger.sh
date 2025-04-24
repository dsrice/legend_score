#!/bin/bash

# Install swag if not already installed
if ! command -v swag &> /dev/null
then
    echo "Installing swag..."
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Generate Swagger documentation
echo "Generating Swagger documentation..."
swag init -g swagger.go -o docs

echo "Swagger documentation generated successfully!"