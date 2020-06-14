package grpc_test

import (
	"github.com/moemoe89/go-grpc-server-tisa/api/v1/api_struct/model"
	usrGrpc "github.com/moemoe89/go-grpc-server-tisa/api/v1/user/delivery/grpc"
	usrproto "github.com/moemoe89/go-grpc-server-tisa/api/v1/user/delivery/grpc/proto"
	"github.com/moemoe89/go-grpc-server-tisa/api/v1/user/mocks"

	"errors"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"github.com/stretchr/testify/mock"
	"github.com/rs/xid"
)

func TestDeliveryCreateFailValidation(t *testing.T) {
	server := grpc.NewServer()
	mockService := new(mocks.Service)
	usrCtrl := usrGrpc.NewAUserServerGrpc(server, mockService)

	createUserRes, err := usrCtrl.Create(context.Background(), &usrproto.UserCreateReq{
		Name:    "John",
		Phone:   "081234567890",
		Email:   "johndoe.com",
		Address: "USA",
	})
	assert.Nil(t, createUserRes)
	assert.Error(t, err)
}

func TestDeliveryCreateFail(t *testing.T) {
	server := grpc.NewServer()
	mockService := new(mocks.Service)
	mockService.On("Create", mock.AnythingOfType("*form.UserForm")).Return(nil, http.StatusInternalServerError, errors.New("Unexpected database error"))
	usrCtrl := usrGrpc.NewAUserServerGrpc(server, mockService)

	createUserRes, err := usrCtrl.Create(context.Background(), &usrproto.UserCreateReq{
		Name:    "John",
		Phone:   "081234567890",
		Email:   "john@doe.com",
		Address: "USA",
	})
	assert.Nil(t, createUserRes)
	assert.Error(t, err)
}

func TestDeliveryCreateSuccess(t *testing.T) {
	user := &model.UserModel{
		ID:      "id",
		Name:    "John",
		Email:   "john@doe.com",
		Phone:   "081234567890",
		Address: "USA",
	}

	server := grpc.NewServer()
	mockService := new(mocks.Service)
	mockService.On("Create", mock.AnythingOfType("*form.UserForm")).Return(user, 0, nil)
	usrCtrl := usrGrpc.NewAUserServerGrpc(server, mockService)

	createUserRes, err := usrCtrl.Create(context.Background(), &usrproto.UserCreateReq{
		Name:    user.Name,
		Phone:   user.Phone,
		Email:   user.Email,
		Address: user.Address,
	})
	assert.NotNil(t, createUserRes)
	assert.NoError(t, err)
}

func TestDeliveryList(t *testing.T) {
	user := &model.UserModel{
		ID:      xid.New().String(),
		Name:    "Momo",
		Email:   "momo@mail.com",
		Phone:   "085640",
		Address: "Indonesia",
	}
	users := []*model.UserModel{}
	users = append(users, user)

	name := "a"
	email := "b"
	phone := "0856"
	createdAtStart := "2020-01-01"
	createdAtEnd := "2020-01-31"

	filter := map[string]interface{}{}
	filter["name"] = "%" + name + "%"
	filter["email"] = "%" + email + "%"
	filter["phone"] = "%" + phone + "%"
	filter["created_at_start"] = createdAtStart
	filter["created_at_end"] = createdAtEnd
	filterCount := filter
	filter["limit"] = 10
	filter["offset"] = 0

	server := grpc.NewServer()
	mockService := new(mocks.Service)
	mockService.On("List", filter, filterCount, "WHERE deleted_at IS NULL AND name LIKE :name AND email LIKE :email AND phone LIKE :phone AND created_at >= :created_at_start AND created_at <= :created_at_end", "created_at DESC", model.UserSelectField).Return(users, 1, 0, nil)
	usrCtrl := usrGrpc.NewAUserServerGrpc(server, mockService)

	listUserRes, err := usrCtrl.List(context.Background(), &usrproto.UsersReq{
		PerPage:        "10",
		Name:           name,
		Email:          email,
		Phone:          phone,
		CreatedAtStart: createdAtStart,
		CreatedAtEnd:   createdAtEnd,
		SelectField:    "id,name,email,phone,address,created_at,updated_at",
	})
	assert.NotNil(t, listUserRes)
	assert.NoError(t, err)
}

