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