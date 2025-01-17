package routes

import (
	"errors"

	"github.com/Akkshatt/fiber_go/database"
	"github.com/Akkshatt/fiber_go/models"
	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID      uint    `json:"id"`
	User    User    `json:"user"`
	Product Product `json:"product"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{ID: order.ID, User: user, Product: product}
}
func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusOK).JSON(err.Error())

	}
	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(fiber.StatusOK).JSON(err.Error())
	}
	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(fiber.StatusOK).JSON(err.Error())
	}
	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order,responseUser,responseProduct)
	return c.Status(200).JSON(responseOrder)

}
func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)
	responseOrders := []Order{}

	for _,order := range orders {
		var user models.User
		var product models.Product
		database.Database.Db.Find(&user,"id=?",order.UserRefer)
		database.Database.Db.Find(&product,"id=?",order.ProductRefer)
		responseOrder := CreateResponseOrder(order,CreateResponseUser(user),CreateResponseProduct(product))
        responseOrders = append(responseOrders,responseOrder)
	}
	return c.Status(fiber.StatusOK).JSON(responseOrders)
}


func FindOrder(id int,order *models.Order) error {
	database.Database.Db.Find(&order,"id=?",id)
	if order.ID == 0  {
		return errors.New("orders does not exist")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id,err :=c.ParamsInt("id")
	var order models.Order
	if err!= nil{
		return c.Status(fiber.StatusOK).JSON("id is not an integer")
	}
	if err:=FindOrder(id ,&order);err!=nil{
		return c.Status(fiber.StatusOK).JSON(err.Error())

	}
	var user models.User
	var product models.Product 
	database.Database.Db.First(&user,order.UserRefer)
	database.Database.Db.First(&product,order.ProductRefer)
	responseUser := CreateResponseUser(user)
	responseProduct:= CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order,responseUser,responseProduct)
	return c.Status(fiber.StatusOK).JSON(responseOrder)
}