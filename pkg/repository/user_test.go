package repository

import (
	"GlassGalore/pkg/utils/models"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
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

func Test_CheckUserAvailability(t *testing.T) {
	tests := []struct {
		name string
		args string
		stub func(sqlmock.Sqlmock)
		want bool
	}{
		{
			name: "user available",
			args: "ashik@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select count\(\*\) from users(.+)$`
				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))

			},
			want: true,
		},
		{
			name: "user not available",
			args: "ashik@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expeectedQuery := `^select count\(\*\) from users(.+)$`
				mockSQL.ExpectQuery(expeectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
			},
			want: false,
		},
		{
			name: "error from database",
			args: "ashik@gmail.com",
			stub: func(s sqlmock.Sqlmock) {
				expectedquery := `^select count\(\*\) from users(.+)$`
				s.ExpectQuery(expectedquery).
					WillReturnError(errors.New("text string"))
			},
			want: false,
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
			result := u.CheckUserAvailability(tt.args)
			assert.Equal(t, tt.want, result)
		})
	}

}
func Test_userBlockStatus(t *testing.T) {
	test := []struct {
		name    string
		args    string
		stub    func(sqlmock.Sqlmock)
		want    bool
		wanterr error
	}{
		{
			name: "user is blocked",
			args: "ashik@gmail.com",
			stub: func(s sqlmock.Sqlmock) {
				expectedQuery := `select blocked from users where email = ?`
				s.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"is_blocked"}).AddRow(true))
			},
			want:    true,
			wanterr: nil,
		},
		{
			name: "user is not blocked",
			args: "ashik@gmail.com",
			stub: func(s sqlmock.Sqlmock) {
				expectedQuery := `select blocked from users where email = ?`
				s.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"is_blocked"}).AddRow(false))
			},
			want:    false,
			wanterr: nil,
		},

		{name: "error from database",
			args: "ashik@gmail.com",
			stub: func(s sqlmock.Sqlmock) {
				expectedQuery := `select blocked from users where email = ?`
				s.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("text string"))
			},
			want:    false,
			wanterr: errors.New("text string"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()
			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})
			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)
			result, err := u.UserBlockStatus(tt.args)
			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wanterr, err)
		})
	}
}

func Test_findUserByEmail(t *testing.T) {
	tests := []struct {
		name    string
		args    models.UserLogin
		stub    func(sqlmock.Sqlmock)
		want    models.UserSignInResponse
		wanterr error
	}{
		{
			name: "success",
			args: models.UserLogin{
				Email:    "ashik@gmail.com",
				Password: "12345",
			},
			stub: func(s sqlmock.Sqlmock) {
				expectQuery := `^SELECT \* FROM users(.+)$`
				s.ExpectQuery(expectQuery).WithArgs("ashik@gmail.com").
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "name", "email", "phone", "password"}).AddRow(1, 1, "ashik", "ashik@gmail.com", "7510468623", "12345"))
			},
			want: models.UserSignInResponse{
				Id:       1,
				UserID:   1,
				Name:     "ashik",
				Email:    "ashik@gmail.com",
				Phone:    "7510468623",
				Password: "12345",
			},
			wanterr: nil,
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

			result, err := u.FindUserByEmail(tt.args)
			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wanterr, err)
		})
	}
}

func Test_Addaddress(t *testing.T) {
	tests := []struct {
		name    string
		args    models.AddAddress
		stub    func(sqlmock.Sqlmock)
		wanterr error
	}{
		{
			name: "success",
			args: models.AddAddress{
				Name:      "Ashik",
				HouseName: "kizhakoodan",
				Street:    "kanakamala",
				City:      "chalakudy",
				State:     "kerala",
				Phone:     "7510468624",
				Pin:       "680680",
			},
			stub: func(s sqlmock.Sqlmock) {
				s.ExpectExec("INSERT INTO addresses").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wanterr: nil,
		},
		{
			name: "error",
			args: models.AddAddress{
				Name:      "Ashik",
				HouseName: "kizhakoodan",
				Street:    "kanakamala",
				City:      "chalakudy",
				State:     "kerala",
				Phone:     "7510468624",
				Pin:       "680680",
			},
			stub: func(s sqlmock.Sqlmock) {
				s.ExpectExec("INSERT INTO addresses").WillReturnError(errors.New("could not add address"))
			},
			wanterr: errors.New("could not add address"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocDB, mocksql, _ := sqlmock.New()
			defer mocDB.Close()
			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mocDB,
			}), &gorm.Config{})
			tt.stub(mocksql)
			u := NewUserRepository(gormDB)
			err := u.AddAddress(1, tt.args, true)
			assert.Equal(t, tt.wanterr, err)
		})
	}
}

func Test_CheckifFirstAddress(t *testing.T) {
	testin := []struct {
		name string
		args int
		stub func(sqlmock.Sqlmock)
		want bool
	}{
		{
			name: "first address",
			args: 1,
			stub: func(s sqlmock.Sqlmock) {
				expectedQuery := `^select count\(\*\) from addresses(.+)$`
				s.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{gomock.Any().String()}).AddRow(2))
			},
			want: true,
		},
		{
			name: "error occured",
			args: 2,
			stub: func(s sqlmock.Sqlmock) {
				expectedquey := `select count(*) from addresses where user_id = ?`
				s.ExpectQuery(expectedquey).WillReturnError(errors.New("error"))
			},
			want: false,
		},
	}
	for _, tt := range testin {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})
			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)
			result := u.CheckIfFirstAddress(tt.args)
			assert.Equal(t, tt.want, result)

		})
	}
}
