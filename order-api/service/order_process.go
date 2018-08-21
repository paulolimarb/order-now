package service

import (
	"order-now/order-api/models"
	"order-now/order-api/queue"

	"io"

	"io/ioutil"

	"github.com/Pallinder/go-randomdata"
)

func ProcessOrder(request io.ReadCloser) models.OrderResponse {

	response := models.OrderResponse{}

	r, _ := ioutil.ReadAll(request)

	err := queue.WriteMessage(string(r))

	if err != nil {
		response.Error.ErrorCode = string(randomdata.Number(100, 500))
		response.Error.ErrorMessage = err.Error()
		return response
	} else {
		response.OrderId = string(randomdata.Number(100, 999))
		response.DeliveryTime = string(randomdata.Number(1, 10))
	}

	return response
}
