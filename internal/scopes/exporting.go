package scopes

import "fmt"

const ExportedConst = 1
const nonExportedConst = 2

var ExportedVariable = ExportedConst
var nonExportedVariable = nonExportedConst

func ExportedFunction() {
	fmt.Println("Меня можно вызвать за пределами пакета,", ExportedVariable)
}

func nonExportedFunction() {
	fmt.Println("Меня нельзя вызвать за пределами пакета, ", nonExportedVariable)
}
