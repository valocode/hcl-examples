package main

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

type TestRun struct {
	Tool  string       `hcl:"tool,attr"`
	Type  string       `hcl:"type,attr"`
	Tests TestCaseEdge `hcl:"tests,block"`
}

type TestCaseEdge struct {
	TestCases []TestCase `hcl:"test_case,block"`
}

type TestCase struct {
	Name        string `hcl:"name,attr"`
	Description string `hcl:"description,optional"`
	Result      bool   `hcl:"result,attr"`
}

type InputDecl struct {
	Name     string         `hcl:",label"`
	TypeExpr hcl.Expression `hcl:"type,optional"`
	Type     cty.Type
}

type InputDef struct {
	Values map[string]cty.Value `hcl:",remain"`
}
