# Internet のドメイン名に含まれるアポストロフィ

## 現象と再現

全体テストで `DomainName()` が `o'reilly.za` を返し、既存の `internet_test.go` のドメイン名の検査に失敗した。`faker.CreateWithSeed(42)` で作成した 1 インスタンスで、ほかの値を生成せずに `f.Internet().DomainName()` を繰り返すと、35 回目に `o'conner.fr` を返す。

`DomainWord()` は姓を小文字化してそのまま返す。データにはアポストロフィを含む姓があり、これをドメインラベルに連結すると引用符が残る。この実装は v0.0.8 にもある。氏名をインターネット識別子に使う `Internet.LastName()` は既にアポストロフィを取り除いている。

## 後続作業

ドメインラベルを作る際の正規化を決める。まず `DomainWord()` で `Internet.LastName()` を再利用できるか確認し、固定入力の姓で `DomainWord()`、`DomainName()` とドメイン名を使うメール・URL の影響を検証する。公開 API の出力に関わるため、修正を含めるリリースは未決。Phase 3 の文書と CI 設定の変更から独立して扱う。
