package handlers_test

import (
	"fmt"
	"gotest/handlers"
	"gotest/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculationDiscount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		amount := 100
		expected := 80
		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expected, nil)

		promoHandler := handlers.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		res, _ := app.Test(req)

		// act
		defer res.Body.Close()

		//Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}
	})

	t.Run("convert fail", func(t *testing.T) {
		// Arrange
		amount := "pasit"
		expected := 80
		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expected, nil)

		promoHandler := handlers.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		res, _ := app.Test(req)

		// act
		defer res.Body.Close()

		//Assert
		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})

	t.Run("service error", func(t *testing.T) {
		// Arrange
		amount := -200
		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(0, services.ErrZeroAmount)

		promoHandler := handlers.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		res, _ := app.Test(req)

		// act
		defer res.Body.Close()

		//Assert
		assert.Equal(t, fiber.StatusNotFound, res.StatusCode)
	})
}
