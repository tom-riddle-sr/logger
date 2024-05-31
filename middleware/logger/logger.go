package logger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func GetLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		logrus.Infof("[req][method => %s][path => %s]", c.Method(), c.Path())
		if c.Method() == fiber.MethodPost {
			logrus.Infof("[req][body => %s]", c.Body())
		}

		defer logrus.Infof("[res][status => %d][path => %s]", c.Response().StatusCode(), c.Path())

		return c.Next()
	}
}
