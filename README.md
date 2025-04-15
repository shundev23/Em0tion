# Em0tion - Gemini APIを使った感情分析ツール

[![Go](https://img.shields.io/badge/Go-1.x-blue?style=flat-square&logo=go&logoColor=white)](https://go.dev/)
![GitHub Public Repository](https://img.shields.io/github/license/YOUR_GITHUB_USERNAME/em0tion?style=flat-square)
![GitHub Last Commit](https://img.shields.io/github/last-commit/YOUR_GITHUB_USERNAME/em0tion?style=flat-square)

Gemini API (Google AI for Developers) を利用したコマンドライン感情分析ツールです。
入力されたテキストの感情を分析し、ポジティブ、ネガティブ、中立のいずれかと、その度合いを簡単なグラフで表示します。

## 概要

このツールは、ユーザーが入力したテキストの感情を客観的に把握することを目的としています。個人的な文章の振り返りや、SNSの投稿の分析などに活用できます。

## インストール

Goがインストールされている環境が必要です。

1.  このリポジトリをクローンします。
    ```bash
    git clone [https://github.com/shundev23/em0tion.git](https://github.com/shundev23/Em0tion.git)
    cd em0tion
    ```
2.  依存関係をダウンロードします。
    ```bash
    go mod tidy
    ```

## 実行方法

1.  Google AI for Developers でAPIキーを取得し、`.env` ファイルを作成してAPIキーを設定します。
    ```
    GEMINI_API_KEY=YOUR_ACTUAL_API_KEY
    ```
2.  コマンドラインから実行します。
    ```bash
    go run . -text "分析したいテキスト"
    ```
    または、テキストを引数として直接渡すこともできます。
    ```bash
    go run . "分析したいテキスト"
    ```
    引数なしで実行すると、デフォルトのテキストが分析されます。

## オプション

* `-text "分析したいテキスト"`: 分析するテキストを指定します。

## 今後の拡張予定

* 結果の出力形式の拡張 (JSONなど)
* 詳細な感情カテゴリの分析
* インタラクティブモードの実装
* UI/UXの拡張

## ライセンス

Em0tion は MIT License の下でライセンスされています。詳細については [LICENSE](LICENSE) ファイルをご覧ください。

## 開発者

[shundev23]