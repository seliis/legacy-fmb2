package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	fiber "github.com/gofiber/fiber/v2"
)

type Doc struct {
	Dir     bool        `json:"dir"`
	Name    string      `json:"name"`
	Contain interface{} `json:"contain,omitempty"`
}

func ScanDir(path string) []Doc {
	elems, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}

	var each Doc
	var data []Doc

	for _, elem := range elems {
		each.Dir = elem.IsDir()
		each.Name = elem.Name()
		if each.Dir {
			each.Contain = ScanDir(
				path + "/" + each.Name,
			)
		}
		data = append(data, each)
	}
	return data
}

func main() {
	app := fiber.New()

	app.Static("/", "./")

	app.Get("docs", func(c *fiber.Ctx) error {
		return c.JSON(ScanDir("./document"))
	})

	app.Get("docs/:name", func(c *fiber.Ctx) error {
		var found string
		filepath.Walk("./document", func(path string, info os.FileInfo, err error) error {
			if info.Name() == c.Params("name") {
				found = path
			}
			return nil
		})
		return c.SendFile(found)
	})

	app.Listen(":8080")
}
