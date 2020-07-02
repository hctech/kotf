package client

import (
	"fmt"
	"testing"
)

func TestKotfIint(t *testing.T) {
	//test := terraform.NewTerraform()
	//var provider map[string]interface{}
	//var cloudRegion map[string]interface{}
	//var hosts map[string]interface{}
	//
	//vars := map[string]interface{}{
	//	"provider":    provider,
	//	"cloudRegion": cloudRegion,
	//	"hosts":       hosts,
	//}
	//result, err := test.Init("", "", vars)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(result)
	client := NewKotfClient("localhost", 8083)
	result, err := client.Init("", "", "", "", "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
