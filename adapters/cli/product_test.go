package cli_test

import (
	"fmt"
	"testing"

	"github.com/edmarfelipe/go-hexagonal/adapters/cli"
	"github.com/edmarfelipe/go-hexagonal/application"

	mock_application "github.com/edmarfelipe/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func createMock(t *testing.T) (*mock_application.MockProductServiceInterface, *application.Product) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25.99

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(product.ID).AnyTimes()
	productMock.EXPECT().GetName().Return(product.Name).AnyTimes()
	productMock.EXPECT().GetStatus().Return(product.Status).AnyTimes()
	productMock.EXPECT().GetPrice().Return(product.Price).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(product.Name, product.Price).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(product.ID).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	return service, product
}

func TestRun_Create(t *testing.T) {
	service, product := createMock(t)

	resultExpected := fmt.Sprintf(
		"Product ID %s with the name %s has been created with the price %f and status %s",
		product.ID,
		product.Name,
		product.Price,
		product.Status,
	)

	result, err := cli.Create(service, product.Name, product.Price)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}

func TestRun_Disable(t *testing.T) {
	service, product := createMock(t)
	resultExpected := fmt.Sprintf("Product ID %s has been disabled", product.Name)
	result, err := cli.Disable(service, product.ID)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}

func TestRun_Enable(t *testing.T) {
	service, product := createMock(t)
	resultExpected := fmt.Sprintf("Product ID %s has been disabled", product.Name)
	result, err := cli.Disable(service, product.ID)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}

func TestRun_Get(t *testing.T) {
	service, product := createMock(t)

	resultExpected := fmt.Sprintf(
		"Product ID: %s\nName: %s\nPrice %f\nStatus: %s",
		product.ID,
		product.Name,
		product.Price,
		product.Status,
	)

	result, err := cli.Get(service, product.ID)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
