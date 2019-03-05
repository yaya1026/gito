package main

import (
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const API_URL = "https://www.gitignore.io/api/"

func main() {
	app := cli.NewApp()
	app.Name = "touch gitignore"
	app.Usage = "This app make .gitignore from gitignore.io"
	app.Version = "1.0.0"


	app.Action = func (context *cli.Context) error {
		tools := context.Args().Get(0)

		macos := ",macos"
		//intellij := ",intellij"
		//vscode := ",visualstudiocode"

		baseReqUrl := API_URL + tools + macos

		resp, _ := http.Get(baseReqUrl)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)

		file, err := os.Create(".gitignore")
		if err != nil {
			log.Fatal(err)
		}

		file.Write(byteArray)

		return nil
	}

	app.Run(os.Args)
}