package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	sqlConnector, err := createSqlConnector(os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Panic().Msg(err.Error())
	}

	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.GET("", createHandler(createMoviesGetter(sqlConnector)))
	if err := g.Run(":5000"); err != nil {
		log.Error().Msg(err.Error())
	}
}
