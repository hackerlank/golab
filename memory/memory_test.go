package main

import (
	"testing"
)
const (
	PlayerCount = 1000000
)


type PlayerInfo struct {
	A1 	int64
	A2 	int64 
	A3  int64  
	A4 	int64
	A5	int64
	A6	int64
	A7	int64
	A8	int64
	A9	int64
	A10	int64
}

func TestMemoryPointer(t *testing.T) {
	// 分配map，value存放obj
	// http://blog.golang.org/go-maps-in-action
	// 180M
	mapdata := make(map[int64]PlayerInfo, PlayerCount)
	// 下面的赋值消耗内存3M
	for i := 0; i < PlayerCount; i++ {
		mapdata[int64(i)] = PlayerInfo{}
	}
	// 分配slice，存放obj
	// http://blog.golang.org/go-slices-usage-and-internals
	// 76M
	slicedata := make([]PlayerInfo, PlayerCount)
	for i := 0; i < PlayerCount; i++ {
		slicedata[i] = mapdata[int64(i)]
	}
	// 分配map, value存放指针
	// 36M
	mapdata2 := make(map[int64]*PlayerInfo, PlayerCount)
	// 赋值消耗内存64.8M
	for i := 0; i < PlayerCount; i++ {
		mapdata2[int64(i)] = &PlayerInfo{}
	}
	// 分配slice, 存放指针
	slicedata2 := make([]*PlayerInfo, PlayerCount)
	for i := 0; i < PlayerCount; i++ {
		slicedata2[i] = mapdata2[int64(i)]
	}
	// 分配slice, 存放指针，10倍
	slicedata3 := make([]*PlayerInfo, PlayerCount*10)
	for i := 0; i < PlayerCount*10; i++ {
		slicedata3[i] = mapdata2[int64(i/10)]
	}
}