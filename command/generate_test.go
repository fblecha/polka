package command

import (
	"fmt"
	"os"
	"testing"
)

func TestTemplatesExist(t *testing.T) {
	endpoints := GetEndpoints()
	//see if each of these templates exist
	for _, endpoint := range endpoints {
		templateFile := fmt.Sprintf("../templates/js/endpoint/%v", endpoint)
		if _, err := os.Stat(templateFile); err != nil {
			t.Error("unable to find template", templateFile)
		}
	}
}
