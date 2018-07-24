package state

import (
	"github.com/MRushtonScratch/terrahawk/internal/domain/entity"
	"strings"
	"fmt"
	"strconv"
)

func BuildImportResources(resourceIds []string, state entity.State) map[string]string {
	resources := make(map[string]string)

	var path, resourceId string


	for _, module := range state.Modules {

		path = ""

		if len(module.Path) > 1 {
			path = "module." + strings.Join(module.Path[1:], ".module.")
		}

		for key, resource := range module.Resources {
			resourceId = key

			if len(path) > 0 {
				resourceId = fmt.Sprintf("%v.%v", path, key)
			}

			resources[resourceId] = resource.Primary.ID
		}
	}

	for key, val := range resources {
		fmt.Println(fmt.Sprintf("%v - %v", key, val))
	}

	imports := make(map[string]string)
	for _, resourceId := range resourceIds {

		identifier := convertFromArrayIdentifier(resourceId)
		fmt.Println(fmt.Sprintf("Identifier %v", identifier))
		if len(resources[identifier]) > 0 {
			imports[identifier] = resources[identifier]
		}
	}

	return imports
}

func convertFromArrayIdentifier(id string) string {
	id = strings.Replace(id,"[", ".", 1)
	id = strings.Replace(id, "]", "", 1)
	return id

}

func convertToArrayIdentifier(id string) string {

	segments := strings.Split(id, ".")
	idx := len(segments)-1
	last := segments[idx]

	if _,err := strconv.Atoi(last); err == nil {
		arrId := fmt.Sprintf("[%s]", last)
		id = strings.Join(segments[:idx],".")
		id = fmt.Sprintf("%v%v", id, arrId)
	} else {
		id = strings.Join(segments,".")
	}

	return id
}
