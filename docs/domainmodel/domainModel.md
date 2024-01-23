```uml
@startuml Test

package ユーザー集約 {
    object ユーザー {
        ユーザーID
        メールアドレス
        タイプ
    }
}

note left of ユーザー
    * 更新時入力されたメールアドレスが更新前のメールアドレスと同じ場合、更新できない
end note

note bottom of ユーザー
    * タイプは2種類がある。(従業員と顧客)
    * メールアドレスのドメインが会社名の時に従業員になる
end note

package 会社集約 {
    object 会社 {
        会社ID
        名前
        従業員数
    }
}
ユーザー "n" -left-* "1" 会社
note bottom of 会社
    * ユーザーが作成された際にメールアドレスが会社のドメイン名と同じ場合、会社の従業員数を+1にする
end note

note bottom of 会社
    * 更新時ユーザーのタイプが変更された時に会社の従業員数も変動する
    * 顧客から従業員の場合は+1
    * 従業員から顧客の場合は-1
end note

@enduml
```