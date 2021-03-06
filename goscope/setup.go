// License: MIT
// Authors:
// 		- Josep Bigorra (averageflow)
package goscope

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	// Import MYSQL Driver
	_ "github.com/go-sql-driver/mysql"
	// Import PostgreSQL Driver
	_ "github.com/lib/pq"
)

// Ensure necessary application variables are set.
func CheckVariablesAreSet() {
	variables := []string{
		"APPLICATION_ID",
		"APPLICATION_NAME",
		"APPLICATION_TIMEZONE",
		"GOSCOPE_DATABASE_CONNECTION",
		"GOSCOPE_DATABASE_TYPE",
		"GOSCOPE_ENTRIES_PER_PAGE",
	}
	for _, s := range variables {
		if os.Getenv(s) == "" {
			panic(fmt.Sprintf("%s variable is not set", s))
		}
	}
}

func Setup(engine *gin.Engine) {
	CheckVariablesAreSet()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	logger := &LoggerGoScope{}
	gin.DefaultErrorWriter = logger

	log.SetFlags(log.Lshortfile)
	log.SetOutput(logger)
	// Use the logging middleware
	engine.Use(ResponseLogger)
	// Setup necessary routes
	goscopeGroup := engine.Group("/goscope")
	goscopeGroup.GET("/", RequestDashboard)
	goscopeGroup.GET("/logs", LogDashboard)
	goscopeGroup.GET("/info", ShowSystemInfo)
	goscopeGroup.GET("/log-records", GetLogs)
	goscopeGroup.GET("/log-records/:id", ShowLog)
	goscopeGroup.GET("/requests", GetRequests)
	goscopeGroup.GET("/requests/:id", ShowRequest)
	goscopeGroup.POST("/search/requests", SearchRequest)
	goscopeGroup.POST("/search/logs", SearchLog)
}
