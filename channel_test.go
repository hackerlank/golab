package main

import (
	"testing"
)
const (
	PlayerCount = 1000000
)

type PlayerInfo struct {
	Id 	int64	// player id
	Money 	int64   // 金币
	FightValue int64 // 战力
	Level 	int32    // 等级
}



func InitData() []PlayerInfo {
	data := make([]PlayerInfo, PlayerCount, PlayerCount)
	for i := 0; i < PlayerCount; i++ {
		data = append(data, PlayerInfo{Id:int64(i+1), Money:99999, FightValue:99999, Level:555})
	}
	return data
}

func Consume(data chan PlayerInfo, done chan bool) {
	for i := 0; i < PlayerCount; i++ {
		_ = <- data
	}
	done <- true
}

// 测试chan []PlayerInfo和 chan []*PlayerInfo的性能差异
func BenchmarkChannelSlice(b *testing.B) { 
	data := InitData()
	ch := make(chan PlayerInfo)
	done := make(chan bool)
	go Consume(ch, done)
	b.ResetTimer()
	for i := 0; i < PlayerCount; i++ {
		ch <- data[i]
	}
	_ = <- done
}

func Consume2(data chan *PlayerInfo, done chan bool) {
	for i := 0; i < PlayerCount; i++ {
		_ = <- data
	}
	done <- true
}

func BenchmarkChannelSlicePointer(b *testing.B) { 
	data := InitData()
	ch := make(chan *PlayerInfo)
	done := make(chan bool)
	go Consume2(ch, done)
	b.ResetTimer()
	for i := 0; i < PlayerCount; i++ {
		ch <- &data[i]
	}
	_ = <- done
}