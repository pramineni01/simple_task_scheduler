package tasks

import (
	"log"
	"os"
	"sync"
	"time"
)

func Run(inp *TasksInput) []TaskOutput {
	if inp == nil {
		log.Println("Invalid input received. Exiting")
		os.Exit(1)
	}

	chOut := make(chan TaskOutput, 3)
	out := make([]TaskOutput, 0)

	done := func() chan bool {
		done := make(chan bool)
		go func(done chan bool) {
			for o := range chOut {
				out = append(out, o)
			}
			done <- true
		}(done)

		return done
	}()

	for i, st := range *&inp.SequentialTasks {
		log.Println("Task execution: concurrent tasks set ", i)
		var cwg sync.WaitGroup
		for _, t := range st.ConcurrentTasks {
			cwg.Add(1)
			go func(ID string, runFor int64) {
				defer cwg.Done()
				start := time.Now()
				time.Sleep(time.Duration(runFor) * time.Second)
				chOut <- TaskOutput{ID, start, time.Now()}
			}(t.ID, t.Runfor)
		}

		// wait till done
		cwg.Wait()
	}

	close(chOut)

	<-done
	return out

}
