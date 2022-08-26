package command

import (
	"testing"
)

func TestSQLGenerate(t *testing.T) {
	g := genModelBySQL{
		ddlPath:     "../testdata/test.sql",
		output:      "../testdata/",
		initialisms: false,
		override:    true,
	}
	err := g.Run()
	if err != nil {
		t.Error(err)
		return
	}
}
