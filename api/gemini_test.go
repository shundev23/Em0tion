package api_test

import (
	"os"
	"testing"

	"Em0tion/api" // プロジェクト名に合わせて修正
)

func TestAnalyzeSentimentPositive(t *testing.T) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		t.Skip("GEMINI_API_KEY environment variable not set, skipping integration test")
	}
	inputText := "今日はとても良い天気で、気分が最高です！"
	result, err := api.AnalyzeSentiment(inputText, apiKey)
	if err != nil {
		t.Fatalf("AnalyzeSentiment failed: %v", err)
	}
	if result.SentimentType != "ポジティブ" {
		t.Errorf("Expected sentiment 'ポジティブ', got '%s'", result.SentimentType)
	}
	if result.Score == 0 {
		t.Errorf("Expected positive score greater than 0, got %d", result.Score)
	}
}

func TestAnalyzeSentimentNegative(t *testing.T) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		t.Skip("GEMINI_API_KEY environment variable not set, skipping integration test")
	}
	inputText := "仕事で大きなミスをしてしまい、本当に落ち込んでいます。"
	result, err := api.AnalyzeSentiment(inputText, apiKey)
	if err != nil {
		t.Fatalf("AnalyzeSentiment failed: %v", err)
	}
	if result.SentimentType != "ネガティブ" {
		t.Errorf("Expected sentiment 'ネガティブ', got '%s'", result.SentimentType)
	}
	if result.Score == 0 {
		t.Errorf("Expected negative score greater than 0, got %d", result.Score)
	}
}

func TestAnalyzeSentimentNeutral(t *testing.T) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		t.Skip("GEMINI_API_KEY environment variable not set, skipping integration test")
	}
	inputText := "今日のニュースは特に何もなかったな。"
	result, err := api.AnalyzeSentiment(inputText, apiKey)
	if err != nil {
		t.Fatalf("AnalyzeSentiment failed: %v", err)
	}
	if result.SentimentType != "中立" {
		t.Errorf("Expected sentiment '中立', got '%s'", result.SentimentType)
	}
	// 中立のスコアが0であるとは限らないので、ここではスコアの存在のみをチェック
	if result.Score == 0 {
		t.Logf("Neutral sentiment with score 0")
	}
}
