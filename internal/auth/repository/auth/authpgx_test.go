package auth_test

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/auth/repository/auth"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

func TestPostgres_CheckExistUserByEmail_True(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputEmail := "CorrectEmail@mail.ru"

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowMain := sqlmock.NewRows([]string{"exist"})

	rowMain.AddRow(true)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CheckExist)).
		WithArgs(inputEmail). // Values in query
		WillReturnRows(rowMain)

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.CheckExistUserByEmail(ctx, inputEmail)
	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, true, actual)
}

func TestPostgres_CheckExistUserByEmail_False(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputEmail := "CorrectEmail@mail.ru"

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowMain := sqlmock.NewRows([]string{"exist"})

	rowMain.AddRow(false)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CheckExist)).
		WithArgs(inputEmail). // Values in query
		WillReturnRows(rowMain)

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.CheckExistUserByEmail(ctx, inputEmail)
	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, false, actual)
}

func TestPostgres_CheckExistUserByEmail_RowError(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputEmail := "CorrectEmail@mail.ru"

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowMain := sqlmock.NewRows([]string{"exist"})

	rowMain.AddRow(false)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CheckExist)).
		WithArgs(inputEmail). // Values in query
		WillReturnError(sql.ErrNoRows)

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.CheckExistUserByEmail(ctx, inputEmail)
	require.NotNil(t, err, fmt.Errorf("expected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, false, actual)
}

func TestPostgres_CheckExistUserByEmail_ScanError(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputEmail := "CorrectEmail@mail.ru"

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowMain := sqlmock.NewRows([]string{"exist"})
	rowMain.AddRow("asfalse")

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CheckExist)).
		WithArgs(inputEmail). // Values in query
		WillReturnRows(rowMain)

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.CheckExistUserByEmail(ctx, inputEmail)
	require.NotNil(t, err, fmt.Errorf("expected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, false, actual)
}

func TestPostgres_CreateUser_OK(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputUser := models.User{
		Email:    "CorrectEmail@mail.ru",
		Password: "CorrectPassword123",
		Nickname: "CorrectNickname",
		Avatar:   "avatar",
	}

	outputUser := models.User{
		ID: 1,
	}

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowCheckExist := sqlmock.NewRows([]string{"exist"})
	rowCheckExist.AddRow(false)

	rowCreateUser := sqlmock.NewRows([]string{"user_id"})
	rowCreateUser.AddRow(1)

	rowAddCollections := sqlmock.NewRows([]string{"collection_id"})
	rowAddCollections.AddRow(1)
	rowAddCollections.AddRow(2)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CheckExist)).
		WithArgs(inputUser.Email). // Values in query
		WillReturnRows(rowCheckExist)

	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CreateUser)).
		WithArgs(
			inputUser.Email,
			inputUser.Nickname,
			[]byte(inputUser.Password),
			inputUser.Avatar).
		WillReturnRows(rowCreateUser)

	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CreateDefCollections)).
		WillReturnRows(rowAddCollections)

	mock.ExpectExec(regexp.QuoteMeta(auth.LinkUserDefCollections)).
		WithArgs(1, 1, 2).
		WillReturnResult(sqlmock.NewResult(2, 2))
	mock.ExpectCommit()

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.CreateUser(ctx, &inputUser)
	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, outputUser, actual)
}

