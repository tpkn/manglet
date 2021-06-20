package utils

import (
	"path/filepath"
	"regexp"
)

const file_suffix = "_manglet"

// MakeOutputPath makes safe outputh path
func MakeOutputPath(input_path, output_path string) string {
	dir := filepath.Dir(input_path)
	name := filepath.Base(input_path)
	new_path := output_path
	
	if output_path == "" || output_path == input_path {
		new_path = filepath.Join(dir, FilenameSuffix(name, file_suffix))
	} else if IsDir(output_path) {
		new_path = filepath.Join(filepath.Dir(output_path), FilenameSuffix(name, file_suffix))
	}
	
	return new_path
}

// FilenameSuffix adds suffix to the filename
func FilenameSuffix(file_name, suffix string) string {
	ext := filepath.Ext(file_name)
	rule := regexp.MustCompile(regexp.QuoteMeta(ext) + "$")
	return rule.ReplaceAllString(file_name, suffix + ext)
}