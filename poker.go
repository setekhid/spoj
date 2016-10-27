package main

import (
	"fmt"
	"sort"
)

func main() {

	var test_amount int
	fmt.Scanln(&test_amount)
	for ; test_amount > 0; test_amount-- {

		var a, b, c, d, e string
		fmt.Scanln(&a, &b, &c, &d, &e)

		hands := "None"

		if isFlush(a, b, c, d, e) {
			if isRoyal(a, b, c, d, e) {
				hands = "royal flush"
			} else if isStraight(a, b, c, d, e) {
				hands = "straight flush"
			} else {
				hands = "flush"
			}
		} else if isStraight(a, b, c, d, e) {
			hands = "straight"
		} else {
			p1, p2 := pairs(a, b, c, d, e)
			switch p1 {
			case 4:
				hands = "four of a kind"
			case 3:
				if p2 >= 2 {
					hands = "full house"
				} else {
					hands = "three of a kind"
				}
			case 2:
				if p2 >= 2 {
					hands = "two pairs"
				} else {
					hands = "pair"
				}
			default:
				hands = "high card"
			}
		}

		fmt.Println(hands)
	}
}

func isFlush(a, b, c, d, e string) bool {
	return a[1] == b[1] && b[1] == c[1] && c[1] == d[1] && d[1] == e[1]
}

var straight = map[uint8]int{
	"A"[0]: 0,
	"2"[0]: 1,
	"3"[0]: 2,
	"4"[0]: 3,
	"5"[0]: 4,
	"6"[0]: 5,
	"7"[0]: 6,
	"8"[0]: 7,
	"9"[0]: 8,
	"T"[0]: 9,
	"J"[0]: 10,
	"Q"[0]: 11,
	"K"[0]: 12,
}

func isStraight(a, b, c, d, e string) bool {
	ls := make([]int, 0, 5)
	ls = append(ls, straight[a[0]])
	ls = append(ls, straight[b[0]])
	ls = append(ls, straight[c[0]])
	ls = append(ls, straight[d[0]])
	ls = append(ls, straight[e[0]])
	sort.Ints(ls)
	return ls[1]+1 == ls[2] && ls[2]+1 == ls[3] && ls[3]+1 == ls[4] && (ls[0]+1 == ls[1] || (ls[4]+1)%13 == ls[0])
}

var royal = map[uint8]bool{
	"A"[0]: true,
	"K"[0]: true,
	"Q"[0]: true,
	"J"[0]: true,
	"T"[0]: true,
}

func isRoyal(a, b, c, d, e string) bool {
	return royal[a[0]] && royal[b[0]] && royal[c[0]] && royal[d[0]] && royal[e[0]]
}

func pairs(a, b, c, d, e string) (p1 int, p2 int) {
	counts := make(map[uint8]int)
	counts[a[0]] = counts[a[0]] + 1
	counts[b[0]] = counts[b[0]] + 1
	counts[c[0]] = counts[c[0]] + 1
	counts[d[0]] = counts[d[0]] + 1
	counts[e[0]] = counts[e[0]] + 1
	for _, count := range counts {
		if count > 1 {
			if count > p1 {
				p2 = p1
				p1 = count
			} else {
				p2 = count
			}
		}
	}
	return
}
