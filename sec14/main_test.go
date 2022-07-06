package main

import "testing"

var Debug bool = false

// 名前がTestで始まる関数がテスト用の関数
func TestIsOne(t *testing.T) {
	i := 1
	if Debug {
		t.Skip("スキップさせる")
	}
	v := IsOne(i)

	if !v {
		t.Errorf("%v != %v", i, 1) // エラー
	}
}

// テスト実行は go test
// すべてのパッケージをテストするときはgo test ./...
// -v verbose
// -cover カバー率
