package app

import (
	"context"
	"dont/hexagonal/app/routes"
	"dont/hexagonal/pkg/configs"
	"dont/hexagonal/pkg/middleware"
	"dont/hexagonal/utils"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/aws/aws-lambda-go/events"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	_ "github.com/joho/godotenv/autoload"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.ItemRoutes(app)
	routes.SwaggerRoute(app)
	routes.NotFoundRoute(app)
	fiberLambda = fiberadapter.New(app)
	// Start server (with or without graceful shutdown).
	// if os.Getenv("STAGE_STATUS") == "dev" {
	// 	utils.StartServer(app)
	// } else {
	// 	utils.StartServerWithGracefulShutdown(app)
	// }
}

func StartServer() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.ItemRoutes(app)
	routes.SwaggerRoute(app)
	routes.NotFoundRoute(app)
	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}

// Handler will deal with Fiber working with Lambda
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return fiberLambda.ProxyWithContext(ctx, req)
}
