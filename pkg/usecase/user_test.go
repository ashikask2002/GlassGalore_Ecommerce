package usecase

import (
	"errors"
	"testing"

	"GlassGalore/pkg/domain"
	mockhelper "GlassGalore/pkg/helper/mock"
	mockrepo "GlassGalore/pkg/repository/mock"

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

