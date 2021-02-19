# opstats-go

## 感想

opstats の Go 実装版。  
いろり CLI パーサーを試したが github.com/alexflint/go-arg がシンプルなIFで簡単に実装でき使い勝手がよかった。  
またビルドもgolang側で一発ででき、速度も早くNodeJSと比べて楽。  
型の恩恵も受けやすく、Rustほどハードでないので開発体験がよかった。  

## 開発

```
go run ./main.go --help
go run ./main.go ./mock/mini
```
