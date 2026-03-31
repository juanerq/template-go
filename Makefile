swagger: ## Genera documentación Swagger
	@echo "Generando documentación Swagger..."
	swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal