package day_04

import (
	"strconv"
	"sync"
)

type Solution struct{}

func (s *Solution) PartOne(input string) (string, error) {
	var rolledNums []int

	var i int
	numStart := 0
	for i = 0; input[i] != '\n'; i++ {
		if input[i] != ',' {
			continue
		}
		num, err := strconv.Atoi(input[numStart:i])
		if err != nil {
			return "", err
		}
		rolledNums = append(rolledNums, num)
		numStart = i + 1
	}
	i++
	i++

	var players []BingoPlayer
	var round sync.WaitGroup
	bingo := make(chan int)

	// TODO: Fix this parser so that I don't have to add a newline on the input file
	for i < len(input) {
		player := BingoPlayer{}
		player.Init(&round, bingo)
		player.ReceiveCard(input[i : i+15*5])
		players = append(players, player)
		i += 15*5 + 1
	}

	go func() {
		for _, n := range rolledNums {
			round.Add(len(players))
			for idx := range players {
				go func(player *BingoPlayer) {
					player.FillIn(n)
				}(&players[idx])
			}
			round.Wait()
		}
	}()
	result := <-bingo

	return strconv.Itoa(result), nil
}

func (s *Solution) PartTwo(input string) (string, error) {
	return "", nil
}

type BingoPlayer struct {
	card     BingoCard
	bingoCh  chan int
	filledIn *sync.WaitGroup
}

type BingoCard struct {
	numbers [5][5]int
	marks   [5][5]bool
}

func (b *BingoPlayer) Init(group *sync.WaitGroup, bingoCh chan int) {
	b.filledIn = group
	b.bingoCh = bingoCh
}

func (b *BingoPlayer) ReceiveCard(card string) {
	row, col := 0, 0
	numStart := 0
	for i, c := range card {
		if c != ' ' && c != '\n' {
			continue
		}
		num, err := strconv.Atoi(card[numStart:i])
		numStart = i + 1
		if err != nil {
			continue
		}
		b.card.numbers[row][col] = num
		col++
		col %= 5
		if c == '\n' {
			row++
		}
	}
}

func (b *BingoPlayer) FillIn(num int) {
	for i, row := range b.card.numbers {
		for j, n := range row {
			if n == num {
				b.card.marks[i][j] = true
				if b.CheckForBingo() {
					b.Bingo(num)
				}
				break
			}
		}
	}
	b.filledIn.Done()
}

func (b *BingoPlayer) CheckForBingo() bool {
	for i := 0; i < len(b.card.marks); i++ {
		rowWise := true
		columnWise := true
		for j := 0; j < len(b.card.marks); j++ {
			rowWise = rowWise && b.card.marks[i][j]
			columnWise = columnWise && b.card.marks[j][i]
		}
		if rowWise || columnWise {
			return true
		}
	}
	return false
}

func (b *BingoPlayer) Bingo(num int) {
	sum := 0
	for i, row := range b.card.marks {
		for j, marked := range row {
			if !marked {
				sum += b.card.numbers[i][j]
			}
		}
	}
	b.bingoCh <- num * sum
}
