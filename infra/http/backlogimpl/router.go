package backlogimpl

import (
	bi "cuddly-octo-palm-tree/service/backlogimpl"
	"net/http"
	"strconv"

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
	users, err := r.biService.GetUsers(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, users)
}

func (r *Router) GetOrder(c *gin.Context) {
	order, err := r.biService.GetOrders(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, order)
}

func (r *Router) GetOrderItems(c *gin.Context) {
	orderItem, err := r.biService.GetOrderItems(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errror": err,
		})
	}
	c.JSON(http.StatusOK, orderItem)
}

func (r *Router) GetProducts(c *gin.Context) {
	products, err := r.biService.GetProducts(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, products)
}

func (r *Router) GetUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user id",
		})
	}

	user, err := r.biService.GetUsersById(c, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot find user",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (r *Router) DeleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user id",
		})
	}

	r.biService.DeleteUser(c, userId)
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete User successfully",
	})
}
