package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurcholisnanda/wallet-record/models"
	"github.com/nurcholisnanda/wallet-record/services"
)

type Controller interface {
	CreateRecord(ctx *gin.Context)
	GetLatest(ctx *gin.Context)
	GetHistory(ctx *gin.Context)
}

type controller struct {
	service services.Service
}

func NewController(s services.Service) Controller {
	return &controller{
		service: s,
	}
}

// CreateRecord godoc
// @Summary      Insert record for wallet balance
// @Description  insert new transaction
// @Tags         wallets
// @Accept       json
// @Produce      json
// @Param        body  body      models.Record  true  "Models of record"
// @Success      201   {object}  models.BasicResponse
// @Failure      400   {object}  models.BasicResponse
// @Router       /records [post]
func (c *controller) CreateRecord(ctx *gin.Context) {
	var record models.Record
	if err := ctx.ShouldBindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BasicResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Success: false,
		})
		return
	}
	if err := record.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BasicResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Success: false,
		})
		return
	}
	if err := c.service.CreateRecord(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BasicResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Success: false,
		})
		return
	}
	ctx.JSON(http.StatusCreated, models.BasicResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Success: true,
	})
}

// GetHistory godoc
// @Summary      Get History from start to end date time
// @Description  get history balance
// @Tags         wallets
// @Accept       json
// @Produce      json
// @Param        body  body      models.History  true  "Models of request"
// @Success      200   {object}  []models.Record
// @Failure      400   {object}  models.BasicResponse
// @Router       /records/history [post]
func (c *controller) GetHistory(ctx *gin.Context) {
	var history models.History
	if err := ctx.ShouldBindJSON(&history); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BasicResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Success: false,
		})
		return
	}
	if err := history.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BasicResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Success: false,
		})
		return
	}
	if history.EndDatetime.Before(history.StartDatetime) {
		ctx.JSON(http.StatusBadRequest, models.BasicResponse{
			Status:  http.StatusBadRequest,
			Message: "endTime needs to be after startTime",
			Success: false,
		})
		return
	}
	records, err := c.service.GetHistory(history.StartDatetime, history.EndDatetime)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.BasicResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Success: false,
		})
		return
	}
	ctx.JSON(http.StatusOK, records)
}

// GetLatest godoc
// @Summary      Get Latest Balance
// @Description  get lattest balance
// @Tags         wallets
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Record
// @Failure      400  {object}  models.BasicResponse
// @Router       /records/latest [get]
func (c *controller) GetLatest(ctx *gin.Context) {
	record, err := c.service.GetLatest()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.BasicResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Success: false,
		})
		return
	}
	ctx.JSON(http.StatusOK, record)
}
