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
		response.Errors.ErrorCode = randomdata.StringSample("400", "500")
		response.Errors.ErrorMessage = "Error"
		return response
	} else {
		response.OrderId = randomdata.StringNumberExt(2, "-", 5)
		response.DeliveryTime = randomdata.StringSample("1 day", "2 days", "3 days")
	}

	return response
}
