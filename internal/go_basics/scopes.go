package go_basics

import (
	"fmt"
	"github.com/tranxton/algorithms/internal/scopes"
)

const GlobalConst = 1

var (
	GlobalVariableTwo = GlobalConst
	GlobalVariableOne = GlobalConst
)

func TestingGlobalScope() {
	fmt.Println(GlobalConst, GlobalVariableOne, GlobalVariableTwo)
}

func TestingExport() {
	fmt.Println(scopes.ExportedConst, scopes.ExportedVariable)

	scopes.ExportedFunction()
}
