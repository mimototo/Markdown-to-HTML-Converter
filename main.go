package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// コマンドラインの引数を取得
	args := os.Args

	// コマンドラインの引数の数をチェック
	if len(args) < 4 {
		fmt.Println("Usage: go run main.go convert <InputFileName> <OutputFileName>")
		fmt.Println(args)
		return
	}

	// コマンドとファイル名を取得
	command := args[1]
	inputFile := args[2]
	outputFile := args[3]

	// コマンドが"convert"であることを確認
	if command != "convert" {
		fmt.Println("Invalid command. Use 'convert' to convert the input file contents.")
		return
	}

	// ファイルを開く
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Can't open this file, sorry")
		return
	}

	// 出力ファイルを作成
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Cant't create the output file:", err)
		return
	}
	defer func() {
		err := output.Close()
		if err != nil {
			fmt.Println("Something to wrong", err)
			return
		}
	}()

	// マークダウンをHTMLに変換する
	err = convert(input, output)
	if err != nil {
		fmt.Println("Conversion failed", err)
	}
}

func convert(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)

	for scanner.Scan() {
		line := scanner.Text()
		htmlLine := convertLine(line)
		_, err := writer.WriteString(htmlLine + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func convertLine(line string) string {
	switch {
	case strings.HasPrefix(line, "# "):
		return "<h1>" + strings.TrimPrefix(line, "# ") + "</h1>"
	case strings.HasPrefix(line, "## "):
		return "<h2>" + strings.TrimPrefix(line, "## ") + "</h2>"
	case strings.HasPrefix(line, "### "):
		return "<h3>" + strings.TrimPrefix(line, "### ") + "</h3>"
	case strings.HasPrefix(line, "#### "):
		return "<h4>" + strings.TrimPrefix(line, "#### ") + "</h4>"
	case strings.HasPrefix(line, "##### "):
		return "<h5>" + strings.TrimPrefix(line, "##### ") + "</h5>"
	default:
		return "<p>" + line + "<p>"
	}
}
