package main

import (
	"fmt"
)

type bitBoard uint64
type GameBoard struct {
	boardBlack *bitBoard
	boardWhite *bitBoard
}


func (bb *bitBoard) Conv2Mat() (mat [8][8]int) {
	// bitBoard -> [8][8]intに変換
	var p bitBoard = 1
	for i := 0; i < 64; i++ {
		if (*bb)&p != 0 {
			mat[i/8][i%8] = 1
		} else {
			mat[i/8][i%8] = 0
		}
		p <<= 1
	}
	return
}


func (bb *bitBoard) Print() {
	// bitBoradを2次元配列として表示
	mat := bb.Conv2Mat()
	for _, l := range mat {
		fmt.Println(l)
	}
}


func (bb *bitBoard) Count() int {
	// bitBoardの駒数(立っているbitの数)を数える
	count := uint64(*bb)
	count -= (count >> 1) & 0x5555555555555555
	count = count&0x3333333333333333 + (count>>2)&0x3333333333333333
	count = (count + (count >> 4)) & 0x0f0f0f0f0f0f0f0f
	count = (count + (count >> 8)) & 0x00ff00ff00ff00ff
	count = (count + (count >> 16)) & 0x0000ffff0000ffff
	count = (count + (count >> 32)) & 0x00000000ffffffff

	return int(count)
}

func NewGameBoard() *GameBoard {
	// GameBoardのコンストラクタ
	var bb bitBoard = 0x0000001008000000
	var bw bitBoard = 0x0000000810000000
	return &GameBoard{boardBlack: &bb, boardWhite: &bw}
}

func (gb *GameBoard) getChangePlaces(bb bitBoard, color int) bitBoard{
	if
} 

func main() {
	board := NewGameBoard()
	fmt.Println(board.boardWhite.Count())
}
