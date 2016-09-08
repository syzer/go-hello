package main

import (
	"log"
	"testing"
)

func TestFlipToAngry(t *testing.T) {
	res := flip("┬─┬ノ(º_ºノ)")

	// TODO "(╯°□°）╯︵ ┻━┻"
	if res != "┻━┻(╯°□°）╯" {
		log.Fatal("table not flipped %s", res)
	}
}

func TestFlipToCalm(t *testing.T) {
	res := flip("┻━┻(╯°□°）╯")

	if res != "┬─┬ノ(º_ºノ)" {
		log.Fatal("table not flipped %s", res)
	}
}
