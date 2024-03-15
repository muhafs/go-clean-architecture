package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhafs/go-clean-architecture/bootstrap"
	"github.com/muhafs/go-clean-architecture/domain/entity"
	"github.com/muhafs/go-clean-architecture/domain/model/request"
	"github.com/muhafs/go-clean-architecture/domain/model/response"
	"github.com/muhafs/go-clean-architecture/domain/usecase"
)

type CategoryController struct {
	CategoryUsecase usecase.CategoryUsecase
	Env             *bootstrap.Env
}

func (cc *CategoryController) Create(c *gin.Context) {
	var request request.CategoryRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Success: false, Message: err.Error()})
		return
	}

	category := entity.Category{
		Name: request.Name,
	}
	if err := cc.CategoryUsecase.Create(&category); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Category created successfully",
		Data:    category,
	})
}

func (cc *CategoryController) FindAll(c *gin.Context) {
	categories, err := cc.CategoryUsecase.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Category list fetched successfully",
		Data:    categories,
	})
}

func (cc *CategoryController) FindOne(c *gin.Context) {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	category, err := cc.CategoryUsecase.FindOne(intID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Category fetched successfully",
		Data:    category,
	})
}
