package main

import (
	"encoding/csv"
	"os"
)

func main(){
	//並び替えたい、正しい順番のヘッダーを読み込む
	header, err := os.Open("header.csv")
	if err != nil {
		panic(err)
	}
	defer header.Close()
	reader := csv.NewReader(header)
	template,err := reader.Read()

	//test.csvを開いて読み込む
	file, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader = csv.NewReader(file)

	//lineには変換元のcsvデータの各行が順にしまわれる
	//sortは変換先のcsvの各列が、変換元のcsvの何列目かを示す
	var line []string
    var sort []int

	//lineに、変換元のcsvの一行目、すなわちヘッダーを読み込む
	line, err = reader.Read()

	//変換元のヘッダーとあるべき並び順とを参照し、sortスライスにしまう
	for i := 0; i < len(line); i++{
		for j := 0; j < len(template); j++{
			if line[i] == template[j] {
				sort = append(sort, j)
				break
			}
		}
	}

	//変換先のcsvの出力ファイル、ない場合はつくる
	file2, err := os.OpenFile("output.csv", os.O_WRONLY|os.O_CREATE, 0600)
	defer file2.Close()

	//書き込み用
	writer := csv.NewWriter(file2)

	//一行目として正しい並び順のヘッダー
	writer.Write(template)

	//変換元のcsvの行が尽きるまで無限ループ
	for {
		//変換元のcsvから一行読んでlineにしまう、行が無ければループを抜ける
		line, err = reader.Read()
		if err != nil {
			break
		}

		//正しい順番に並び替えた後の当該行のデータを入れておく用
		var row []string

		//lineに格納したデータをsortに従い一つずつ取り出していってrowに正しい順番でしまう
		for i := 0; i < len(line); i++{
			row = append(row, line[sort[i]])
		}
		writer.Write(row)
	}

	//貯めといたやつ全部書き込む
	writer.Flush()
}
