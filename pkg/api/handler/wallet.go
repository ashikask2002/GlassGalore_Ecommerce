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
