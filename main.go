package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	fiber "github.com/gofiber/fiber/v2"
)

type Doc struct {
	Dir     bool        `json:"dir"`
	Name    string      `json:"name"`
	Contain interface{} `json:"contain,omitempty"`
}

type Post struct {
	Directory string `json:"directory"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
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

	app.Get("doc/docs", func(c *fiber.Ctx) error {
		return c.JSON(ScanDocs("./document"))
	})

	app.Get("doc/dirs", func(c *fiber.Ctx) error {
		var dirs []string
		filepath.Walk("./document", func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				dirs = append(dirs, path)
			}
			return nil
		})
		return c.JSON(dirs)
	})

	app.Get("doc/getdoc/:name", func(c *fiber.Ctx) error {
		var found string
		filepath.Walk("./document", func(path string, info os.FileInfo, err error) error {
			if info.Name() == c.Params("name")+".html" {
				found = path
			}
			return nil
		})
		if found != "" {
			return c.SendFile(found)
		} else {
			return c.SendStatus(http.StatusNotFound)
		}
	})

	app.Post("doc/send", func(c *fiber.Ctx) error {
		recv := new(Post) // why var declare make pointer error?
		if err := c.BodyParser(recv); err != nil {
			return err
		}
		file, err := os.Create(recv.Directory + "/" + recv.Subject + ".html")
		if err != nil {
			return err
		}
		defer file.Close()
		file.WriteString(recv.Content)
		return c.SendStatus(http.StatusOK)
	})

	app.Listen(":8080")
}
