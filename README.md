## ToDoアプリ概要
Goの学習過程で作成したToDoアプリです。<br>
ユーザ管理として、ユーザ登録、ログイン、ログアウトの機能を実装しています。<br>
ToDoアプリとしての機能は、ToDoの作成、修正、削除、完了、完了の取消となっております。<br><br>

## アプリの使い方
1. signupに進み、好きなお名前、メールアドレス、パスワードを入力し、ユーザ登録を行います。<br>
2. loginに進み、signupで入力したメールアドレスとパスワードを入力します。<br>
3. [Create]から新しいToDoを作成できます。<br>
4. [Edit]で入力したToDoの内容を修正できます。<br>
5. [Delete]ボタンを押すと、ToDoの削除ができます。<br>
6. [Completed]ボタンを押すと、ToDoが完了扱いになります。<br>
7. [Incompleted]ボタンを押すと、完了扱いになったToDoが再度活性化します。<br><br>

## インストール方法およびアクセス方法
下記を実行<br>
```bush
$ git clone https://github.com/mamehaku-prod/go_todoapp.git
$ cd go_todoapp
$ go mod tidy
$ go run main.go
<br><br>

実行した状態でローカルホストにアクセス<br>
```code
http://localhost:8080
