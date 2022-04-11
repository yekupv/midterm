package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var (
	storeMutex sync.Mutex
)

var store = map[string]string{
	"1":     "Berik",
	"2":     "Serik",
	"Zerik": "3",
}

func get(c *fiber.Ctx) error {
	storeMutex.Lock()
	fmt.Println(store[c.Params("key")])
	storeMutex.Unlock()
	return c.SendStatus(200)

}
func put(c *fiber.Ctx) error {
	storeMutex.Lock()

	store[c.Params("key")] = c.Params("value")
	fmt.Println("OK")
	storeMutex.Unlock()
	return c.SendStatus(200)

}

func main() {
	app := fiber.New()
	app.Get("store/:key", get)
	app.Get("store/:key/:value", put)
	log.Fatal(app.Listen(":3000"))
}
