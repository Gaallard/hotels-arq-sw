package reservas

import (
	"context"
	"fmt"
	"net/http"
	Domain "reserva-api/domain"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetReservaById(ctx context.Context, id int64) (Domain.Reserva, error)
	InsertReserva(ctx context.Context, reserva Domain.Reserva) (Domain.Reserva, error)
	UpdateReserva(ctx context.Context, reserva Domain.Reserva) (Domain.Reserva, error)
	DeleteReserva(ctx context.Context, reserva Domain.Reserva) error
}

type Controller struct {
	service Service
}

func NewController(service Service) Controller {
	return Controller{
		service: service,
	}
}

func (controller Controller) GetReservaById(ctx *gin.Context) {
	// Validate ID param
	idParam := strings.TrimSpace(ctx.Param("id"))
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid id: %s", idParam),
		})
		return
	}

	// Get hotel by ID using the service
	reserva, err := controller.service.GetReservaById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error getting hotel: %s", err.Error()),
		})
		return
	}

	// Send response
	ctx.JSON(http.StatusOK, reserva)
}

func (controller Controller) InsertReserva(ctx *gin.Context) {
	var reservaDomain Domain.Reserva
	err := ctx.BindJSON(&reservaDomain)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid data"),
		})
		return
	}
	reservaDomain, er := controller.service.InsertReserva(ctx, reservaDomain)
	if er != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error getting reserva: %s", er.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, reservaDomain)
}

func (controller Controller) UpdateReserva(ctx *gin.Context) {
	var reservaDomain Domain.Reserva
	err := ctx.BindJSON(&reservaDomain)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid data"),
		})
		return
	}
	reservaDomain, er := controller.service.UpdateReserva(ctx, reservaDomain)
	if er != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error getting reserva: %s", er.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, reservaDomain)
}

func (controller Controller) DeleteReserva(ctx *gin.Context) {
	var reservaDomain Domain.Reserva
	err := ctx.BindJSON(&reservaDomain)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid data"),
		})
		return
	}
	er := controller.service.DeleteReserva(ctx, reservaDomain)
	if er != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error getting reserva: %s", er.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
