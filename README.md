<!-- ABOUT THE PROJECT -->
## About The Project

This is a project template for golang fiber microservice with hexagonal architect support GORM and aws lambda serverless

<!-- GETTING STARTED -->
## Getting Started

### Installation
1. https://github.com/gofiber/fiber
2. https://github.com/swaggo/swag
3. https://github.com/golang-migrate/migrate
4. https://github.com/go-gorm/gorm
5. https://github.com/awslabs/aws-lambda-go-api-proxy

<!-- USAGE EXAMPLES -->
## Usage
1. Generate docs
  - swag init
2. Start mysql & migrates data
  - make docker.mysql
  - make migrate.up
3. Start server
  - go run main.go
4. (Optinal) build for lambda
  - Modify main.go uncoment lambda.Start(app.Handler) and comment app.StartServer()
  - make docker.fiber.build
