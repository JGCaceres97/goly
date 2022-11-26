package routes

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jgcaceres97/goly/app/controllers"
	"github.com/jgcaceres97/goly/app/model"
	"github.com/jgcaceres97/goly/app/utils"
)

func Redirect(c *fiber.Ctx) error {
	golyUrl := c.Params("redirect")

	goly, err := controllers.FindByGolyUrl(&golyUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not find goly redirect: " + err.Error(),
		})
	}

	goly.Clicked += 1
	err = controllers.UpdateGoly(&goly)
	if err != nil {
		log.Printf("error updating: %v\n", err)
	}

	return c.Redirect(goly.Redirect, fiber.StatusTemporaryRedirect)
}

func GetAllRedirects(c *fiber.Ctx) error {
	golies, err := controllers.GetAllGolies()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all goly links: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(golies)
}

func GetGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id: " + err.Error(),
		})
	}

	goly, err := controllers.GetGoly(&id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error could not retrieve goly from db: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func CreateGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var goly model.Goly
	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON: " + err.Error(),
		})
	}

	for {
		if goly.IsRandom {
			goly.Goly = utils.RandomURL(8)
		}

		err = controllers.CheckGoly(&goly.Goly)
		if err == nil {
			break
		}

		if !goly.IsRandom {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	err = controllers.CreateGoly(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create goly in db: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(goly)
}

func UpdateGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var goly model.Goly
	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse JSON: " + err.Error(),
		})
	}

	err = controllers.UpdateGoly(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not update goly link in db: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func DeleteGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse id from url: " + err.Error(),
		})
	}

	err = controllers.DeleteGoly(&id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not delete from db: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "goly deleted.",
	})
}
