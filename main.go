package main

import (
	"com.ai.bff-purchase-order-inquiry/server"
	"github.com/gin-gonic/gin"
)

func main() {
	// Start the GraphQL server
	engine := gin.Default()
	server.StartGraphQLServer(engine)
}
