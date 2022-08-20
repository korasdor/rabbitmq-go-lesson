package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/korasdor/rabbitmq-go/producer/app"
	"github.com/korasdor/rabbitmq-go/producer/utils"
	"github.com/rs/zerolog/log"
)

func init() {
	mode := utils.GetEnvVar("GIN_MODE")
	gin.SetMode(mode)
}

func main() {
	app := app.SetupApp()

	addr := utils.GetEnvVar("GIN_ADDR")
	port := utils.GetEnvVar("GIN_PORT")
	https := utils.GetEnvVar("GIN_HTTPS")

	if https == "true" {
		certFile := utils.GetEnvVar("GIN_CERT")
		certKey := utils.GetEnvVar("GIN_CERT_KEY")

		log.Info().Msgf("Starting service on https://%s:%s", addr, port)

		if err := app.RunTLS(fmt.Sprintf("%s-%s", addr, port), certFile, certKey); err != nil {
			log.Fatal().Err(err).Msg("Error occurred while setting up the server in HTTPS mode")
		}
	}

	log.Info().Msgf("Starting service on http://%s:%s", addr, port)

	if err := app.Run(fmt.Sprintf("%s:%s", addr, port)); err != nil {
		log.Fatal().Err(err).Msg("Error occurred while setting up the server")
	}
}
