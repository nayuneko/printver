package printver

import "fmt"

var VER = ""
var NoSetVer = "<no set>"

func Version() {
	fmt.Println(GetVersion())
}

func GetVersion() string {
	if VER == "" {
		return NoSetVer
	}
	return VER
}
