package utils

import (
	"fmt"
	"testing"
)

type step struct {
	label     string
	input     string
	output    string
	should_be string
}

func TestMakeOutputPath(t *testing.T) {
	steps := []step{
		{ "Output file", "C:\\file_name.txt", "D:\\file_name.txt", "D:\\file_name.txt" },
		{ "Output is dir", "C:\\file_name.txt", "D:\\", "D:\\file_name_manglet.txt" },
		{ "Output same as input", "C:\\file_name.txt", "C:\\file_name.txt", "C:\\file_name_manglet.txt" },
		{ "Empty output", "C:\\file_name.txt", "", "C:\\file_name_manglet.txt" },
	}
	
	for _, k := range steps {
		result := MakeOutputPath(k.input, k.output)
		if result != k.should_be {
			t.Errorf("%v: %v != %v", k.label, result, k.should_be)
		}
		
		fmt.Println(k.label, result)
	}
}