func TestPostgres_CreateUser_ExistErr(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputUser := models.User{
		Email:    "CorrectEmail@mail.ru",
		Password: "CorrectPassword123",
		Nickname: "CorrectNickname",
		Avatar:   "avatar",
	}

	outputUser := models.User{
		ID: 0,
	}

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowCheckExist := sqlmock.NewRows([]string{"exist"})
	rowCheckExist.AddRow("oops")

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CheckExist)).
		WithArgs(inputUser.Email). // Values in query
		WillReturnRows(rowCheckExist)

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.CreateUser(ctx, &inputUser)
	require.NotNil(t, err, fmt.Errorf("expected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, outputUser, actual)
}

func TestPostgres_CreateUser_Exist(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputUser := models.User{
		Email:    "CorrectEmail@mail.ru",
		Password: "CorrectPassword123",
		Nickname: "CorrectNickname",
		Avatar:   "avatar",
	}

	outputUser := models.User{
		ID: 0,
	}

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowCheckExist := sqlmock.NewRows([]string{"exist"})
	rowCheckExist.AddRow(true)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CheckExist)).
		WithArgs(inputUser.Email). // Values in query
		WillReturnRows(rowCheckExist)

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.CreateUser(ctx, &inputUser)
	require.NotNil(t, err, fmt.Errorf("expected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, outputUser, actual)
}

func TestPostgres_CreateUser_RowUserErr(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputUser := models.User{
		Email:    "CorrectEmail@mail.ru",
		Password: "CorrectPassword123",
		Nickname: "CorrectNickname",
		Avatar:   "avatar",
	}

	outputUser := models.User{
		ID: 0,
	}

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowCheckExist := sqlmock.NewRows([]string{"exist"})
	rowCheckExist.AddRow(false)

	rowCreateUser := sqlmock.NewRows([]string{"user_id"})
	rowCreateUser.AddRow(1)

	rowAddCollections := sqlmock.NewRows([]string{"collection_id"})
	rowAddCollections.AddRow(1)
	rowAddCollections.AddRow(2)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CheckExist)).
		WithArgs(inputUser.Email). // Values in query
		WillReturnRows(rowCheckExist)

	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CreateUser)).
		WithArgs(
			inputUser.Email,
			inputUser.Nickname,
			[]byte(inputUser.Password),
			inputUser.Avatar).
		WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.CreateUser(ctx, &inputUser)
	require.NotNil(t, err, fmt.Errorf("expected err: %s", sql.ErrNoRows))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, outputUser, actual)
}

func TestPostgres_CreateUser_RowUserScanErr(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputUser := models.User{
		Email:    "CorrectEmail@mail.ru",
		Password: "CorrectPassword123",
		Nickname: "CorrectNickname",
		Avatar:   "avatar",
	}

	outputUser := models.User{
		ID: 0,
	}

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowCheckExist := sqlmock.NewRows([]string{"exist"})
	rowCheckExist.AddRow(false)

	rowCreateUser := sqlmock.NewRows([]string{"user_id"})
	rowCreateUser.AddRow("oops")

	rowAddCollections := sqlmock.NewRows([]string{"collection_id"})
	rowAddCollections.AddRow(1)
	rowAddCollections.AddRow(2)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CheckExist)).
		WithArgs(inputUser.Email). // Values in query
		WillReturnRows(rowCheckExist)

	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CreateUser)).
		WithArgs(
			inputUser.Email,
			inputUser.Nickname,
			[]byte(inputUser.Password),
			inputUser.Avatar).
		WillReturnRows(rowCreateUser)
	mock.ExpectRollback()

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.CreateUser(ctx, &inputUser)
	require.NotNil(t, err, fmt.Errorf("expected err: %s", sql.ErrNoRows))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, outputUser, actual)
}

func TestPostgres_CreateUser_CollectionsErr(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputUser := models.User{
		Email:    "CorrectEmail@mail.ru",
		Password: "CorrectPassword123",
		Nickname: "CorrectNickname",
		Avatar:   "avatar",
	}

	outputUser := models.User{
		ID: 0,
	}

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowCheckExist := sqlmock.NewRows([]string{"exist"})
	rowCheckExist.AddRow(false)

	rowCreateUser := sqlmock.NewRows([]string{"user_id"})
	rowCreateUser.AddRow(1)

	rowAddCollections := sqlmock.NewRows([]string{"collection_id"})
	rowAddCollections.AddRow(1)
	rowAddCollections.AddRow(2)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CheckExist)).
		WithArgs(inputUser.Email). // Values in query
		WillReturnRows(rowCheckExist)

	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CreateUser)).
		WithArgs(
			inputUser.Email,
			inputUser.Nickname,
			[]byte(inputUser.Password),
			inputUser.Avatar).
		WillReturnRows(rowCreateUser)

	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CreateDefCollections)).
		WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.CreateUser(ctx, &inputUser)
	require.NotNil(t, err, fmt.Errorf("expected err: %s", sql.ErrNoRows))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, outputUser, actual)
}

