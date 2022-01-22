package main

import (
	"io"
	"net/http"
)

func infoHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `
<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="UTF-8">
<title>インフォメーション</title>
</head>
<body>
<h1>ようこそ</h1>
</body>
</html>
    `)
}

func main() {
	http.HandleFunc("/info", infoHandler)
	http.ListenAndServe(":8080", nil)
}
