package hotels

import (
	"context"
	"fmt"
	hotelsDomain "hotels-api/domain/hotels"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetHotelByID(ctx context.Context, id string) (hotelsDomain.Hotel, error)
	InsertHotel(ctx context.Context, hotel hotelsDomain.Hotel) (string, error)
	UpdateHotel(ctx context.Context, id string, hotel hotelsDomain.Hotel) error
	GetHotelsAvailability(ctx context.Context) (map[string]int64, error)
	GetAllHotels2(ctx context.Context) ([]hotelsDomain.Hotel, error)
}

type Controller struct {
	service Service
}

func NewController(service Service) Controller {
	return Controller{
		service: service,
	}
}

func (controller Controller) GetAllHotels2(ctx *gin.Context) {
	result, err := controller.service.GetAllHotels2(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error al obtener habitaciones disponibles: %v", err),
		})
		return
	}

	// Responder con el resultado en formato JSON
	ctx.JSON(http.StatusOK, result)
}

func (controller Controller) GetHotelByID(ctx *gin.Context) {
	objectID := strings.TrimSpace(ctx.Param("id"))
	println("recibimos3: ", objectID)

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

	result, err := controller.service.InsertHotel(ctx.Request.Context(), hotel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error creating hotel: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":  "hotel created successfully",
		"id mongo": result,
	})
}

func (controller Controller) UpdateHotel(ctx *gin.Context) {

	objectID := strings.TrimSpace(ctx.Param("id"))

	var hotelDomain hotelsDomain.Hotel
	if err := ctx.ShouldBindJSON(&hotelDomain); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("error passing hotel: %s", err),
		})
		return
	}

	err := controller.service.UpdateHotel(ctx.Request.Context(), objectID, hotelDomain)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error updating hotel: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "hotel updated successfully",
		"id":      objectID,
	})
}

func (controller Controller) GetAvailableRooms(ctx *gin.Context) {
	// Llama al servicio para obtener hoteles con habitaciones disponibles
	result, err := controller.service.GetHotelsAvailability(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error al obtener habitaciones disponibles: %v", err),
		})
		return
	}

	// Responder con el resultado en formato JSON
	ctx.JSON(http.StatusOK, gin.H{
		"available_rooms": result,
	})
}
