package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/th002/goswagger/swagger"
)

//Version application version
const Version = "0.1.0"

var version = flag.Bool("v", false, "show version")
var apiPackage = flag.String("apiPackage", "", "The package that implements the API controllers, relative to $GOPATH/src")
var mainFile = flag.String("mainFile", "", "The file that contains the general API annotations, relative to $GOPATH/src")
var outPutPath = flag.String("output", "./swagger.json", "Output (path) for the generated file(s)")
var ignore = flag.String("ignore", "swagger", "Ignore packages that satisfy this match")

func main() {
	flag.Parse()
	if *version {
		fmt.Printf("go-swagger verison:%s\nswagger version:%s\ngo version:%s\n", Version, swagger.SwaggerVersion, runtime.Version())
		return
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ignores := strings.Split(*ignore, ";")
	params := swagger.Params{
		OutputPath:  *outPutPath,
		Ignore:      ignores,
		APIPackage:  *apiPackage,
		MainAPIFile: *mainFile,
	}
	if err := swagger.Run(params); err != nil {
		log.Fatalln(err)
	}
}
