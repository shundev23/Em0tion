package api

import (
	"Em0tion/pkg/sentiment"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	geminiAPIEndpoint = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"
)

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type GenerateContentRequest struct {
	Contents []Content `json:"contents"`
}

type GenerateContentResponse struct {
	Candidates []struct {
		Content Content `json:"content"`
	} `json:"candidates"`
}

func AnalyzeSentiment(inputText string, apiKey string) (sentiment.Result, error) {

	prompt := fmt.Sprintf("このテキストの感情を分析してください: %s\n結果をポジティブ、ネガティブ、中立と0から10で度合いを出力してください。", inputText)

	requestBody := GenerateContentRequest{
		Contents: []Content{
			{
				Parts: []Part{
					{Text: prompt},
				},
			},
		},
	}
	requestBytes, err := json.Marshal(requestBody)
	if err != nil {
		return sentiment.Result{}, err
	}

	req, err := http.NewRequest("POST", geminiAPIEndpoint+"?key="+apiKey, bytes.NewBuffer(requestBytes))
	if err != nil {
		return sentiment.Result{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return sentiment.Result{}, err
	}
	defer resp.Body.Close()

	// デバッグ: レスポンスの詳細
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sentiment.Result{}, err
	}

	var response GenerateContentResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return sentiment.Result{}, err
	}

	if len(response.Candidates) == 0 {
		return sentiment.Result{}, fmt.Errorf("APIからの応答に候補が含まれていません")
	}

	candidate := response.Candidates[0]
	if len(candidate.Content.Parts) == 0 {
		return sentiment.Result{}, fmt.Errorf("APIからの応答にコンテンツが含まれていません")
	}

	rawResponse := ""
	if len(response.Candidates) > 0 && len(response.Candidates[0].Content.Parts) > 0 {
		rawResponse = response.Candidates[0].Content.Parts[0].Text
	}
	fmt.Printf("APIからの生レスポンス:\n%s\n", rawResponse)

	result := sentiment.Result{
		RawResponse: rawResponse,
	}

	// 感情の種類を抽出
	sentimentRegex := regexp.MustCompile(`感情[：:]\s*(\S+)`)
	sentimentMatch := sentimentRegex.FindStringSubmatch(rawResponse)
	fmt.Printf("正規表現マッチの結果: %v¥n", sentimentMatch)
	if len(sentimentMatch) > 1 {
		result.SentimentType = strings.TrimSpace(sentimentMatch[1])
		// 度合いを抽出 (パターン1: "**度合い:** 9 (10が最高)")
		degreeRegex := regexp.MustCompile(`度合い[：:]\s*(\d+)/10`)
		degreeMatch := degreeRegex.FindStringSubmatch(rawResponse)
		if len(degreeMatch) > 1 {
			score, _ := strconv.Atoi(degreeMatch[1])
			result.Score = score
		}
	} else {
		// 度合いを抽出 (パターン2: "**ポジティブ:** (\d+)")
		positiveRegex := regexp.MustCompile(`\*\s*\*\*ポジティブ:\*\*\s*(\d+)`)
		positiveMatch := positiveRegex.FindStringSubmatch(rawResponse)
		negativeRegex := regexp.MustCompile(`\*\s*\*\*ネガティブ:\*\*\s*(\d+)`)
		negativeMatch := negativeRegex.FindStringSubmatch(rawResponse)
		neutralRegex := regexp.MustCompile(`\*\s*\*\*中立:\*\*\s*(\d+)`)
		neutralMatch := neutralRegex.FindStringSubmatch(rawResponse)

		if len(positiveMatch) > 1 && len(negativeMatch) > 1 && len(neutralMatch) > 1 {
			positive, _ := strconv.Atoi(positiveMatch[1])
			negative, _ := strconv.Atoi(negativeMatch[1])
			neutral, _ := strconv.Atoi(neutralMatch[1])

			if positive > negative && positive > neutral {
				result.SentimentType = "ポジティブ"
				result.Score = positive
			} else if negative > positive && negative > neutral {
				result.SentimentType = "ネガティブ"
				result.Score = negative
			} else {
				result.SentimentType = "中立"
				result.Score = neutral
			}
		}
	}

	return result, nil
}
