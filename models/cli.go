package models

type CLI struct {
	About        string // Prints app info
	Input        string // Input file path
	Output       string // Output file path. Could be a file or just a directory
	MangleGroups string // Mangle options: 'l' = letters; 'd' = digits; 's' = symbols
	Skip         string // Rows to skip, separated by commas
	Verbose      bool   // Prints output file path to STDOUT
}
