package main

var PublicSlice []byte
var privateArray [20]rune

const PublicFloat64  float64 = 3.1415
const privateFloat32 float32 = 3.14

type privateStruct struct {
	privateField int
	PublicField  int64
}

type publicStruct struct {
	privateField int32
	PublicField  string
}

type myFuncType func(param string) (returnValue int)

func main() {}
