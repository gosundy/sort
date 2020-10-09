package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Heap struct {
	datas []int
}
func main() {
	datas:=rand.Perm(100)
	fmt.Println(datas)
	heap:=NewHeap()
	for _,data:=range datas{
		heap.Push(data)
	}
	fmt.Println(heap.datas)
	for {
		data,err:=heap.Pop()
		if err!=nil{
			break
		}
		fmt.Println(data)
	}

}
func NewHeap()*Heap{
	return &Heap{datas:make([]int,1)}
}
func (heap *Heap)Push(data int)error{
	heap.datas=append(heap.datas,data)
    parentIdx:=(len(heap.datas)-1)/2
    if parentIdx<=1{
    	return nil
	}
	for parentIdx!=0{
		//compare parent,left child, right child
		leftIdx:=parentIdx*2
		rightIdx:=parentIdx*2+1
		tmpIdx:=parentIdx
		if leftIdx<len(heap.datas){
			if heap.datas[leftIdx]<heap.datas[tmpIdx]{
				tmpIdx=leftIdx
			}
		}
		if rightIdx<len(heap.datas){
			if heap.datas[rightIdx]<heap.datas[tmpIdx]{
				tmpIdx=rightIdx
			}
		}
		if tmpIdx==parentIdx{
			break
		}else{
			tmpData:=heap.datas[tmpIdx]
			heap.datas[tmpIdx]=heap.datas[parentIdx]
			heap.datas[parentIdx]=tmpData
			parentIdx=parentIdx/2
		}
	}
	return nil
}
func (heap *Heap)Pop()(int,error){
	if len(heap.datas)<=1{
		return 0,errors.New("empty")
	}
	data:=heap.datas[1]
	heap.datas[1]=heap.datas[len(heap.datas)-1]
	parentIdx:=1
	heap.datas=heap.datas[:len(heap.datas)-1]
	for parentIdx<len(heap.datas){
		tmpIdx:=parentIdx
		leftIdx:=parentIdx*2
		rightIdx:=parentIdx*2+1
		if leftIdx<len(heap.datas){
			if heap.datas[leftIdx]<heap.datas[tmpIdx]{
				tmpIdx=leftIdx
			}
		}
		if rightIdx<len(heap.datas){
			if heap.datas[rightIdx]<heap.datas[tmpIdx]{
				tmpIdx=rightIdx
			}
		}
		if parentIdx==tmpIdx{
			break
		}else{
			tmpData:=heap.datas[tmpIdx]
			heap.datas[tmpIdx]=heap.datas[parentIdx]
			heap.datas[parentIdx]=tmpData
			parentIdx=tmpIdx
		}
	}
	return data,nil
}



