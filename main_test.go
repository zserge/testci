package main

import "testing"


var array []int

func init() {
	array = []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}
}

func TestIt(t *testing.T) {
	t.Log(array[6] + array[7])
}
