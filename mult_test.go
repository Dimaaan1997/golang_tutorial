package main

import (
	"fmt"
	"testing"
)

func TestMult1(t *testing.T) {
	/* Кейс когда 2 разных указателя и 3 значения */
	a := 5
	b := 10
	testNums := []*int{&a, &a, &b}
	mult(testNums, 2)

	if a != 10 {
		t.Errorf("expected 10, got %d", a)
	}

	if b != 20 {
		t.Errorf("expected 20, got %d", b)
	}
	fmt.Println("multTest1 is succssful.")
}

func TestMult2(t *testing.T) {
	/* Кейс когда 1  указатель и 4 значения */
	a := 5

	testNums := []*int{&a, &a, &a, &a}
	mult(testNums, 2)

	if a != 10 {
		t.Errorf("expected 10, got %d", a)
	}
	fmt.Println("multTest2 is succssful.")

}

func TestMult3(t *testing.T) {
	/* Кейс когда, передан пустой слайс */

	testNums := []*int{}
	mult(testNums, 2)
	fmt.Println("multTest3 is succssful.")
}
