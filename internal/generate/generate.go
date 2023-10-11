package generate

import (
	"fmt"
	"os"

	"github.com/agp745/Go-kata-machine/internal/dsa"
)

func createFile(fileName string, functionName string, parameters []dsa.Parameter, returnType string) {

	var params string
	for i, param := range parameters {
		params += param.ParamName + " " + param.ParamType
		if i < len(parameters)-1 {
			params += ", "
		}
	}

	function := "func " + functionName + "(" + params + ") " + returnType + " {}"

	body := "package dsa\n\n" + function

	err := os.WriteFile("./internal/dsa/"+fileName+".go", []byte(body), 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func Generate() {

	for _, data := range dsa.Arr {
		createFile(data.Name, data.FunctionName, data.Parameters, data.ReturnType)
	}

}
