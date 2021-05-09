package controller

import (
	"net/http"
	"strconv"

	"calc/model"
	"calc/service"
	"github.com/gin-gonic/gin"
)

func CalcAdd(c *gin.Context) {
	var p model.CalcAddPaylod
	if err := c.ShouldBindUri(&p); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	s, err := service.NewCalcService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ret, err :=s.Add(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, strconv.Itoa(ret))
}

func CalcDiv(c *gin.Context) {
	var p model.CalcDivPayload
	if err := c.BindUri(&p); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if *p.Y == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "div by zero",
		})
		return
	}
	s, err := service.NewCalcService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ret, err :=s.Div(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, strconv.Itoa(ret))
}
