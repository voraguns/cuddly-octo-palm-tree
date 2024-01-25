package backlogimpl

import (
	bi "cuddly-octo-palm-tree/service/backlogimpl"

	"github.com/gin-gonic/gin"
)

type Router struct {
	biService bi.BiService
}

func NewRouter(svc bi.BiService) *Router {
	return &Router{
		biService: svc,
	}
}

func (r *Router) GetUsers(c *gin.Context) {

}

func (r *Router) GetOrders(c *gin.Context) {

}

func (r *Router) GetOrderItems(c *gin.Context) {

}

func (r *Router) GetProducts(c *gin.Context) {

}
