package helper

import (
	cfg "GlassGalore/pkg/config"
	"GlassGalore/pkg/helper/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	"golang.org/x/crypto/bcrypt"

	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

type helper struct {
	cfg cfg.Config
}

func NewHelper(config cfg.Config) interfaces.Helper {
	return &helper{
		cfg: config,
	}
}

var client *twilio.RestClient

type AuthcustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func (h *helper) PasswordHashing(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}

	hash := string(hashedPassword)
	return hash, nil
}

func (h *helper) GenerateTokenClients(user models.UserDetailsResponse) (string, error) {
	claims := &AuthcustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "client",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("comebuyglass"))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (h *helper) CompareHashAndPassword(a string, b string) error {
	err := bcrypt.CompareHashAndPassword([]byte(a), []byte(b))
	if err != nil {
		return err
	}
	return nil
}

func (helper *helper) GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, string, error) {
	accessTokenClaims := &AuthcustomClaims{
		Id: admin.ID,
		//Email: admin.Email,
		Role: "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 50).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshTokenClaims := &AuthcustomClaims{
		Id: admin.ID,
		//Email: admin.Email,
		Role: "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accesToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accesToken.SignedString([]byte("1234"))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte("refreshsecret"))
	if err != nil {
		return "", "", err
	}

	fmt.Println("accegshshjskl;", accessTokenString)
	return accessTokenString, refreshTokenString, nil
}

func (h *helper) TwilioSetup(username string, password string) {
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: username,
		Password: password,
	})

}

func (h *helper) TwilioSendOTP(phone string, serviceID string) (string, error) {
	to := "+91" + phone
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(serviceID, params)
	if err != nil {
		return " ", err
	}
	return *resp.Sid, nil
}

func (h *helper) TwilioVerifyOTP(serviceID string, code string, phone string) error {

	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo("+91" + phone)
	params.SetCode(code)
	resp, err := client.VerifyV2.CreateVerificationCheck(serviceID, params)

	if err != nil {
		return err
	}

	if *resp.Status == "approved" {
		return nil
	}

	return errors.New("failed to validate otp")
}

func (h *helper) PhoneValidation(phone string) bool {
	phoneNumber := phone
	pattern := `^\d{10}$`
	regex := regexp.MustCompile(pattern)
	value := regex.MatchString(phoneNumber)
	return value
}

func (h *helper) IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	value := regex.MatchString(email)
	return value
}

func (h *helper) IsValidPIN(pincode string) bool {
	// Regular expression for a simple PIN code validation (6 digits)
	pattern := `^\d{6}$`
	regex := regexp.MustCompile(pattern)
	value := regex.MatchString(pincode)

	return value
}

func (h *helper) GetTimeFromPeriod(timePeriod string) (time.Time, time.Time) {

	endDate := time.Now()

	if timePeriod == "week" {
		startDate := endDate.AddDate(0, 0, -6)
		return startDate, endDate
	}

	if timePeriod == "month" {
		startDate := endDate.AddDate(0, -1, 0)
		return startDate, endDate
	}

	if timePeriod == "year" {
		startDate := endDate.AddDate(-1, 0, 0)
		return startDate, endDate
	}

	return endDate.AddDate(0, 0, -6), endDate

}

func (h *helper) ValidateAlphabets(data string) (bool, error) {
	for _, char := range data {
		if !unicode.IsLetter(char) {
			return false, errors.New("data contains non-alphabetic characters")
		}
	}
	return true, nil
}

func ConvertToExel(sales []models.OrderDetailsAdmin) (*excelize.File, error) {

	filename := "salesReport/sales_report.xlsx"
	file := excelize.NewFile()

	file.SetCellValue("Sheet1", "A1", "Item")
	file.SetCellValue("Sheet1", "B1", "Total Amount Sold")

	for i, sale := range sales {
		col1 := fmt.Sprintf("A%d", i+1)
		col2 := fmt.Sprintf("B%d", i+1)

		file.SetCellValue("Sheet1", col1, sale.ProductName)
		file.SetCellValue("Sheet1", col2, sale.TotalAmount)

	}

	if err := file.SaveAs(filename); err != nil {
		return nil, err
	}

	return file, nil
}

func GetImageMimeType(filename string) string {
	extension := strings.ToLower(strings.Split(filename, ".")[len(strings.Split(filename, "."))-1])

	imageMimeTypes := map[string]string{
		"jpg":  "image/jpeg",
		"jpeg": "image/jpeg",
		"png":  "image/png",
		"gif":  "image/gif",
		"bmp":  "image/bmp",
		"webp": "image/webp",
	}

	if mimeType, ok := imageMimeTypes[extension]; ok {
		return mimeType
	}

	return "application/octet-stream"
}
func AddImageToS3(file *multipart.FileHeader) (string, error) {
	f, openErr := file.Open()
	if openErr != nil {

		fmt.Println("opening error:", openErr)
		return "", openErr
	}
	defer f.Close()
	if err := godotenv.Load(); err != nil {
		fmt.Println("error 1", err)
		return "", err
	}
	mimeType := GetImageMimeType(file.Filename)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})
	if err != nil {
		fmt.Println("error in session config", err)
		return "", err
	}
	// Create an S3 uploader with the session and default options
	uploader := s3manager.NewUploader(sess)
	BucketName := "glassgalore"
	//upload data(video or image)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(BucketName),
		Key:         aws.String(file.Filename),
		Body:        f,
		ContentType: aws.String(mimeType),
	})
	if err != nil {
		fmt.Println("error 2", err)
		return "", err
	}
	url := fmt.Sprintf("https://glassgalore.s3.ap-south-1.amazonaws.com/%s", file.Filename)
	return url, nil
}
