package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toBase36(n int) string {
	return strconv.FormatInt(int64(n), 36)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	row := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "rowcol")
		var out strings.Builder

		for col, part := range parts[:len(parts)-1] {
			out.WriteString(part)
			out.WriteString(toBase36(row))
			out.WriteString(toBase36(col))
		}

		out.WriteString(parts[len(parts)-1])
		fmt.Println(out.String())
		row++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}
}
