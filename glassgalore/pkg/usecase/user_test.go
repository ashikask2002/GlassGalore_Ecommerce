package usecase

import (
	"errors"
	"testing"

	"GlassGalore/pkg/domain"
	mockhelper "GlassGalore/pkg/helper/mock"
	mockrepo "GlassGalore/pkg/repository/mock"
	"GlassGalore/pkg/utils/models"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func Test_GetAddresses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userrepo := mockrepo.NewMockUserRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)

	userUsecase := NewUserUseCase(userrepo, helper)

	testData := map[string]struct {
		input   int
		stub    func(*mockrepo.MockUserRepository, *mockhelper.MockHelper, int)
		want    []domain.Address
		wantErr error
	}{
		"success": {
			input: 1,
			stub: func(userrepo *mockrepo.MockUserRepository, helper *mockhelper.MockHelper, data int) {
				userrepo.EXPECT().GetAddresses(data).Times(1).Return([]domain.Address{}, nil)
			},
			want:    []domain.Address{},
			wantErr: nil,
		},
		"failed": {
			input: 1,
			stub: func(userrepo *mockrepo.MockUserRepository, helper *mockhelper.MockHelper, data int) {
				userrepo.EXPECT().GetAddresses(data).Times(1).Return([]domain.Address{}, errors.New("error"))
			},
			want:    []domain.Address{},
			wantErr: errors.New("error in getting addresses"), // Corrected error string
		},
	}
	for _, test := range testData {
		test.stub(userrepo, helper, test.input)
		result, err := userUsecase.GetAddresses(test.input)
		assert.Equal(t, test.want, result)
		assert.Equal(t, test.wantErr, err)
	}
}

func Test_AddAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockrepo.NewMockUserRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)

	userUsecase := NewUserUseCase(userRepo, helper)
	testData := map[string]struct {
		input   models.AddAddress
		stub    func(*mockrepo.MockUserRepository, *mockhelper.MockHelper, models.AddAddress)
		WantErr error
	}{
		"success": {
			input: models.AddAddress{
				Name:      "ashik",
				HouseName: "kizhakoodan",
				Street:    "kanakamala",
				City:      "chalakudy",
				State:     "kerala",
				Phone:     "7510468624",
				Pin:       "680680",
			},
			stub: func(userrepo *mockrepo.MockUserRepository, helper *mockhelper.MockHelper, data models.AddAddress) {
				userRepo.EXPECT().CheckIfFirstAddress(1).Times(1).Return(false) // Setting expectation for CheckIfFirstAddress
				userRepo.EXPECT().AddAddress(1, data, true).Times(1).Return(nil)
			},
			WantErr: nil,
		},
		"failure": {
			input: models.AddAddress{
				Name:      "ashik",
				HouseName: "kizhakoodan",
				Street:    "kanakamala",
				City:      "chalakudy",
				State:     "kerala",
				Phone:     "7510468624",
				Pin:       "680680",
			},
			stub: func(userrepo *mockrepo.MockUserRepository, helper *mockhelper.MockHelper, data models.AddAddress) {
				userRepo.EXPECT().CheckIfFirstAddress(1).Times(1).Return(false) // Setting expectation for CheckIfFirstAddress
				userRepo.EXPECT().AddAddress(1, data, true).Times(1).Return(errors.New("error in adding"))
			},
			WantErr: errors.New("error in adding address"),
		},
	}
	for _, test := range testData {
		test.stub(userRepo, helper, test.input)
		err := userUsecase.AddAddress(1, test.input)
		assert.Equal(t, test.WantErr, err)
	}
}
