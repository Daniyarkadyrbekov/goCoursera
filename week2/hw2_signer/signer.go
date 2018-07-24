package main

import (
	"fmt"
	"strconv"
)

// сюда писать код
func ExecutePipeline(jobs... job){
	fmt.Print("jobs len ", len(jobs), "\n")

	var chans []chan interface{}
	chans = make([]chan interface{}, len(jobs) + 1)
	for i := range  chans{
		chans[i] = make(chan interface{})
	}
	for i, curJob := range jobs{
		if i != len(jobs) - 1{
			go curJob(chans[i], chans[i + 1])
		}else{
			go curJob(chans[i], chans[i + 1])
		}
	}
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
	inputString, ok := input.(string)
	if !ok {
		fmt.Errorf("Unable to convert to steing in MultiHash")
	}
	result := ""
	for i := 0; i < 6; i++{
		stringOfIteration := strconv.Itoa(i)
		result = result + DataSignerCrc32(stringOfIteration + inputString)
	}
	//fmt.Printf("result ")
	out <- result
}

func CombineResults(in, out chan interface{}){
	result := ""
	for {
		select{
		case <-in:
			resultChan := <-in
			result = resultChan.(string) + "_"
		}
	}
	out <- result
	close(out)
}