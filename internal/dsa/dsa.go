package dsa

type Parameter struct {
	ParamName string
	ParamType string
}

type Dsa struct {
	Name         string
	FunctionName string
	Parameters   []Parameter
	ReturnType   string
}

var Arr = []Dsa{
	{
		Name:         "LinearSearch",
		FunctionName: "linear_search",
		Parameters: []Parameter{
			{
				ParamName: "arr",
				ParamType: "[]int",
			},
			{
				ParamName: "target",
				ParamType: "int",
			},
		},
		ReturnType: "bool",
	},
	{
		Name:         "TwoCrystalBalls",
		FunctionName: "two_crystal_balls",
		Parameters: []Parameter{
			{
				ParamName: "breaks",
				ParamType: "[]bool",
			},
		},
		ReturnType: "int",
	},
}
