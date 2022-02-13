package cli

import (
	"fmt"

	"github.com/edmarfelipe/go-hexagonal/application"
)

func Create(service application.ProductServiceInterface, productName string, price float64) (string, error) {
	var result = ""

	product, err := service.Create(productName, price)
	if err != nil {
		return result, err
	}

	result = fmt.Sprintf(
		"Product ID %s with the name %s has been created with the price %f and status %s",
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)

	return result, nil
}

func Enable(service application.ProductServiceInterface, productId string) (string, error) {
	var result = ""

	product, err := service.Get(productId)
	if err != nil {
		return result, err
	}
	res, err := service.Enable(product)
	if err != nil {
		return result, err
	}

	result = fmt.Sprintf(
		"Product ID %s has been enabled",
		res.GetName(),
	)

	return result, nil
}

func Disable(service application.ProductServiceInterface, productId string) (string, error) {
	var result = ""

	product, err := service.Get(productId)
	if err != nil {
		return result, err
	}
	res, err := service.Disable(product)
	if err != nil {
		return result, err
	}

	result = fmt.Sprintf(
		"Product ID %s has been disabled",
		res.GetName(),
	)

	return result, nil
}

func Get(service application.ProductServiceInterface, productId string) (string, error) {
	var result = ""

	product, err := service.Get(productId)
	if err != nil {
		return result, err
	}

	text := "Product ID: %s\nName: %s\nPrice %f\nStatus: %s"
	result = fmt.Sprintf(
		text,
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)

	return result, nil
}
