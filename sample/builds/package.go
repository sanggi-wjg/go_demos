package builds

import (
	"C"
	"encoding/json"
	"fmt"
)

/*
build 이전에 반드시 할 것
1. 함수들은 main package로 옮길 것
2. import "C" 를 추가할

go build -buildmode=c-shared -o packages.so main.go

https://pkg.go.dev/cmd/go#Compile_packages_and_dependencies
https://fluhus.github.io/snopher/
*/

//export TestPing
func TestPing() {
	fmt.Println("Pong")
}

//export Add
func Add(a, b int) int {
	return a + b
}

//export JsonDecode
func JsonDecode(target []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal(target, &data); err != nil {
		return nil, err
	}
	return data, nil
}
