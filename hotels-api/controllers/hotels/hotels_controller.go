package hotels

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	hotelsDomain "hotels-api/domain/hotels"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetHotelByID(ctx context.Context, id primitive.ObjectID) (hotelsDomain.Hotel, error)
	InsertHotel(ctx context.Context, hotel hotelsDomain.Hotel) (error, string)
	UpdateHotel(ctx context.Context, id primitive.ObjectID, hotel hotelsDomain.Hotel) (hotelsDomain.Hotel, error)
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
	idParam := strings.TrimSpace(ctx.Param("_id"))
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid id: %s", idParam),
		})
		return
	}

	hotel, err := controller.service.GetHotelByID(ctx.Request.Context(), objectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error getting hotel: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, hotel)
}

func (controller Controller) InsertHotel(ctx *gin.Context) {
	var hotel hotelsDomain.Hotel

	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("error passing hotel: %s", err),
		})
		return
	}

	err, hotelIdMongo := controller.service.InsertHotel(ctx.Request.Context(), hotel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error creating hotel: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "hotel created successfully",
		"hotelID": hotelIdMongo,
	})
}

func (controller Controller) UpdateHotel(ctx *gin.Context) {
	var hotel hotelsDomain.Hotel

	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("error passing hotel: %s", err),
		})
		return
	}

	updatedHotel, err := controller.service.UpdateHotel(ctx.Request.Context(), hotel.IdMongo, hotel)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error updating hotel: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "hotel updated successfully",
		"hotel":   updatedHotel,
	})
}
