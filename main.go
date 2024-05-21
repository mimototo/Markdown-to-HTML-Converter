package main

import (
	"fmt"
	"io"
	"os"
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
	defer input.Close()

	// ファイル内容を読み取る
	var data []byte
	buffer := make([]byte, 1024)
	for {
		count, err := input.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to load file", err)
				return
			}
			break
		}
		data = append(data, buffer[:count]...)
	}

	// ファイルの内容を逆順にする
	reversedData := reverseBytes(data)

	// Outputファイルを作成する
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Can't create the output file:", err)
		return
	}
	defer output.Close()

	// 逆順にした内容をoutputファイルに書き込む
	_, err = output.Write(reversedData)
	if err != nil {
		fmt.Println("fail to write to output file:", err)
		return
	}

	fmt.Println("File content reversed and written to", output)
}

func reverseBytes(data []byte) []byte {
	length := len(data)
	reversed := make([]byte, length)

	for i := range reversed {
		reversed[length-i-1] = data[i]
	}

	return reversed
}
