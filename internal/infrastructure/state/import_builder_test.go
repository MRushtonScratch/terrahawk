package state

import (
"testing"
"path/filepath"
)

func TestImportBuilderSimple(t *testing.T) {

	path := filepath.Join("testData", "simple.tfstate")
	state, err := Detach(path)

	if err != nil {
		t.Error(err)
	}

	ids := []string{"module.bootstrap.aws_route53_record.oasis-consul-bootstrap-a"}

	result := BuildImportResources(ids, state)

	if len(result) != 1 {
		t.Error("A single import statement was expected")
	}
}

func TestImportBuilderArrays(t *testing.T) {

	path := filepath.Join("testData", "simple.tfstate")
	state, err := Detach(path)

	if err != nil {
		t.Error(err)
	}

	ids := []string{"module.bootstrap.aws_route53_zone.oasis-consul-bootstrap[0]"}

	result := BuildImportResources(ids, state)

	if len(result) != 1 {
		t.Error("A single import statement was expected")
	}
}