package main

import (
	routers "Sample/pkg/routers"
	middleware "Sample/pkg/utils"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	// "io/ioutil"
// 	"net/http"

// 	"github.com/gofiber/fiber/v2"
// )

func main() {

	// app := fiber.New()
	// app.Post("/", func(c *fiber.Ctx) error {

	// 	m := make(map[string]interface{})
	// 	c.BodyParser(&m)
	// 	// fmt.Println(c.Request().Header.Header()("Signature"))
	// 	// fmt.Println(c.Get("Signature"))

	// 	marshallReq, _ := json.Marshal(m)
	// 	req, _ := http.NewRequest("POST", "https://test.atmprepaidsolutions.com/plr0422/api/v1/balance", bytes.NewBuffer(marshallReq))

	// 	auth := "FCCeLp3ZrBuf9WJNIxV35yCdcXRtyeqCjt9aQP345ckrXGkPo7Zh78ykuCscbBKmQu4g8xiDV0pNFG26t3moAs+srIdbpel19JlxYc8jNmFG+2uXCpL1IFFnKRrcLgoPql4jDxw+m1wYSCNVOT6TeHn0XIq0MTX+ltZYXAKCWZZZGaaoQ/tgAzTCjLepDI/Ke/560m7QP6VmPuTNU56NVdD0AlEcR8LwQDgIeRq0CtwCM/pKswc/aEN6JlWN2PPaqnq+EphZHifYfivhLaIcfSEJtrVxjd+lACfQ7e1w11zm7Qd1NgJ+wljpkVPgOQSk4PPzGGkV8H1YSYXH96Z1oQ=="
	// 	// req.Header = http.Header{
	// 	// 	"Signature":    {auth},
	// 	// 	"Content-Type": {"application/json"},
	// 	// 	"charset":      {"utf-8"},
	// 	// }
	// 	req.Header.Set("Signature", auth)
	// 	fmt.Println(req.Header["Signature"])

	// 	// client := &http.Client{}
	// 	// resp, _ := client.Do(req)

	// 	// defer resp.Body.Close()

	// 	// body, _ := ioutil.ReadAll(resp.Body)

	// 	// result := make(map[string]interface{})
	// 	// json.Unmarshal(body, &result)
	// 	// fmt.Println(result)
	// 	return c.JSON("result")
	// })
	// app.Listen(":2000")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading Env File: ", err)
	}
	envi := os.Getenv("ENVIRONMENT")

	err = godotenv.Load(fmt.Sprintf(".env-%v", envi)) //
	if err != nil {
		log.Fatal("Error Loading Env File: ", err)
	}

	// Initialize DB here

	// Declare & initialize fiber
	app := fiber.New(fiber.Config{
		UnescapePath: true,
	})

	// For GoRoutine implementation
	// appb := fiber.New(fiber.Config{
	// 	UnescapePath: true,
	// })

	// Configure application CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// For GoRoutine implementation
	// appb.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	// }))

	// Declare & initialize logger
	app.Use(logger.New())

	// For GoRoutine implementation
	// appb.Use(logger.New())

	// Declare & initialize routes
	routers.SetupPublicRoutes(app)
	routers.SetupPrivateRoutes(app)

	// For GoRoutine implementation
	// routers.SetupPublicRoutesB(appb)
	// go func() {
	// 	err := appb.Listen(fmt.Sprintf(":8002"))
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// }()

	fmt.Println("Port: ", middleware.GetEnv("PORT"))
	// Serve the application
	if middleware.GetEnv("SSL") == "enabled" {
		log.Fatal(app.ListenTLS(
			fmt.Sprintf(":%s", middleware.GetEnv("PORT")),
			middleware.GetEnv("SSL_CERTIFICATE"),
			middleware.GetEnv("SSL_KEY"),
		))
	} else {
		err := app.Listen(fmt.Sprintf(":%s", middleware.GetEnv("PORT")))
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
