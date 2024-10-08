package routes

import (
	"errors"

	"github.com/Akkshatt/fiber_go/database"
	"github.com/Akkshatt/fiber_go/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product
	if err != nil {
		return c.Status(400).JSON("Please ensure that  :id is an Integer")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(fiber.StatusOK).JSON(err.Error())

	}
	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusOK).JSON(err.Error())
	}
	return c.Status(200).SendString("sucessfully deleted product")

}

func CreateResponseProduct(product models.Product) Product {
	return Product{ID: product.ID, Name: product.Name, SerialNumber: product.SerialNumber}

}
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusOK).JSON(err.Error())
	}
	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(fiber.StatusOK).JSON(responseProduct)

}
func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.Database.Db.Find(&products)
	responseProducts := []Product{}
	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)

	}
	return c.Status(200).JSON(responseProducts)
}

func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("product does not exist")

	}
	return nil
}
func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product
	if err != nil {
		return c.Status(fiber.StatusOK).JSON("please ensure thhat id is an int")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(fiber.StatusOK).JSON(err.Error())
	}
	responseProduct := CreateResponseProduct(product)
	return c.Status(fiber.StatusOK).JSON(responseProduct)
}
func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product
	if err != nil {
		return c.Status(fiber.StatusOK).JSON("please ensure that id is an Integer")

	}

	err = findProduct(id, &product)

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(err.Error())
	}

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateData UpdateProduct
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusOK).JSON(err.Error())
	}
	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber
	database.Database.Db.Save(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(fiber.StatusOK).JSON(responseProduct)
}
