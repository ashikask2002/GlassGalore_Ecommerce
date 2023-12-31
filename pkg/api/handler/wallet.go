package handler

import (
	"GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	WalletUseCase interfaces.WalletUseCase
}

func NewWalletHandler(usecase interfaces.WalletUseCase) *WalletHandler {
	return &WalletHandler{
		WalletUseCase: usecase,
	}

}

// @Summary View user's wallet details
// @Description Retrieve details of the wallet for the authenticated user
// @Accept json
// @Produce json
// @Tags Wallet
// @Security BearerTokenAuth
// @Param id path int true "User ID"
// @Success 200 {object} response.Response "Wallet details"
// @Failure 400 {object} response.Response "Error in retrieving wallet details"
// @Router /users/wallet [get]
func (i *WalletHandler) ViewWallet(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	walletDetails, err := i.WalletUseCase.GetWallet(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "couldnt get the wallet", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	succesRes := response.ClientResponse(http.StatusBadRequest, "successfully got the wallet", walletDetails, nil)
	c.JSON(http.StatusOK, succesRes)

}
