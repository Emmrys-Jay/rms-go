package user

import (
	"github.com/gofiber/fiber/v2"
	"net/http"

	"gateway-service/utility"
)

func (base *Controller) CreateUser(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) GetUser(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) GetAllUsers(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) UpdateUser(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) DeleteUser(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) AddProfilePicture(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}
