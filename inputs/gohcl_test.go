package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/ext/typeexpr"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"
)

func TestGohcl(t *testing.T) {
	// We need a "wrapper" for the HCL structs
	var hclWrap struct {
		TestRun TestRun `hcl:"test_run,block"`
	}
	// hclsimple is just a wrapper around gohcl
	err := hclsimple.DecodeFile("test_run.hcl", nil, &hclWrap)
	require.NoError(t, err)
	// Dump the output
	spew.Dump(hclWrap)
}

func TestGohclWithInputs(t *testing.T) {
	const filename = "test_run_with_inputs.hcl"
	// We need a "wrapper" for the HCL structs
	var inputs struct {
		InputDecls []InputDecl `hcl:"input,block"`
		// Make it optional with pointer
		InputDefs *InputDef `hcl:"inputs,block"`
		// Store the rest of the stuff here
		Leftovers hcl.Body `hcl:",remain"`
	}
	// hclsimple is just a wrapper around gohcl
	err := hclsimple.DecodeFile(filename, nil, &inputs)
	require.NoError(t, err)

	// Convert the type expressions into actual cty.Types
	for idx := range inputs.InputDecls {
		var (
			// Get a pointer to the decl so that we can update it
			decl  = &inputs.InputDecls[idx]
			diags hcl.Diagnostics
		)
		decl.Type, diags = typeexpr.TypeConstraint(decl.TypeExpr)
		require.Empty(t, diags.Errs())
	}

	// Validate the inputs
	if inputs.InputDefs != nil {
		for name, def := range inputs.InputDefs.Values {
			var defFound bool
			for _, decl := range inputs.InputDecls {
				if name == decl.Name {
					defFound = true
					// Check the type
					assert.Equal(t, def.Type(), decl.Type)
					continue
				}
			}
			// Check that the input definition was declared
			assert.True(t, defFound)
		}
	}

	// Create the eval context
	evalCtx := hcl.EvalContext{
		Variables: map[string]cty.Value{
			"input": cty.ObjectVal(inputs.InputDefs.Values),
		},
	}

	var hclWrap struct {
		TestRun TestRun `hcl:"test_run,block"`
	}
	diags := gohcl.DecodeBody(inputs.Leftovers, &evalCtx, &hclWrap)
	require.Empty(t, diags.Errs())
	// Dump the output
	spew.Dump(hclWrap.TestRun)
}
