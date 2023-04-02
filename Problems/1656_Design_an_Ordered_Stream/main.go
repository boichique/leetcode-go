package main

import "fmt"

type OrderedStream struct {
	Id     int
	Stream []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		Id:     0,
		Stream: make([]string, n),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.Stream[idKey-1] = value
	res := []string{}
	for this.Id < len(this.Stream) && this.Stream[this.Id] != "" {
		res = append(res, this.Stream[this.Id])
		this.Id++
	}
	return res
}

func main() {
	conv := Constructor(5)
	fmt.Println(conv)
	conv.Insert(3, "ccccc")
	fmt.Println(conv)
	conv.Insert(1, "aaaaa")
	fmt.Println(conv)
	conv.Insert(2, "bbbbb")
	fmt.Println(conv)
	conv.Insert(5, "eeeee")
	fmt.Println(conv)
	conv.Insert(4, "ddddd")
	fmt.Println(conv)
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
