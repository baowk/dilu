package module

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"

	"github.com/baowk/dilu-core/common/utils"

	"github.com/spf13/cobra"
)

var (
	modName  string
	StartCmd = &cobra.Command{
		Use:     "new",
		Short:   "Create a new module,Path:${project}/module/",
		Long:    "Use when you need to create a new module",
		Example: "dilu new -n yourname",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&modName, "name", "n", "", "Start server with provided configuration file")
}

func run() {

	fmt.Println(`start init`)
	//1. 读取配置

	fmt.Println(`generate migration file`)
	_ = genFile()

}

const ROOT = "modules/"

func genFile() error {
	if modName == "" {
		return errors.New("arg `name` invalid ：name is empty")
	}
	path := ROOT
	appPath := path + modName
	err := utils.IsNotExistMkDir(appPath)
	if err != nil {
		return err
	}
	apiPath := appPath + "/apis/"
	err = utils.IsNotExistMkDir(apiPath)
	if err != nil {
		return err
	}
	modelsPath := appPath + "/models/"
	err = utils.IsNotExistMkDir(modelsPath)
	if err != nil {
		return err
	}
	routerPath := appPath + "/router/"
	err = utils.IsNotExistMkDir(routerPath)
	if err != nil {
		return err
	}
	servicePath := appPath + "/service/"
	err = utils.IsNotExistMkDir(servicePath)
	if err != nil {
		return err
	}
	dtoPath := appPath + "/service/dto/"
	err = utils.IsNotExistMkDir(dtoPath)
	if err != nil {
		return err
	}

	t1, err := template.ParseFiles("resources/template/cmd_api.template")
	if err != nil {
		return err
	}
	m := map[string]string{}
	m["modName"] = modName
	var b1 bytes.Buffer
	err = t1.Execute(&b1, m)
	if err != nil {
		return err
	}
	utils.FileCreate(b1, "cmd/start/"+modName+".go")
	t2, err := template.ParseFiles("resources/template/router.template")
	if err != nil {
		return err
	}
	var b2 bytes.Buffer
	err = t2.Execute(&b2, m)
	if err != nil {
		return err
	}
	utils.FileCreate(b2, appPath+"/router/router.go")
	return nil
}
