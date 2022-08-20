package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korasdor/rabbitmq-go/producer/consts"
	"github.com/korasdor/rabbitmq-go/producer/models"
	"github.com/korasdor/rabbitmq-go/producer/utils"
	"github.com/rs/zerolog/log"
)

func Example(ctx *gin.Context) {
	var msg models.Message

	request_id := ctx.GetString("x-request-id")

	if err := ctx.ShouldBindJSON(&msg); err != nil {
		log.Error().Err(err).Str("request_id", request_id).Msg("Error occured while binding request data")

		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	connectionString := utils.GetEnvVar("RMQ_URL")
	rmqProducer := utils.RMQProducer{
		Queue:            consts.EXAMPLE_QUEUE,
		ConnectionString: connectionString,
	}

	rmqProducer.PublishMessage("text/plain", []byte(msg.Message))
}
