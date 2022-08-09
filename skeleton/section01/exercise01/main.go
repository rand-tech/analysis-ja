package main

import (
	"errors"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"strconv"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) < 1 {
		return errors.New("source code must be specified")
	}

	fname := args[0]
	fset := token.NewFileSet()
	// TODO: fnameのファイルをパースする

	if err != nil {
		return err
	}

	for _, spec := range /* TODO: *ast.ImportSpecのスライスを取得する */ {
		// TODO: パスの文字列リテラルから文字列を取得

		if err != nil {
			return err
		}

		if path == "unsafe" {
			// TODO: token.Pos型からtoken.Position型の値を取得

			fmt.Fprintf(os.Stderr, "%s: import unsafe\n", pos)
		}
	}

	return nil
}
