package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/pramineni01/simple_task_scheduler/tasks"
)

func Execute() {
	jsonFile, err := os.Open("input/input.json")
	if err != nil {
		log.Println("Fatal error opening input file: ", err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	byteVal, _ := ioutil.ReadAll(jsonFile)

	var inpTasks tasks.TasksInput
	if err := json.Unmarshal(byteVal, &inpTasks); err != nil {
		log.Println("Fatal error while unmarshaling input: ", err)
		os.Exit(1)
	}

	// print output
	out := tasks.Run(&inpTasks)
	for _, o := range out {
		fmt.Printf("ID: %s\t Start: %s\t Complete: %s\n", o.ID, o.Start.Format(time.RFC1123), o.End.Format(time.RFC1123))
	}
}
