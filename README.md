## Member Center Service
一個基於DDD架構重新開發的會員中心服務。

## 特點
- 使用RSA256加密的JWT
- 使用雙Token設計權限架構
- 基於DDD架構開發
- 搭配Redis進行快取增加效能
- 全面docker化，並能使用docker-compose進行管理

## Demo
[API DOC](https://member-center.jmh-su.com/swagger/index.html)

[Live DEMO](https://member.jmh-su.com)

## 已完成功能
- [x] 註冊會員
- [x] 登入
- [x] 忘記密碼
- [x] 重寄驗證信
- [x] 修改密碼
- [X] 取得Profile
- [X] 修改Profile
- [x] 上傳大頭貼
- [X] 取得大頭貼
- [X] 驗證信箱驗證信
- [x] 重新取得Atomic Token

## 如何啟動服務
### 前置作業
- 首先先取得RSA KEY，然後將public key和private key放到專案目錄底下
    - public key檔名=pubkey.pem
    - private key檔名=key.pem
- 取得Google Cloud Storage的金鑰，你會得到一個json檔的金鑰
    - 請將檔名設置成file-center.json
- 依照自身需求設置config.yaml
    - 請先將config-example.yaml更名成config.yaml
- 依照自身需求設置docker-compose.yaml
    - 請先將docker-compose-example.yaml更名成docker-compose.yaml
### 透過docker compose快速部署
```
docker-compose up -d
```

## 如何重新建置Swagger文件
```
swag init -g cmd/main.go
```
