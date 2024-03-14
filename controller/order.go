package controller

import (
	"net/http"
	"second-assignment/model"
	"second-assignment/repository"
	"second-assignment/util"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderRepository repository.IOrderRepository
}

func NewOrderController(orderRepository repository.IOrderRepository) *OrderController {
	return &OrderController{
		orderRepository: orderRepository,
	}
}

func (or *OrderController) Create(ctx *gin.Context) {
	var newOrder model.Order

	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		res := util.CreateResponse(http.StatusBadRequest,false,err, nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	createdPerson, err := or.orderRepository.Create(newOrder)
	if err != nil {
		res := util.CreateResponse(http.StatusInternalServerError,false,err, nil,)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := util.CreateResponse(http.StatusCreated,true,nil, createdPerson,)
	ctx.JSON(http.StatusOK, res)
}

func (or *OrderController) GetAll (ctx *gin.Context) {
	persons, err := or.orderRepository.GetAll()
	if err != nil {
		res := util.CreateResponse(http.StatusInternalServerError, false, err, nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
	return
	}
	res := util.CreateResponse(http.StatusOK, true, nil, persons)
	ctx.JSON(http.StatusOK,res)
}

func (or *OrderController) DeleteByUUID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err :=or.orderRepository.Delete(uuid)
	if err != nil {
		res := util.CreateResponse(http.StatusInternalServerError, false, err, nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
	return
	} 
	res := util.CreateResponse(http.StatusOK,true,nil,nil)
	ctx.JSON(http.StatusOK,res)
}

func (or *OrderController) Update (ctx *gin.Context){
	var val model.Order
	err :=ctx.ShouldBindJSON(&val)
		if err!= nil {
			res := util.CreateResponse(http.StatusBadRequest, false, err, nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	errUp:=or.orderRepository.Update(&val)
		if errUp!=nil {
		res := util.CreateResponse(http.StatusInternalServerError, false, err, nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
		}
			res := util.CreateResponse(http.StatusOK, true, nil, nil)
	ctx.JSON(http.StatusOK, res)
}