func TestPostgres_CreateUser_CollectionsScanErr(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputUser := models.User{
		Email:    "CorrectEmail@mail.ru",
		Password: "CorrectPassword123",
		Nickname: "CorrectNickname",
		Avatar:   "avatar",
	}

	outputUser := models.User{
		ID: 0,
	}

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowCheckExist := sqlmock.NewRows([]string{"exist"})
	rowCheckExist.AddRow(false)

	rowCreateUser := sqlmock.NewRows([]string{"user_id"})
	rowCreateUser.AddRow(1)

	rowAddCollections := sqlmock.NewRows([]string{"collection_id"})
	rowAddCollections.AddRow("oops")
	rowAddCollections.AddRow("spoo")

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CheckExist)).
		WithArgs(inputUser.Email). // Values in query
		WillReturnRows(rowCheckExist)

	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CreateUser)).
		WithArgs(
			inputUser.Email,
			inputUser.Nickname,
			[]byte(inputUser.Password),
			inputUser.Avatar).
		WillReturnRows(rowCreateUser)

	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CreateDefCollections)).
		WillReturnRows(rowAddCollections)
	mock.ExpectRollback()

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.CreateUser(ctx, &inputUser)
	require.NotNil(t, err, fmt.Errorf("expected err: %s", sql.ErrNoRows))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, outputUser, actual)
}

func TestPostgres_CreateUser_LinkCollectionsError(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputUser := models.User{
		Email:    "CorrectEmail@mail.ru",
		Password: "CorrectPassword123",
		Nickname: "CorrectNickname",
		Avatar:   "avatar",
	}

	outputUser := models.User{
		ID: 0,
	}

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowCheckExist := sqlmock.NewRows([]string{"exist"})
	rowCheckExist.AddRow(false)

	rowCreateUser := sqlmock.NewRows([]string{"user_id"})
	rowCreateUser.AddRow(1)

	rowAddCollections := sqlmock.NewRows([]string{"collection_id"})
	rowAddCollections.AddRow(1)
	rowAddCollections.AddRow(2)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CheckExist)).
		WithArgs(inputUser.Email). // Values in query
		WillReturnRows(rowCheckExist)

	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CreateUser)).
		WithArgs(
			inputUser.Email,
			inputUser.Nickname,
			[]byte(inputUser.Password),
			inputUser.Avatar).
		WillReturnRows(rowCreateUser)

	mock.
		ExpectQuery(regexp.QuoteMeta(auth.CreateDefCollections)).
		WillReturnRows(rowAddCollections)

	mock.ExpectExec(regexp.QuoteMeta(auth.LinkUserDefCollections)).
		WithArgs(1, 1, 2).
		WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.CreateUser(ctx, &inputUser)
	require.NotNil(t, err, fmt.Errorf("expected err: %s", sql.ErrNoRows))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, outputUser, actual)
}

func TestPostgres_GetUserByEmail_OK(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputEmail := "correctemail@gmail.com"

	outputUser := models.User{
		ID:       1,
		Email:    inputEmail,
		Nickname: "CorrectNickname",
		Password: "CorrectPassword123",
		Avatar:   "avatar",
	}

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowGetUser := sqlmock.NewRows([]string{"user_id", "email", "nickname", "password", "avatar"})
	rowGetUser.AddRow(
		outputUser.ID,
		outputUser.Email,
		outputUser.Nickname,
		[]byte(outputUser.Password),
		outputUser.Avatar)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.GetUserByEmail)).
		WithArgs(inputEmail).
		WillReturnRows(rowGetUser)

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.GetUserByEmail(ctx, inputEmail)
	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, outputUser, actual)
}

func TestPostgres_GetUserByEmail_RowUserError(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputEmail := "correctemail@gmail.com"

	outputUser := models.User{
		ID:       0,
		Email:    "",
		Nickname: "",
		Password: "",
		Avatar:   "",
	}

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowGetUser := sqlmock.NewRows([]string{"user_id", "email", "nickname", "password", "avatar"})
	rowGetUser.AddRow(
		outputUser.ID,
		outputUser.Email,
		outputUser.Nickname,
		[]byte(outputUser.Password),
		outputUser.Avatar)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.GetUserByEmail)).
		WithArgs(inputEmail).
		WillReturnError(sql.ErrNoRows)

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.GetUserByEmail(ctx, inputEmail)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", sql.ErrNoRows))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, outputUser, actual)
}

func TestPostgres_GetUserByEmail_ScanError(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputEmail := "correctemail@gmail.com"

	outputUser := models.User{
		ID:       0,
		Email:    "",
		Nickname: "",
		Password: "",
		Avatar:   "",
	}

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowGetUser := sqlmock.NewRows([]string{"user_id", "email", "nickname", "password", "avatar"})
	rowGetUser.AddRow(
		"oops",
		outputUser.Email,
		outputUser.Nickname,
		[]byte(outputUser.Password),
		outputUser.Avatar)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.GetUserByEmail)).
		WithArgs(inputEmail).
		WillReturnRows(rowGetUser)

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.GetUserByEmail(ctx, inputEmail)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, outputUser, actual)
}

