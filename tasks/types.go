package tasks

type TasksInput struct {
	SequentialTasks []struct {
		ConcurrentTasks []struct {
			ID     string `json:"id"`
			Runfor int64  `json:"runfor"`
		} `json:"concurrent_tasks"`
	} `json:"sequential_tasks"`
}

type TaskOutput struct {
	ID    string `json:"id"`
	Start string `json:"start_time"`
	End   string `json:"completion_time"`
}
