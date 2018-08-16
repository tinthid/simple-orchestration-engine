package workflow

type Workflow struct {
	Tasks      []Task
	Subscriber string
	Publisher  string
}

type Task struct {
	Name string
}

func CreateWorkflow() (w *Workflow, err error) {
	w = new(Workflow)
	err = nil
	return
}
