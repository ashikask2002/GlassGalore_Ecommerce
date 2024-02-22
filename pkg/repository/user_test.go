package repository

import (
	"GlassGalore/pkg/utils/models"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func Test_GetUserDetails(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		args    int
// 		stub    func(mockSQL sqlmock.Sqlmock)
// 		want    models.UserDetailsResponse
// 		wantErr error
// 	}{
// 		{
// 			name: "success",
// 			args: 1,
// 			stub: func(mockSQL sqlmock.Sqlmock) {
// 				expectQuery := `^select\s+id\s*,\s*name\s*,\s*email\s*,\s*phone\s+from\s+users.*$`

// 				mockSQL.ExpectQuery(expectQuery).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone"}).AddRow(1, "ashik", "ashik@gmail.com", "7510468623"))
// 			},
// 			want: models.UserDetailsResponse{
// 				Id:    1,
// 				Name:  "ashik",
// 				Email: "ashik@gmail.com",
// 				Phone: "7510468623",
// 			},
// 			wantErr: nil,
// 		},
// 		{
// 			name: "error",
// 			args: 1,
// 			stub: func(mockSQL sqlmock.Sqlmock) {
// 				expectQuery := `^select\s+id\s*,\s*name\s*,\s*email\s*,\s*phone\s+from\s+users.*$
// 				`
// 				mockSQL.ExpectQuery(expectQuery).WillReturnError(sqlmock.ErrCancelled)

// 			},
// 			want:    models.UserDetailsResponse{},
// 			wantErr: errors.New("could not get the user  details"),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mockDB, mockSQL, _ := sqlmock.New()
// 			defer mockDB.Close()
// 			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
// 				Conn: mockDB,
// 			}), &gorm.Config{})
// 			tt.stub(mockSQL)
// 			u := NewUserRepository(gormDB)
// 			result, err := u.GetUserDetails(tt.args)
// 			assert.Equal(t, tt.want, result)
// 			assert.Equal(t, tt.wantErr, err)

// 		})
// 	}

// }

func Test_user_UsersignUp(t *testing.T) {
	type args struct {
		input models.UserDetails
	}
	tests := []struct {
		name    string
		args    args
		stub    func(sqlmock.Sqlmock)
		want    models.UserDetailsResponse
		wantErr error
	}{
		{
			name: "success",
			args: args{
				input: models.UserDetails{Name: "Ashik", Email: "ashik@gmail.com", Phone: "7510468624", Password: "12345", ConfirmPassword: "12345"},
			},
			stub: func(mocSQL sqlmock.Sqlmock) {
				expectedQuery := `^INSERT INTO users (.+)$`
				mocSQL.ExpectQuery(expectedQuery).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone"}).AddRow(1, "Ashik", "ashik@gmail.com", "7510468624"))
			},
			want:    models.UserDetailsResponse{Id: 1, Name: "Ashik", Email: "ashik@gmail.com", Phone: "7510468624"},
			wantErr: nil,
		},
		{
			name: "failure",
			args: args{
				input: models.UserDetails{Name: "", Email: "", Phone: "", Password: "", ConfirmPassword: ""},
			},
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^INSERT INTO users (.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(errors.New("not possible"))
			},
			want:    models.UserDetailsResponse{},
			wantErr: errors.New("not possible"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()
			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})
			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)
			got, err := u.UserSignUp(tt.args.input)
			assert.Equal(t, tt.wantErr, err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.UserSignUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func test_CheckUserAvailability(t *testing.T) {
	tests := []struct {
		name string
		args string
		stub func(sqlmock.Sqlmock)
		want bool
	}{
		{name : "user available",
	     args:  "ashik@gmail.com",
		 stub: func(s sqlmock.Sqlmock) {
			expectedQuery:= `^select count\(\*\) from users(.+)$`
			mock
		 },
	}
	}

}
