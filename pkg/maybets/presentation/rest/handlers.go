package rest

import (
	"net/http"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/usecases"
	"github.com/gin-gonic/gin"
)

// HandlersInterfacesImpl represents the usecase implementation object
type HandlersInterfacesImpl struct {
	usecase *usecases.UsecaseMayBets
}

// NewHandlersInterfaces initializes a new rest handlers usecase
func NewHandlersInterfaces(i *usecases.UsecaseMayBets) *HandlersInterfacesImpl {
	return &HandlersInterfacesImpl{i}
}

// GetUserTotalBets endpoint to get all a user's total bets
func (h HandlersInterfacesImpl) GetUserTotalBets(c *gin.Context) {
	userID := c.Query("user_id")

	user, err := h.usecase.GetUserTotalBets(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": user,
	})
}

// GetUserTotalWinnings endpoint to get a user's total winnings.
func (h HandlersInterfacesImpl) GetUserTotalWinnings(c *gin.Context) {
	userID := c.Query("user_id")

	user, err := h.usecase.GetUserTotalWinnings(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": user,
	})
}

// GetTopFiveUsers endpoint to get top 5 users with highest betting volume
func (h HandlersInterfacesImpl) GetTopFiveUsers(c *gin.Context) {
	users, err := h.usecase.GetTopFiveUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": users,
	})
}

// GetAllAnomalousUsers endpoint to get all users with significantly higher
// betting activity than the average.
func (h HandlersInterfacesImpl) GetAllAnomalousUsers(c *gin.Context) {
	users, err := h.usecase.GetAllAnomalousUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": users,
	})
}
