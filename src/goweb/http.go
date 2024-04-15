package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func createHandler(getMovies moviesGetter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		movies, err := getMovies()
		if err != nil {
			log.Error().Msg(err.Error())
			ctx.AbortWithError(500, err)
			return
		}

		ctx.JSON(200, movies)
	}
}
