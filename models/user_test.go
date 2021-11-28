package models

import (
	"log"
	"math"
	"math/rand"
	"sync"
	"testing"
)

var userFmt = PaintStringFunc("user")

func Test_UserSerializable(t *testing.T) {
	for _, user := range users {
		log.Println(userFmt(user))
	}
}

// TODO :Mirgrate
func Test_Balance(t *testing.T) {
	totalCost := 9000.0 //which should be enough
	var wg sync.WaitGroup
	wg.Add(int(totalCost))

	appendOrder := func(dealNums int) {
		for i := 0; i < dealNums; i++ {
			meal := meals[rand.Intn(len(meals))]
			users[rand.Intn(len(users))].Request(meal)
			users[rand.Intn(len(users))].Accept(meal)
			wg.Done()
		}
	}

	N := 3
	for i := 0; i < N; i++ {
		go appendOrder(int(totalCost) / N)
	}
	wg.Wait()

	var sum float64

	for _, user := range users {
		sum += user.Balance
		log.Println(userFmt(user))
	}

	if math.Abs(sum) > 1e4 {
		t.Errorf("Bill checks failed: sum is %f", sum)
	}
}
