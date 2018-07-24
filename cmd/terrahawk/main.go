package main

import (
	"flag"
	"fmt"
	"github.com/MRushtonScratch/terrahawk/internal/infrastructure/state"
	"github.com/MRushtonScratch/terrahawk/internal/infrastructure/file"
)

func main () {

	stateFile := flag.String("state", "", "Path to the state file")
	resourcesFile := flag.String("resources", "", "Path to the resources file")
	env := flag.String("env", "", "The environment name")

	flag.Parse()
	fmt.Println("Arguments:", flag.Args())

	tfState, err := state.Detach(*stateFile)
	check(err)

	resourceIds, err := file.ReadLines(*resourcesFile)

	imports := state.BuildImportResources(resourceIds, tfState)

	for key, id := range imports {
		fmt.Println(fmt.Sprintf("./working/import_state.sh \"%s\" %s %s", *env, key, id))

	}


	fmt.Printf("Resources %v matched to %v resourceIds", len(imports), len(resourceIds))
	fmt.Println("")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}