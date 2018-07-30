package main

import (
	"fmt"
	"strconv"
	"time"
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
		go curJob(chans[i], chans[i + 1])
		fmt.Printf("job %d executed\n", i)
	}
	for{
		select {
		case v1 := <-chans[len(chans) - 1]:
			fmt.Printf("log %v from pipeline\n", v1)
			return
		}
	}
	//<-chans[len(chans) - 1]
}

func SingleHash(in, out chan interface{}){
	timer := time.NewTimer(10 * time.Millisecond)
LOOP:
	for{
		select {
		case data := <-in:
			dataInt, ok := data.(int)
			if !ok {
				fmt.Errorf("can't convert to string in SingleHash %v", data)
			}
			dataString := strconv.Itoa(dataInt)
			result := DataSignerCrc32(dataString) + "~" + DataSignerCrc32(DataSignerMd5(dataString))
			fmt.Printf("Singlehash Result %s\n", result)
			out <- result
		case <-timer.C:
			break LOOP
		}
	}
	close(out)
}

func MultiHash(in, out chan interface{}){
	for input := range in{
		inputString, ok := input.(string)
		if !ok {
			fmt.Errorf("Unable to convert to steing in MultiHash")
		}
		for i := 0; i < 6; i++{
			stringOfIteration := strconv.Itoa(i)
			DataSignerCrc32(stringOfIteration + inputString)
			out <- DataSignerCrc32(stringOfIteration + inputString)
		}
	}
	fmt.Printf("===result Multihash executed\n")
	close(out)
}

func CombineResults(in, out chan interface{}){
	result := ""
	for hash := range in{
		result += hash.(string) + "_"
	}
	out <- result
	//close(out)
}