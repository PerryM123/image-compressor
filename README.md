# image-compressor-microservice

## なぜ作ってるか

レシートを見てOpenAIのOCR（文字認識）でレシートの文字を取得し計算してくれるSplit My Receipt Upというアプリ（[参考](https://github.com/PerryM123/memories_backend)）は制作してますがユーザがレシート写真をアップロードしていただくと写真のサイズは大きすぎて圧縮されてない写真でs3のスペースが取られてしまうのがもったいないなーと思いますのでユーザが写真をアップロードする前に写真のサイズを圧縮してくれるマイクロサービスをGo言語で作りたいです！

## 何を学びたいか

僕は主にフロントエンドの開発をやってますがフロントエンド以外の開発も面白くてGo言語で楽しく実装したいです！

## ローカル環境構築

> [!WARNING]
> Docker Desktop は導入必須です。導入後Docker Desktopを開いてください

```sh
$ cd ~/workspace
$ git clone git@github.com:PerryM123/image-compressor.git
$ cd image-compressor
# コンテナのビルドです. もし既にビルドを実行されたら $ make up-with-build で再ビルドできます
$ make up
# dockerでローカルを起動
$ make air-docker
# ホストパソコン内でローカルを起動したい場合 ( airの導入方法: https://github.com/air-verse/air?tab=readme-ov-file#installation )
$ make air-local
```

## 簡単なBFFアーキテクチャ設計

![alt text](/docs/images/open-ai-project-ver3.jpg)

## 関連リポジトリ

### API仕様書リポジトリ (Swagger)
- [PerryM123/split-my-receipt-up-swagger-doc](https://github.com/PerryM123/split-my-receipt-up-swagger-doc)

### バックエンド側
- [PerryM123/memories_backend](https://github.com/PerryM123/memories_backend)

### MOCK環境
OpenAI APIを利用するとトークンがかかるので動作確認用のモック環境を用意しました。
- [PerryM123/OpenAI API Mock Environment (Split My Receipt Up)](https://github.com/PerryM123/open-ai-api-mock-environment)