func TestPostgres_GetUserByEmail_AvatarInvalid(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputEmail := "correctemail@gmail.com"

	outputUser := models.User{
		ID:       1,
		Email:    inputEmail,
		Nickname: "CorrectNickname",
		Password: "CorrectPassword123",
		Avatar:   "avatar",
	}

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowGetUser := sqlmock.NewRows([]string{"user_id", "email", "nickname", "password", "avatar"})
	rowGetUser.AddRow(
		outputUser.ID,
		outputUser.Email,
		outputUser.Nickname,
		[]byte(outputUser.Password),
		nil)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.GetUserByEmail)).
		WithArgs(inputEmail).
		WillReturnRows(rowGetUser)

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.GetUserByEmail(ctx, inputEmail)
	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, outputUser, actual)
}

func TestPostgres_GetUserByID_OK(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputID := 1
	outputUser := models.User{
		Email:    "CorrectEmail@gmail.com",
		Nickname: "CorrectNickname",
		Password: "CorrectPassword123",
		Avatar:   "avatar",
	}

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowGetUser := sqlmock.NewRows([]string{"email", "nickname", "password", "avatar"})
	rowGetUser.AddRow(
		outputUser.Email,
		outputUser.Nickname,
		[]byte(outputUser.Password),
		outputUser.Avatar)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.GetUserByID)).
		WithArgs(inputID).
		WillReturnRows(rowGetUser)

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.GetUserByID(ctx, inputID)
	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, outputUser, actual)
}

func TestPostgres_GetUserByID_RowUserError(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputID := 1
	outputUser := models.User{
		Email:    "",
		Nickname: "",
		Password: "",
		Avatar:   "",
	}

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowGetUser := sqlmock.NewRows([]string{"email", "nickname", "password", "avatar"})
	rowGetUser.AddRow(
		outputUser.Email,
		outputUser.Nickname,
		[]byte(outputUser.Password),
		outputUser.Avatar)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.GetUserByID)).
		WithArgs(inputID).
		WillReturnError(sql.ErrNoRows)

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.GetUserByID(ctx, inputID)
	require.NotNil(t, err, fmt.Errorf("expected err: %s", sql.ErrNoRows))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, outputUser, actual)
}

func TestPostgres_GetUserByID_ScanError(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputID := 1
	outputUser := models.User{
		Email:    "",
		Nickname: "",
		Password: "",
		Avatar:   "",
	}

	// Input global
	ctx := context.TODO()

	// Create required setup for handling
	rowGetUser := sqlmock.NewRows([]string{"nickname", "password", "avatar"})
	rowGetUser.AddRow(
		nil,
		nil,
		nil)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(auth.GetUserByID)).
		WithArgs(inputID).
		WillReturnRows(rowGetUser)

	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	actual, err := repo.GetUserByID(ctx, inputID)
	require.NotNil(t, err, fmt.Errorf("expected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// CheckNewNotification actual
	require.Equal(t, outputUser, actual)
}

func TestPostgres_UpdatePassword_OK(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputUser := models.User{
		ID: 1,
	}
	inputPassword := "newpassword123"

	// Input global
	ctx := context.TODO()

	// Settings mock
	mock.ExpectBegin()
	mock.
		ExpectExec(regexp.QuoteMeta(auth.UpdateUserPassword)).
		WithArgs([]byte(inputPassword), inputUser.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	err = repo.UpdatePassword(ctx, &inputUser, inputPassword)
	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestPostgres_UpdatePassword_ExecError(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputUser := models.User{
		ID: 1,
	}
	inputPassword := "newpassword123"

	// Input global
	ctx := context.TODO()

	// Settings mock
	mock.ExpectBegin()
	mock.
		ExpectExec(regexp.QuoteMeta(auth.UpdateUserPassword)).
		WithArgs([]byte(inputPassword), inputUser.ID).
		WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()
	// Init
	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})

	// CheckNewNotification result
	err = repo.UpdatePassword(ctx, &inputUser, inputPassword)
	require.NotNil(t, err, fmt.Errorf("expected err: %s", sql.ErrNoRows))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}
