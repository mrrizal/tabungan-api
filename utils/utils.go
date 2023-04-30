package utils

import (
	"fmt"
	"math/rand"
	"tabungan-api/app/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GenerateAccountNumber(length int) string {
	rand.Seed(time.Now().UnixNano())
	min := int64(pow(10, length-1))
	max := int64(pow(10, length) - 1)
	number := rand.Int63n(max-min+1) + min
	return fmt.Sprintf("%0*d", length, number)
}

func pow(x, y int) int {
	p := 1
	for i := 0; i < y; i++ {
		p *= x
	}
	return p
}

func ErrorResp(c *fiber.Ctx, message string) error {
	var result models.ErrorResponse
	result.Remark = message
	c.Status(400)
	return c.JSON(result)
}
