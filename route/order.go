package route

import (
	"second-assignment/controller"

	"github.com/gin-gonic/gin"
)

type orderRouter struct {
	orderController controller.OrderController
}

func NewOrderController (oc controller.OrderController) *orderRouter  {
	return &orderRouter{
		orderController: oc,
	}

}

func (or orderRouter) OrderRouter (r *gin.Engine) {
	r.POST("/order", or.orderController.Create)
	r.GET("/order", or.orderController.GetAll)
	r.DELETE("/order/:uuid", or.orderController.DeleteByUUID)
	r.PUT("/order", or.orderController.Update)
}