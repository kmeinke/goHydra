package demon

type workerId interface{} //dont know the type yet

//thats a Worker with a gracefull shotdown
type GracefullWorker struct {
	wid workerId
	config *Configuration
	Do func()
}