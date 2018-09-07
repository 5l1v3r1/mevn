## 概要

gitのCommitMessageをCircleCI内でexportすることができます.

## 使い方

CommitMessageに ` [key]=value ` のような値を設定し
circleci上で `$(./menv)` と実行するとexportされる