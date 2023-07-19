package server

import (
	"com.ai.bff-purchase-order-inquiry/graph"
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"log"
	"os"
)

var (
	serviceName = os.Getenv("SERVICE_NAME")
)

// StartGraphQLServer starts the GraphQL server
func StartGraphQLServer(engine *gin.Engine) {

	// initial tracer
	initTracer()
	engine.Use(otelgin.Middleware(serviceName))
	// Create a new GraphQL server
	// Serve the GraphQL API endpoint
	engine.POST("/query", graphqlHandler())
	// Serve the GraphQL Playground (optional)
	engine.GET("/playground", playgroundHandler())
	// Start the server on port 8080
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("GraphQL server started on http://localhost:8080")
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func initTracer() func(context.Context) error {

	exporter, err := stdouttrace.New()

	if err != nil {
		log.Fatal(err)
	}
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Printf("Could not set resources: ", err)
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(resources),
		),
	)
	return exporter.Shutdown
}
