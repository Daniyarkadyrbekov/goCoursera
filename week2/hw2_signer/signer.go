package main

import (
	"fmt"
	"strconv"
)

// сюда писать код
func ExecutePipeline(jobs... job){
	fmt.Print("jobs len ", len(jobs), "\n")
}

func SingleHash(in, out chan interface{}){
	data := <- in
	dataString, ok := data.(string)
	if !ok {
		fmt.Errorf("can't convert to string in SingleHash %v", data)
	}
	result := DataSignerCrc32(dataString) + "~" + DataSignerCrc32(DataSignerMd5(dataString))

	out <- result
}

func MultiHash(in, out chan interface{}){
	input := <- in
	for i := 0; i < 6; i++{
		stringOfIteration := strconv.Itoa(i)
		inputString, ok := input.(string)
		if !ok {
			fmt.Errorf("Unable to convert to steing in MultiHash")
		}
		result := DataSignerCrc32(stringOfIteration + inputString)
		out <- result
	}
}

func CombineResults(in, out chan interface{}){

}