package go_basics

import (
	"algorithms/scopes"
	"fmt"
)

const GlobalConst = 1

var (
	GlobalVariableOne, GlobalVariableTwo = GlobalConst, GlobalConst
)

func TestingGlobalScope() {
	fmt.Println(GlobalConst, GlobalVariableOne, GlobalVariableTwo)
}

func TestingExport() {
	fmt.Println(scopes.ExportedConst, scopes.ExportedVariable)

	scopes.ExportedFunction()
}
