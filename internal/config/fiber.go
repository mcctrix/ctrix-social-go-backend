package config

import "github.com/gofiber/fiber/v3"

const (
	BODYDATALIMIT = 25 // in MB
)

var Fiber fiber.Config = fiber.Config{
	TrustProxy: true,
	TrustProxyConfig: fiber.TrustProxyConfig{
		Proxies:  []string{"127.0.0.1", "0.0.0.0"},
		Loopback: true,
	},
	BodyLimit: BODYDATALIMIT * 1024 * 1024,
}
