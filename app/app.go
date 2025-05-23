package app

import (
	"github.com/gin-gonic/gin"
	"github.com/parthvinchhi/bank-app/domain"
	"github.com/parthvinchhi/bank-app/service"
)

func Start() {
	r := gin.Default()

	// wiring
	ch := CustomerHandlers{
		// service: service.NewCustomerService(domain.NewCustomerRepositoryStub()),
		service: service.NewCustomerService(domain.NewCustomerRepositoryDB()),
	}

	r.GET("/customers", ch.getAllCustomer)
	r.GET("/customers/:customer_id", ch.getCustomer)

	r.Run(":8123")
}
