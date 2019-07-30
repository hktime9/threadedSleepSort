package main

import (
    "fmt"
    "time"
    "math/rand"
    "sync"
)

 func delayMilliSec(yourTime int){
	time.Sleep(time.Duration(yourTime) * time.Millisecond)
 }

 func threadedSort(array []int)(sorted []int){
 	start:= make(chan bool)
 	startInsideLoop:= make(chan bool)
	var mutex = &sync.Mutex{}
	index:= 0
	sortedArray:= make([]int,len(array))
 	for i := 0; i < len(array); i++ {
 		go func(){
 			currNum:=array[i]
 			<-startInsideLoop
 			delayMilliSec(currNum*10)
 			mutex.Lock()
 			sortedArray[index]= currNum
 			index++
 			if(index == len(array)){
 				<-start
 			}
 			mutex.Unlock()
 		}()
 		startInsideLoop<-true
 	}
 	start<-true
 	return sortedArray
 }

 func makeRandomArray(number int, thresh int)([]int){
 	returnArray:= make([]int, number)
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<number;i++{
		returnArray[i]= rand.Intn(thresh)
	}
	return returnArray
 }

func main(){
	array:= makeRandomArray(300,10)
	fmt.Println("Random Array:",array)
	sorted:= threadedSort(array)
	fmt.Println("Sorted Array:",sorted)
}