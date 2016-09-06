package main

import (
	"testing"
	"log"
)

func TestFlip(t *testing.T) {
	res := flip("┬─┬ノ(º_ºノ)\n")

	if res != "(╯°□°）╯︵ ┻━┻" {
		log.Fatal("table not flipped %s", res)
		t.Fail()
	}
}
