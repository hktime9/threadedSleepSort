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
	var mutex = &sync.Mutex{}
	index:= 0
	sortedArray:= make([]int,len(array))
 	for i := 0; i < len(array); i++ {
 		go func(){
 			fmt.Println("in")
 			<-start
 			delayMilliSec(array[i])
 			mutex.Lock()
 			sortedArray[index]= array[i]
 			index++
 			mutex.Unlock()
 		}()
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
	array:= makeRandomArray(2,100)
	fmt.Println(array)
	sorted:= threadedSort(array)
	fmt.Println(sorted)
}