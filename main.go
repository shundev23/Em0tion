package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"Em0tion/api"

	"github.com/joho/godotenv"
)

func main() {
	// .envファイルからAPIKeyを読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading .env file: %v", err)
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		fmt.Println("エラー: GEMINI_API_KEY .envファイルに環境変数が設定されていません。")
		os.Exit(1)
	}

	// コマンドライン引数の取得
	args := os.Args[1:] // os.Args[0]はプログラム自身のパスなので、args[0]から始まる引数を取得する

	// 分析したいテキスト
	var inputText string
	if len(args) > 0 {
		inputText = strings.Join(args, " ")
	} else {
		// 引数が指定されなかった場合のデフォルトテキスト
		inputText = "今日の天気はどうですか？"
		fmt.Println("入力テキストが指定されなかったため、デフォルトテキストを使用します:", inputText)
	}
	analyzeAndDisplaySentiment(inputText, apiKey)
}

func analyzeAndDisplaySentiment(inputText string, apiKey string) {

	result, err := api.AnalyzeSentiment(inputText, apiKey)
	if err != nil {
		fmt.Printf("感情分析エラー: %v¥n", err)
		return
	}

	fmt.Println("--- 解析結果 ---")
	if result.SentimentType != "" {
		fmt.Printf("感情: %s\n", result.SentimentType)
		if result.Score != 0 {
			fmt.Printf("度合い: %d", result.Score)
		}
		fmt.Println()
	} else {
		fmt.Println("感情を特定できませんでした。")
	}

}
