package main

import (
    "fmt"
    "time"
    "math/rand"
    "sync"
    "sort"
    "math"
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
	delayCoeff:= estimateDelayCoeff(len(array))
 	for i := 0; i < len(array); i++ {
 		go func(){
 			currNum:=array[i]
 			<-startInsideLoop
 			delayMilliSec(currNum*delayCoeff)
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

 func estimateDelayCoeff(n int)(int){
 	exp:= 4.5503685360*(math.Pow(1.0004632701,float64(n)))
 	return int(math.Ceil(exp))
 }

 func isSorted(original []int, sorted []int)(bool){
 	sort.Ints(original)
 	for i:= 0;i<len(original);i++{
 		if(original[i]!=sorted[i]){
 			return false
 		}
 	}
 	return true
 }
func main(){
	array:= makeRandomArray(3000,100)
	// fmt.Println("Random Array:",array)
	sorted:= threadedSort(array)
	// fmt.Println("Sorted Array:",sorted)
	fmt.Println(isSorted(array, sorted))
}