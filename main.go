package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// 設定される変数のポインタを取得
// var msg = flag.String("msg", "デフォルト値", "説明")

var n = flag.Bool("n", false, "bool flag")

func main() {
	//fmt.Println(os.Args[1])
	// ここで実際に設定される
	flag.Parse()
	//fmt.Println(strings.Repeat(_, n))
	i := 0

	// 読み込み用にファイルを開く

	for _, fn := range flag.Args() {

		sf, err := os.Open(fn)
		if err != nil {
			fmt.Fprintln(os.Stderr, "エラー")
			os.Exit(1)
		}
		// 関数終了時に閉じる
		defer sf.Close()

		scanner := bufio.NewScanner(sf)

		for scanner.Scan() {

			//1行分を出力する
			if *n {
				i++
				fmt.Printf("%d: ", i)
			}
			fmt.Println(scanner.Text())
		}
	}
}
