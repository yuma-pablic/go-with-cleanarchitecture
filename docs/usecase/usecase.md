```uml
@startuml
' 1. 名前は書こう

    ' 2. 左から右に矢印 デフォルトは上から下
    left to right direction

    ' 3. 1行コメント
    /' 
        複数行コメント
     '/
    ' 4. 棒人間 or 上半身
    skinparam actorStyle awesome

    ' 5. ユースケースの設定
    skinparam usecase {
        ' 5-1. ユースケースの背景色
        BackgroundColor DarkSeaGreen
        ' 5-2. ユースケースの境界線
        BorderColor DarkSlateGray
    }

    ' 6. アクターの設定
    skinparam actor {
        ' 6-1. アクターの背景色 RGBでも指定できる
        BackgroundColor #333333
        ' 6-2. アクターの境界線 RGBでも指定できる
        BorderColor #999999
    }

    ' 6-3. アクター（ペルソナ）人の設定
    ' 5-3. rectangle は四角で、packageは牛乳パックみたい
    actor ユーザー
    package 試合 {
        usecase ユーザー一覧
        usecase ユーザー登録 
        usecase 会社を更新 
        ' 10. 仕切り文字
        usecase ユーザー更新
    }
    ユーザー --> ユーザー一覧
    ユーザー --> ユーザー登録
    ユーザー --> ユーザー更新
    ' 8-3. 点線
    ユーザー登録 .> 会社を更新
    ユーザー更新 .> 会社を更新
@enduml
```