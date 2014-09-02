package demon

import (
	"log"
	"os"

)

type workerId interface{} //dont know the type yet

//thats a Worker with a gracefull shotdown
type GracefullWorker struct {
	wid workerId
	Do func()
}

//thats the Demon - it holds all the state of the Demon, and ALL of it can be inherited by a parent killing child monster
type Demon struct {
}

//setup state at start & restart. Also inherit state from parent <. that is the load state
func NewDemon(isInherited bool) (*Demon) {
	return nil
}

//start this instance as a new process - wtf?
func (*Demon) Start() {}

//stop this instance - if gracefull = true, node finish pending work befor stopping
func (*Demon) Stop(beGracefull bool) {}

//reload configuration - if gracefull = true, new configuration apply only to new work
func (*Demon) Reload(beGracefull bool) {}

//restart this instance as a new process - if gracefull = true, node finish pending work befor stopping
func (*Demon) Restart(beGracefull bool) {}

/**/

//register a file to be inherited if demon restarts
func (*Demon) AddFile(f *os.File) {
}

//unregister a file to be inherited if demon restarts
func (*Demon) RemFile() (*os.File) {
	return nil
}

//registers a gracefull Worker in the demon
func (*Demon) AddWorker(w *GracefullWorker) workerId { 
	return nil
}

//unregister a gracefull Worker in the demon
func (*Demon) RemWorker(wid workerId) (*GracefullWorker) {
	return nil
}

/**/

//starts a go routine to listen to signals about start, stop, reload and restart
func (*Demon) handleSignals() {}

//kills parent process after restart
func (*Demon) terminateParent() {}

//waits untill all workers are done
func (*Demon) beGracefull() {}

//spawns a child process and inherits the current state
func (*Demon) spawnChild() {}