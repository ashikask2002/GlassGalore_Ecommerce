package usecase

import (
	"GlassGalore/pkg/domain"
	interfaces "GlassGalore/pkg/repository/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
	"fmt"

	helper_interface "GlassGalore/pkg/helper/interfaces"
	services "GlassGalore/pkg/usecase/interfaces"

	"github.com/jinzhu/copier"
	"github.com/jung-kurt/gofpdf"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
	helper          helper_interface.Helper
}

func NewAdminUseCase(repo interfaces.AdminRepository, h helper_interface.Helper) services.AdminUseCase {
	return &adminUseCase{
		adminRepository: repo,
		helper:          h,
	}
}

func (ad *adminUseCase) LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error) {

	// getting details of the admin based on the email provided
	adminCompareDetails, err := ad.adminRepository.LoginHandler(adminDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	// compare password from database and that provided from admins
	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return domain.TokenAdmin{}, err
	}
	var adminDetailsResponse models.AdminDetailsResponse

	// copy all details except password and sent it back to the front end
	err = copier.Copy(&adminDetailsResponse, &adminCompareDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}
	access, refresh, err := ad.helper.GenerateTokenAdmin(adminDetailsResponse)

	if err != nil {
		return domain.TokenAdmin{}, err
	}

	return domain.TokenAdmin{
		Admin:        adminDetailsResponse,
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil

}

func (ad *adminUseCase) GetUsers(page int) ([]models.UserDetailsAtAdmin, error) {

	userDetails, err := ad.adminRepository.GetUsers(page)
	if err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}
	return userDetails, nil
}

func (ad *adminUseCase) BlockUser(id string) error {
	user, err := ad.adminRepository.GetUserByID(id)
	if err != nil {
		return err
	}
	if user.Blocked {
		return errors.New("already blocked")
	} else {
		user.Blocked = true
	}
	err = ad.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}
	return nil
}
func (ad *adminUseCase) UnBlockUser(id string) error {
	user, err := ad.adminRepository.GetUserByID(id)
	if err != nil {
		return err

	}
	if user.Blocked {
		user.Blocked = false
	} else {
		return errors.New("already unblocked")
	}

	err = ad.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}
	return nil
}

func (i *adminUseCase) NewPaymentMethod(id string) error {
	if id == "" {
		return errors.New("not allowed empty name")
	}
	exists, err := i.adminRepository.CheckIfPaymentMethodAlreadyExists(id)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("payment method already exist")
	}

	err = i.adminRepository.NewPaymentMethod(id)
	if err != nil {
		return err
	}
	return nil
}

func (i *adminUseCase) ListPaymentMethods() ([]domain.PaymentMethod, error) {
	categories, err := i.adminRepository.ListPaymentMethods()
	if err != nil {
		return []domain.PaymentMethod{}, err
	}
	return categories, nil
}

func (i *adminUseCase) DeletePaymentMethod(id int) error {
	if id <= 0 {
		return errors.New("id must be positive")
	}
	err := i.adminRepository.DeletePaymentMethod(id)
	if err != nil {
		return err
	}
	return nil
}
func (i *adminUseCase) DashBoard() (models.CompleteAdminDashboard, error) {
	userDetails, err := i.adminRepository.DashBoardUserDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	productDetails, err := i.adminRepository.DashBoardProductDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	orderDetails, err := i.adminRepository.DashBoardOrder()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	totalRevenue, err := i.adminRepository.TotalRevenue()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	amountDetails, err := i.adminRepository.AmountDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	return models.CompleteAdminDashboard{
		DashboardUser:    userDetails,
		DashboardProduct: productDetails,
		DashboardOrder:   orderDetails,
		DashboardRevenue: totalRevenue,
		DashboardAmount:  amountDetails,
	}, nil
}

func (i *adminUseCase) FilteredSalesReport(timePeriod string) (*gofpdf.Fpdf, error) {
	if timePeriod == "" {
		return nil, errors.New("must have to provide something")
	}
	if timePeriod != "week" && timePeriod != "month" && timePeriod != "year" {
		return nil, errors.New("provided time period is not correct (week, month, year) are only available")
	}

	startTime, endTime := i.helper.GetTimeFromPeriod(timePeriod)
	salesReport, err := i.adminRepository.FilteredSalesReport(startTime, endTime)
	if err != nil {
		return nil, err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "Sales Report")

	// Output the sales report data in the PDF
	pdf.Ln(10)
	pdf.Ln(10)
	y := pdf.GetY()
	// Output the sales report data in the PDF
	pdf.Ln(10)
	pdf.Ln(10)
	pdf.SetX(10)
	pdf.SetY(y + 10)
	pdf.Cell(0, 10, "Total Sales: "+fmt.Sprintf("%.2f", salesReport.TotalSales))

	pdf.Ln(10)
	pdf.SetX(10)
	pdf.SetY(y + 20)
	pdf.Cell(0, 10, "Total Orders: "+fmt.Sprintf("%d", salesReport.TotalOrders))

	pdf.Ln(10)
	pdf.SetX(10)
	pdf.SetY(y + 30)
	pdf.Cell(0, 10, "Completed Orders: "+fmt.Sprintf("%d", salesReport.CompletedOrders))

	pdf.Ln(10)
	pdf.SetX(10)
	pdf.SetY(y + 40)
	pdf.Cell(0, 10, "Pending Orders: "+fmt.Sprintf("%d", salesReport.PendingOrders))

	pdf.Ln(10)
	pdf.SetX(10)
	pdf.SetY(y + 50)
	pdf.Cell(0, 10, "Returned Orders: "+fmt.Sprintf("%d", salesReport.ReturnedOrders))

	pdf.Ln(10)
	pdf.SetX(10)
	pdf.SetY(y + 60)
	pdf.Cell(0, 10, "Cancelled Orders: "+fmt.Sprintf("%d", salesReport.CancelledOrders))

	pdf.Ln(10)
	pdf.SetX(10)
	pdf.SetY(y + 70)
	pdf.Cell(0, 10, "Trending Product: "+salesReport.TrendingProduct)

	// Save the PDF file
	// err = pdf.OutputFileAndClose("sales_report.pdf")
	// if err != nil {
	// 	return pdf, err
	// }

	return pdf, nil
}
