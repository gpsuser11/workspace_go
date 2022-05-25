package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var player = [3]int{0, 0, 0}

	for count := 0; count < 100; count++ {
		var card1 [10]int
		var card2 [10]int
		var check bool = true
		var temp_card int
		var temp1 int
		var temp2 int

		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		card1[0] = r1.Intn(10) + 1
		for temp := 0; temp < 9; temp++ {
			check = true
			for check {
				check = false
				s1 := rand.NewSource(time.Now().UnixNano())
				r1 := rand.New(s1)
				temp_card = r1.Intn(10) + 1

				for _, value := range card1 {
					if value == temp_card {
						check = true
					}
				}
			}
			card1[temp+1] = temp_card
		}
		// fmt.Println(card1)

		s2 := rand.NewSource(time.Now().UnixNano())
		r2 := rand.New(s2)

		card2[0] = r2.Intn(10) + 1
		for temp := 0; temp < 9; temp++ {
			check = true
			for check {
				check = false
				s2 := rand.NewSource(time.Now().UnixNano())
				r2 := rand.New(s2)
				temp_card = r2.Intn(10) + 1

				for _, value := range card2 {
					if value == temp_card {
						check = true
					}
				}
			}
			card2[temp+1] = temp_card
		}
		// fmt.Println(card2)

		temp1 = (card1[0] + card2[0]) % 10
		temp2 = (card1[1] + card2[1]) % 10
		// fmt.Println(temp1, temp2)
		if temp1 > temp2 {
			player[0] = player[0] + 1
		} else if temp1 < temp2 {
			player[1] = player[1] + 1
		} else {
			player[2] = player[2] + 1
		}
	}
	fmt.Println("player 1 win ", player[0], " %")
	fmt.Println("player 2 win ", player[1], " %")
	fmt.Println("draw ", player[2], " %")
}
