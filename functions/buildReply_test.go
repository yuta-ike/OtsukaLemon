package main

import (
	"testing"
)

func TestBuildReply(t *testing.T) {
	if result := buildReply("おつかれ"); result != "おつかれもん" {
		t.Error("「おつかれ」が「おつかれもん」になっていません。現在の出力：" + result)
	}
	if result := buildReply("ただいま"); result != "ただいまんもす" {
		t.Error("「ただいま」が「ただいまんもす」になっていません。現在の出力：" + result)
	}
	if result := buildReply("ひいらぎ"); result != "ひいらぎんなん" {
		t.Error("「ひいらぎ」が「ひいらぎんなん」になっていません。現在の出力：" + result)
	}
	if result := buildReply("しゃみせん"); result != "しゃみせん" {
		t.Error("「しゃみせん」が「しゃみせん」になっていません。現在の出力：" + result)
	}
	if result := buildReply("どれみふぁ"); result != "どれみふぁ" {
		t.Error("「どれみふぁ」が「どれみふぁ」になっていません。現在の出力：" + result)
	}
	if result := buildReply("オツカレ"); result != "オツカレ" {
		t.Error("「オツカレ」が「オツカレ」になっていません。現在の出力：" + result)
	}
	if result := buildReply("オツカレ！"); result != "オツカレ！" {
		t.Error("「オツカレ！」が「オツカレ！」になっていません。現在の出力：" + result)
	}
	if result := buildReply("京都大学"); result != "京都大学" {
		t.Error("「京都大学」が「京都大学」になっていません。現在の出力：" + result)
	}
	if result := buildReply("Osaka"); result != "Osaka" {
		t.Error("「Osaka」が「Osaka」になっていません。現在の出力：" + result)
	}
}
