package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/pramineni01/simple_task_handler/tasks"
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

	tasks.Run(inpTasks)

}
