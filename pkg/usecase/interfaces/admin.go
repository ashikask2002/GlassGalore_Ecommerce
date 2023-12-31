package interfaces

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/utils/models"

	"github.com/jung-kurt/gofpdf"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
	GetUsers(page int) ([]models.UserDetailsAtAdmin, error)
	BlockUser(id string) error
	UnBlockUser(id string) error
	NewPaymentMethod(id string) error
	ListPaymentMethods() ([]domain.PaymentMethod, error)
	DeletePaymentMethod(id int) error
	DashBoard() (models.CompleteAdminDashboard, error)
	FilteredSalesReport(timePeriod string) (*gofpdf.Fpdf, error)
}
