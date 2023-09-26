package api

import (
	"api-example/model"
	"api-example/pkg/dao"
	"api-example/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	router       *gin.Engine
	srv          service.CustomerManager
	CustomerList model.CustomerList
}

func (a *ApiServer) Run() {
	a.init()
	a.start()
}

func (a *ApiServer) init() error {
	a.CustomerList = model.CustomerList{}
	cdao, err := dao.NewCustomerDao()
	if err != nil {
		return err
	}
	srv, err := service.NewCustomerService(cdao)
	if err != nil {
		return err
	}
	a.srv = srv

	a.router = gin.New()
	a.router.GET("/healthcheck", Healthcheck)

	a.router.GET("/customers", a.ListCustomers)
	a.router.POST("/customers", a.AddCustomer)
	a.router.GET("/customers/:id", a.GetCustomersById)

	return nil
}

func (a *ApiServer) start() error {
	err := a.router.Run(":8080")
	return err
}

func Healthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "I'm alive")
}

func (a *ApiServer) ListCustomers(ctx *gin.Context) {
	list, err := a.srv.GetCustomers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "customer listing error")
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func (a *ApiServer) AddCustomer(ctx *gin.Context) {
	customer := &model.Customer{}
	err := ctx.BindJSON(customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "customer creation error")
		return
	}
	err = a.srv.AddCustomer(customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "customer adding error")
		return
	}
	ctx.JSON(http.StatusOK, customer)
}

func (a *ApiServer) GetCustomersById(ctx *gin.Context) {
	requestedCustomerId := ctx.Param("id")
	if !isIdInt(requestedCustomerId) {
		ctx.JSON(http.StatusBadRequest, "customer id must be integer")
		return
	}

	if !a.isCustomerListPopulated() {
		ctx.JSON(http.StatusBadRequest, "customer list is empty")
		return
	}

	customer := a.matchCustomerId(requestedCustomerId)
	if customer == nil {
		ctx.JSON(http.StatusBadRequest, "customer not found")
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (a *ApiServer) isCustomerListPopulated() bool {
	list, err := a.srv.GetCustomers()
	if err != nil {
		return false
	}
	if len(list) == 0 {
		return false
	}
	return true
}

func (a *ApiServer) matchCustomerId(id string) *model.Customer {
	customerId, _ := strconv.Atoi(id)
	customers, _ := a.srv.GetCustomers()

	for _, customer := range customers {
		if customer.ID == customerId {
			return &customer
		}
	}
	return nil
}

func isIdInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
