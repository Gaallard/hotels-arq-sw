package hotels

import (
	"context"
	"fmt"
	hotelsDomain "hotels-api/domain/hotels"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetHotelByID(ctx context.Context, id int64) (hotelsDomain.Hotel, error)
	InsertHotel(ctx context.Context, hotel hotelsDomain.Hotel) error
	DeleteHotel(ctx context.Context, id int64) error
}

type Controller struct {
	service Service
}

func NewController(service Service) Controller {
	return Controller{
		service: service,
	}
}

func (controller Controller) GetHotelByID(ctx *gin.Context) {
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
	hotel, err := controller.service.GetHotelByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error getting hotel: %s", err.Error()),
		})
		return
	}

	// Send response
	ctx.JSON(http.StatusOK, hotel)
}

func (controller Controller) InsertHotel(ctx *gin.Context) {
	var hotel hotelsDomain.Hotel

	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Error passing hotel: %s", err),
		})
		return
	}

	if err := controller.service.InsertHotel(ctx.Request.Context(), hotel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error creating hotel: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "hotel created successfully",
		"hotel":   hotel,
	})

}

func (controller Controller) DeleteHotel(ctx *gin.Context) {

	idParam := strings.TrimSpace(ctx.Param("id"))
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid id: %s", idParam),
		})
		return
	}

	// Get hotel by ID using the service
	err = controller.service.DeleteHotel(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error deleting hotel: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "hotel deleted successfully",
	})

}
