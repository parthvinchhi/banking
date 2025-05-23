package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parthvinchhi/bank-app/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomer(c *gin.Context) {
	status := c.Query("status")

	customers, err := ch.service.GetAllCustomer(status)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, customers)
}

func (ch *CustomerHandlers) getCustomer(c *gin.Context) {
	id := c.Param("customer_id")

	if customer, err := ch.service.GetCustomer(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, customer)
	}
}
