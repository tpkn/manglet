package manglet

import (
	"bufio"
	"os"
	"strings"
	
	"manglet/models"
	"manglet/skip"
	"manglet/utils"
)

func Run(options *models.CLI) (string, error) {
	// Make custom output file if needed
	options.Output = utils.MakeOutputPath(options.Input, options.Output)
	
	// Input reader
	rs, err := os.Open(options.Input)
	if err != nil {
		return "", err
	}
	defer rs.Close()
	
	// Output writer
	out, err := os.OpenFile(options.Output, os.O_TRUNC|os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ws := bufio.NewWriterSize(out, 4096 * 2)
	if err != nil {
		return "", err
	}
	defer out.Close()
	
	// Mangle
	var row = 0
	var skip_rows = skip.Rows(options.Skip)
	var skipping_enabled = len(skip_rows) > 0
	var mangle_groups = strings.Split(options.MangleGroups, "")
	
	var scanner = bufio.NewScanner(rs)
	for scanner.Scan() {
		row++
		line := scanner.Text()
		
		if skipping_enabled {
			if _, found := skip_rows[row]; found {
				_, e := ws.WriteString(line + "\n")
				if e != nil {
					return "", e
				}
				continue
			}
		}
		
		line = Scenarios(line, mangle_groups)
		
		_, e := ws.WriteString(line + "\n")
		if e != nil {
			return "", e
		}
	}
	
	err = ws.Flush()
	if err != nil {
		return "", err
	}
	
	return options.Output, nil
}
