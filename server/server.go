package server

import (
	"com.ai.bff-purchase-order-inquiry/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"log"
)

// StartGraphQLServer starts the GraphQL server
func StartGraphQLServer(engine *gin.Engine) {
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
