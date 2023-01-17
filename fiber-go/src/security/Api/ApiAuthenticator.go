package api

import (
	"context"
	"fmt"
	firebase "marketplace/config/services/Firebase"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func MiddlewareAuthenticatorApi(c *fiber.Ctx) error {

	headers := c.GetReqHeaders()

	authorization := headers["Authorization"]

	if authorization == "" {
		c.SendStatus(http.StatusUnauthorized)
		return nil
	}

	connectionFirebase := firebase.NewConnectionFirebase()
	token, err := connectionFirebase.Client.GetUser(context.Background(), authorization)

	if err != nil {
		fmt.Printf("Error verify id: %v", err)
		c.SendStatus(http.StatusUnauthorized)
	}

	fmt.Printf("Verified ID token: %s \n", token.ProviderID)

	c.Next()
	return nil
}
