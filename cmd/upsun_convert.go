package main

import (
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
	app "github.com/upsun/convsun"
	logic "github.com/upsun/convsun/internal/logic"
	utils "github.com/upsun/lib-sun/utility"
)

const (
	APP_NAME = "convsun"

	ARG_SRC = "src"
	ARG_DST = "dst"
)

func init() {
	flag.StringVarP(&app.ArgsC.ProjectSource, ARG_SRC, "", "./", "Source project path to convert")
	flag.StringVarP(&app.ArgsC.ProjectDestination, ARG_DST, "", "", "Destination project path where converted")
	flag.StringVarP(&app.ArgsC.TypeMount, "mount_type", "", "storage", "Change 'Local' mount to upsun compatible mode : storage or instance.")
	flag.BoolVarP(&app.Args.Verbose, "verbose", "v", false, "Enable verbose mode")
	flag.CommandLine.SortFlags = false
	flag.Parse()
}

func main() {
	utils.InitLogger(APP_NAME)
	utils.Disclaimer(APP_NAME)
	utils.StartReporters(APP_NAME)

	// Get Arguments
	app.ArgsC.ProjectSource = utils.RequireFlag(
		ARG_SRC,
		"Enter the project path to convert [%v]: ",
		app.ArgsC.ProjectSource,
		false)

	// Add default path if dst argument is not define.
	if app.ArgsC.ProjectDestination == "" {
		app.ArgsC.ProjectDestination = utils.NormalizePath(filepath.Join(app.ArgsC.ProjectSource, ".upsun"))
	}

	// Make folder and generate convert.
	err := os.MkdirAll(app.ArgsC.ProjectDestination, os.ModePerm)
	if err != nil {
		fmt.Printf("Error when created .upsun folder : %q", err)
	} else {
		logic.Convert(app.ArgsC.ProjectSource, app.ArgsC.ProjectDestination)
	}
}
