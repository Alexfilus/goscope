// Copyright © 2020 Pro Warehouse B.V.
// All Rights Reserved
package goscope

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Setup(engine *gin.Engine) {
	logger := &LoggerGoScope{}
	gin.DefaultErrorWriter = logger
	log.SetFlags(0)
	log.SetOutput(logger)
	// Use the logging middleware
	engine.Use(ResponseLogger)
	// Setup necessary routes
	goscopeGroup := engine.Group("/goscope")
	goscopeGroup.GET("/", Dashboard)
	goscopeGroup.GET("/logs", LogDashboard)
	goscopeGroup.GET("/log-records", GetLogs)
	goscopeGroup.GET("/log-records/:id", ShowLog)
	goscopeGroup.GET("/requests", GetRequests)
	goscopeGroup.GET("/requests/:id", ShowRequest)
}
