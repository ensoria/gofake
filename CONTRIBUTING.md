# Contributing

While the faker project is relatively simple in terms of implementation, it has unique management challenges such as multi-language support and copyright concerns.
For this reason, we are not accepting contributors until the management process is established.

However, bug reports and improvement suggestions are very welcome. For now, we would appreciate your contributions in that form.

## About Coding Agent Instructions

When using coding agents, please use a symbolic link to `ai-instructions.md` as the common instruction file for each agent.

```sh
# Claude
ln -s ./ai-instructions.md ./CLAUDE.md

# Codex
ln -s ./ai-instructions.md ./AGENTS.md

# GitHub Copilot
mkdir .github
ln -s ../ai-instructions.md ./.github/copilot-instructions.md
```

## Adding a locale

1. Test the provider through `faker.CreateWithLocaleAndSeed`. Check local address, postcode and domestic phone formats, and record the source of any reserved numbering ranges.
2. Update the [encli seed command](https://github.com/ensoria/encli) alongside the provider. Add the locale to the help in `cmd/seed.go`, the accepted values in `internal/common/validation/validation.go`, and the runner imports and constructor selection in `internal/seed/generator/generator.go`.
3. Update the consuming project's gofake dependency and use the matching encli binary. Run `encli seed --locale <code>` against a disposable database and inspect the inserted address, postcode and phone number. Also test the default locale and rejection of an unknown code.

> **日本語の説明は英語の説明の後に記載されています。** [→ 日本語版へジャンプ](#contributing日本語)
> *Japanese documentation follows the English documentation below.*

---

# Contributing（日本語）

このfakerプロジェクトは、基本的には実装は簡単ですが、他言語対応であることや、著作権の問題など、独特の「管理の難しさ」があります。
そのため、管理方法が確立されるまで、当面はcontributorsの参加は受け付けていません。

ただし、不具合の報告や、改善案などは大歓迎です。当面は、そちらで貢献していただけましたら幸いです。

## About Coding Agent Instructions

コーディングエージェントを利用する場合は、`ai-instructions.md`のシンボリックリンクを各エージェントの共通の指示として使ってください。

```sh
# Claude
ln -s ./ai-instructions.md ./CLAUDE.md

# Codex
ln -s ./ai-instructions.md ./AGENTS.md

# GitHub Copilot
mkdir .github
ln -s ../ai-instructions.md ./.github/copilot-instructions.md
```

## ロケールを追加するとき

1. `faker.CreateWithLocaleAndSeed` から住所・郵便番号・国内電話番号の書式を検証します。フィクション用の予約番号帯を採用した場合は出典を記載します。
2. [encli](https://github.com/ensoria/encli) の `cmd/seed.go` でヘルプを、`internal/common/validation/validation.go` で許容値を、`internal/seed/generator/generator.go` で生成ランナーの import とコンストラクタ選択を更新します。
3. 利用側プロジェクトの gofake 依存を更新し、対応版の encli バイナリで `encli seed --locale <code>` を検証用 DB に対して実行します。投入された住所・郵便番号・電話番号を確認し、ロケール省略時の既定値と不明なコードの拒否もテストします。
