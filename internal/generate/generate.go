package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/agp745/Go-kata-machine/internal/dsa"
)

func createDir(fileName string) {

	path := "./day1/" + fileName

	if err := os.Mkdir(path, 0755); err != nil {
		println("Error creating directory...\n", err)
	}

}

func createFile(fileName string, functionName string, parameters []dsa.Parameter, returnType string) {

	var params string
	for i, param := range parameters {
		params += param.ParamName + " " + param.ParamType
		if i < len(parameters)-1 {
			params += ", "
		}
	}

	function := "func " + functionName + "(" + params + ") " + returnType + " {}"

	body := "package " + strings.ToLower(fileName) + "\n\n" + function

	err := os.WriteFile("./day1/"+fileName+"/"+fileName+".go", []byte(body), 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func Generate() {

	if err := os.Mkdir("./day1", 0755); err != nil {
		println("Error creating day1 directory...\n", err)
	}

	for _, data := range dsa.Arr {
		createDir(data.Name)
		createFile(data.Name, data.FunctionName, data.Parameters, data.ReturnType)
	}

}