func TestDeliveryListFail(t *testing.T) {
	filter := map[string]interface{}{}
	filterCount := filter
	filter["limit"] = 10
	filter["offset"] = 0

	server := grpc.NewServer()
	mockService := new(mocks.Service)
	mockService.On("List", filter, filterCount, "WHERE deleted_at IS NULL", "created_at DESC", model.UserSelectField).Return(nil, 0, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later"))
	usrCtrl := usrGrpc.NewAUserServerGrpc(server, mockService)

	listUserRes, err := usrCtrl.List(context.Background(), &usrproto.UsersReq{
		PerPage: "10",
	})
	assert.Nil(t, listUserRes)
	assert.Error(t, err)
}

func TestDeliveryListFailPagination(t *testing.T) {
	filter := map[string]interface{}{}
	filterCount := filter
	filter["limit"] = 10
	filter["offset"] = 0

	server := grpc.NewServer()
	mockService := new(mocks.Service)
	mockService.On("List", filter, filterCount, "WHERE deleted_at IS NULL", "created_at DESC", model.UserSelectField).Return(nil, 0, http.StatusInternalServerError, errors.New("Invalid parameter per_page: not an int"))
	usrCtrl := usrGrpc.NewAUserServerGrpc(server, mockService)

	listUserRes, err := usrCtrl.List(context.Background(), &usrproto.UsersReq{
		PerPage: "a",
	})
	assert.Nil(t, listUserRes)
	assert.Error(t, err)
}

func TestDeliveryDetailFail(t *testing.T) {
	server := grpc.NewServer()
	mockService := new(mocks.Service)
	mockService.On("Detail", mock.AnythingOfType("string"),  mock.AnythingOfType("string")).Return(nil, http.StatusInternalServerError, errors.New("Unexpected database error"))
	usrCtrl := usrGrpc.NewAUserServerGrpc(server, mockService)

	detailUserRes, err := usrCtrl.Detail(context.Background(), &usrproto.UserIDReq{
		Id: "id",
	})
	assert.Nil(t, detailUserRes)
	assert.Error(t, err)
}

func TestDeliveryDetailSuccess(t *testing.T) {
	user := &model.UserModel{
		ID:      "id",
		Name:    "John",
		Email:   "john@doe.com",
		Phone:   "081234567890",
		Address: "USA",
	}

	server := grpc.NewServer()
	mockService := new(mocks.Service)
	mockService.On("Detail", mock.AnythingOfType("string"),  mock.AnythingOfType("string")).Return(user, 0, nil)
	usrCtrl := usrGrpc.NewAUserServerGrpc(server, mockService)

	detailUserRes, err := usrCtrl.Detail(context.Background(), &usrproto.UserIDReq{
		Id: "id",
	})
	assert.NotNil(t, detailUserRes)
	assert.NoError(t, err)
}

func TestDeliveryUpdateFailValidation(t *testing.T) {
	user := &model.UserModel{
		ID:      "id",
		Name:    "John",
		Email:   "john@doe.com",
		Phone:   "081234567890",
		Address: "USA",
	}

	server := grpc.NewServer()
	mockService := new(mocks.Service)
	usrCtrl := usrGrpc.NewAUserServerGrpc(server, mockService)
	mockService.On("Detail", mock.AnythingOfType("string"),  mock.AnythingOfType("string")).Return(user, 0, nil)

	updateUserRes, err := usrCtrl.Update(context.Background(), &usrproto.UserUpdateReq{
		Id:      "id",
		Name:    "John",
		Phone:   "081234567890",
		Email:   "johndoe.com",
		Address: "USA",
	})
	assert.Nil(t, updateUserRes)
	assert.Error(t, err)
}

func TestDeliveryUpdateFailDetail(t *testing.T) {
	server := grpc.NewServer()
	mockService := new(mocks.Service)
	mockService.On("Detail", mock.AnythingOfType("string"),  mock.AnythingOfType("string")).Return(nil, http.StatusInternalServerError, errors.New("Unexpected database error"))
	usrCtrl := usrGrpc.NewAUserServerGrpc(server, mockService)

	updateUserRes, err := usrCtrl.Update(context.Background(), &usrproto.UserUpdateReq{
		Id:      "id",
		Name:    "John",
		Phone:   "081234567890",
		Email:   "john@doe.com",
		Address: "USA",
	})
	assert.Nil(t, updateUserRes)
	assert.Error(t, err)
}

func TestDeliveryUpdateFail(t *testing.T) {
	user := &model.UserModel{
		ID:      "id",
		Name:    "John",
		Email:   "john@doe.com",
		Phone:   "081234567890",
		Address: "USA",
	}

	server := grpc.NewServer()
	mockService := new(mocks.Service)
	mockService.On("Detail", mock.AnythingOfType("string"),  mock.AnythingOfType("string")).Return(user, 0, nil)
	mockService.On("Update", mock.AnythingOfType("*form.UserForm"), mock.AnythingOfType("string")).Return(nil, http.StatusInternalServerError, errors.New("Unexpected database error"))
	usrCtrl := usrGrpc.NewAUserServerGrpc(server, mockService)

	updateUserRes, err := usrCtrl.Update(context.Background(), &usrproto.UserUpdateReq{
		Id:      "id",
		Name:    "John",
		Phone:   "081234567890",
		Email:   "john@doe.com",
		Address: "USA",
	})
	assert.Nil(t, updateUserRes)
	assert.Error(t, err)
}

func TestDeliveryUpdateSuccess(t *testing.T) {
	user := &model.UserModel{
		ID:      "id",
		Name:    "John",
		Email:   "john@doe.com",
		Phone:   "081234567890",
		Address: "USA",
	}

	server := grpc.NewServer()
	mockService := new(mocks.Service)
	mockService.On("Detail", mock.AnythingOfType("string"),  mock.AnythingOfType("string")).Return(user, 0, nil)
	mockService.On("Update", mock.AnythingOfType("*form.UserForm"), mock.AnythingOfType("string")).Return(user, 0, nil)
	usrCtrl := usrGrpc.NewAUserServerGrpc(server, mockService)

	updateUserRes, err := usrCtrl.Update(context.Background(), &usrproto.UserUpdateReq{
		Id:      user.ID,
		Name:    user.Name,
		Phone:   user.Phone,
		Email:   user.Email,
		Address: user.Address,
	})
	assert.NotNil(t, updateUserRes)
	assert.NoError(t, err)
}

func TestDeliveryDeleteFail(t *testing.T) {
	server := grpc.NewServer()
	mockService := new(mocks.Service)
	mockService.On("Delete", mock.AnythingOfType("string")).Return(http.StatusInternalServerError, errors.New("Unexpected database error"))
	usrCtrl := usrGrpc.NewAUserServerGrpc(server, mockService)

	_, err := usrCtrl.Delete(context.Background(), &usrproto.UserIDReq{
		Id: "id",
	})
	assert.Error(t, err)
}

func TestDeliveryDeleteSuccess(t *testing.T) {
	server := grpc.NewServer()
	mockService := new(mocks.Service)
	mockService.On("Delete", mock.AnythingOfType("string")).Return(0, nil)
	usrCtrl := usrGrpc.NewAUserServerGrpc(server, mockService)

	_, err := usrCtrl.Delete(context.Background(), &usrproto.UserIDReq{
		Id: "id",
	})
	assert.NoError(t, err)
}
