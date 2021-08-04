# CA Tech Dojo
## 制作物名：「カメガチャ」  

![IMG_0223](https://user-images.githubusercontent.com/66200485/127758779-ed5f97f5-f406-414d-b137-75d913fcff27.JPG)  
↑僕の家に住むかめさん

## ディレクトリ構造
<img width="209" alt="スクリーンショット 0003-08-01 午後1 18 15" src="https://user-images.githubusercontent.com/66200485/127759060-9ac560cd-0026-4abd-894c-7b5297b1cb1a.png">

## MVCモデル
- M : model層  
DBへアクセスしたり、構造体を作成するところ  
- V : view = handler層  
tokenからユーザー情報を読み取ったり、paramからidを取得したりするなど、クライアントのリクエストとレスポンスを行うところ  
- C : controller層  
tokenを作ったり、確率に応じてキャラを引いたりするなどのビジネスロジックを組み立てるところ  
 
## 機能一覧
1. ユーザー関連API  
    1.1. ユーザー情報作成API(POST) http://127.0.0.1:8080/user/create  
    1.2. ユーザー情報取得API(GET) http://127.0.0.1:8080/user/get  
    1.3. ユーザー情報更新API(PUT) http://127.0.0.1:8080/user/update  
1. ガチャ実行API(POST) http://127.0.0.1:8080/gacha/draw  
1. ユーザ所持キャラクター一覧取得API(GET) http://127.0.0.1:8080/character/list  

## カメのレア度
- ☆1から☆7まである
- レアなキャラほどガチャを引く確率が低くなるように設定
- MySQLより、新しいキャラを追加することも可能

## アプリの使用方法
1. このリポジトリをclone
2. 以下を実行
以下を実行
```

```

## 使用技術詳細
**言語**：Golang 1.16.3  
**フレームワーク**：echo v3.3.10  
**開発期間**:4週間  
**開発環境**：MacOS, Docker  
**外部サーバー(OS)**：vultur(ubuntu)  
**DB**：MySQL  


