package main

import (
	"flag"
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

	// コマンドラインオプションの定義
	textPtr := flag.String("text", "", "分析するテキスト")
	flag.Parse()

	// 分析したいテキスト
	var inputText string
	if *textPtr != "" {
		inputText = *textPtr
	} else {
		args := flag.Args()
		if len(args) > 0 {
			inputText = strings.Join(args, " ")
		} else {
			// 引数が指定されなかった場合のデフォルトテキスト
			inputText = "今日の天気はどうですか？"
			fmt.Println("入力テキストが指定されなかったため、デフォルトテキストを使用します:", inputText)
		}
	}
	analyzeAndDisplaySentiment(inputText, apiKey)
}

func analyzeAndDisplaySentiment(inputText string, apiKey string) {
	fmt.Printf("¥n入力テキスト: %s¥n", inputText)

	result, err := api.AnalyzeSentiment(inputText, apiKey)
	if err != nil {
		fmt.Printf("感情分析エラー: %v¥n", err)
		return
	}

	fmt.Println("--- 解析結果 ---")
	if result.SentimentType != "" {
		fmt.Printf("感情: %s\n", result.SentimentType)
		if result.Score != 0 {
			var bar string
			scale := result.Score
			switch result.SentimentType {
			case "ポジティブ":
				bar = strings.Repeat("+", scale)
			case "ネガティブ":
				bar = strings.Repeat("-", scale)
			case "中立":
				bar = strings.Repeat("=", scale)

			}
			fmt.Printf("[%s] (%d)¥n", bar, result.Score)
		} else {
			fmt.Println()
		}
	} else {
		fmt.Println("感情を特定できませんでした。")
	}
}
