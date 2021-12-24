package main

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)
}

func main() {
	e := echo.New()
	e.Use(loggingMiddleware)

	e.GET("/isInt", func(c echo.Context) error {
		a := c.QueryParam("a")
		log.WithField("a", a).Debug("parsing a string")
		if _, err := strconv.Atoi(a); err != nil {
			log.WithField("a", a).Error("unable to parse the string")
			return c.String(http.StatusBadRequest, "not ok")
		}
		log.WithField("a", a).Debug("string is parsed")
		return c.String(http.StatusOK, "ok")
	})

	e.Start(":5050")
}

func loggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		res := next(c)
		log.WithFields(log.Fields{
			"methods":    c.Request().Method,
			"path":       c.Path(),
			"status":     c.Response().Status,
			"latency_ns": time.Since(start).Nanoseconds(),
		}).Info("request details")
		return res
	}
}
