package constValues

type ConfigValueType string

const (
	Array   ConfigValueType = "Array"
	Boolean ConfigValueType = "Boolean"
	Object  ConfigValueType = "Object"
	String  ConfigValueType = "String"
	Number  ConfigValueType = "Number"
)

func (e ConfigValueType) String() string {
	types := [...]string{"Array", "Boolean", "Object", "String", "Number"}

	x := string(e)
	for _, v := range types {
		if v == x {
			return x
		}
	}

	return ""
}
