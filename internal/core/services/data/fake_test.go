package data

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func Test_Pass(t *testing.T) {
	log.Println("test pass")
}

func Test_Fail(t *testing.T) {
	log.Println("test fail")
	t.Fail()
}

func Test_Random(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 100
	rand := rand.Intn(max-min+1) + min

	log.Println(rand)

	if rand > 50 {
		t.Fail()
	}
}
