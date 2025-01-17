package routes

import (
	"errors"

	"github.com/Akkshatt/fiber_go/models"
	"github.com/gofiber/fiber/v2"
	"github.com/Akkshatt/fiber_go/database"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName}

}
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	// Create the user in the database
	if err := database.Database.Db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot create user",
		})
	}

	// Return the created user as a JSON response
	responseUser := CreateResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}
func GetUsers(c *fiber.Ctx) error {
	users :=[]models.User{}
	database.Database.Db.Find(&users)
	responseUsers := []User{}
	for _,user := range users{
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers,responseUser)
	}
	return c.Status(fiber.StatusOK).JSON(responseUsers)
}
func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user,"id=?",id)
	 if user.ID == 0{
		return errors.New("user does not exist")
	 }
	 return nil
}
func GetUser(c *fiber.Ctx) error {
	id ,err := c.ParamsInt("id")
	 var user models.User 
	 if err != nil {
		return c.Status(400).JSON("Please ensure that  :id is an Integer")
	 }
	  if err := findUser(id,&user); err!= nil {
		return c.Status(fiber.StatusOK).JSON(err.Error())

	  }
	  responseUser :=CreateResponseUser(user)
	  return c.Status(fiber.StatusOK).JSON(responseUser)
}

func UpdateUser( c *fiber.Ctx) error {
	id,err := c.ParamsInt("id")
	var user models.User 
	 if err != nil {
		return c.Status(400).JSON("Please ensure that  :id is an Integer")
	 }
	  if err := findUser(id,&user); err!= nil {
		return c.Status(fiber.StatusOK).JSON(err.Error())

	  }
	//   responseUser :=CreateResponseUser(user)
	//   return c.Status(fiber.StatusOK).JSON(responseUser)

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LasttName string `json:"last_name"`
	}
	var updateData UpdateUser
	if err := c.BodyParser(&updateData); err!= nil{
		return c.Status(500).JSON(err.Error())
	}
	user.FirstName = updateData.FirstName
	user.LastName = updateData.LasttName
	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id ,err := c.ParamsInt("id")
	 var user models.User 
	 if err != nil {
		return c.Status(400).JSON("Please ensure that  :id is an Integer")
	 }
	  if err := findUser(id,&user); err!= nil {
		return c.Status(fiber.StatusOK).JSON(err.Error())

	  }
	  if err:= database.Database.Db.Delete(&user).Error; err!= nil{
		return c.Status(fiber.StatusOK).JSON(err.Error())
	  }
	 return c.Status(200).SendString("sucessfully deleted user")
	
	}