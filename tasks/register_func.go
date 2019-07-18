package tasks

var TaskFuncs = map[string]interface{}{
	"hello":          HelloWorld,
	"get_first_user": GetUserFirst,
}

func GetTasks() (m map[string]interface{}) {
	return TaskFuncs
}
