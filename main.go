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

func ScanDocs(path string) []Doc {
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
			each.Contain = ScanDocs(
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
		return c.JSON(ScanDocs("./document"))
	})

	app.Get("dirs", func(c *fiber.Ctx) error {
		var dirs []string
		filepath.Walk("./document", func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				dirs = append(dirs, path)
			}
			return nil
		})
		return c.JSON(dirs)
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
