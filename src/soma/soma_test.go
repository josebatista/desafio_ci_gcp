package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(5, 5)
	if total != 10 {
		t.Errorf("A soma esta incorreta. Recebemos %d, mas esperava o valor %d", total, 10)
	}
}
