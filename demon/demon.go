package demon

import (
	"log"
	"os"
	"syscall"
	"os/signal"

)

type workerId interface{} //dont know the type yet

type Configuration struct {

}

//thats a Worker with a gracefull shotdown
type GracefullWorker struct {
	wid workerId
	config *Configuration
	Do func()
}

//thats the Demon - it holds all the state of the Demon, and ALL of it can be inherited by a parent killing child monster
type Demon struct {
	workers map[workerId] *GracefullWorker
	files map[string] *os.File
}

//setup state at start & restart. Also inherit state from parent <. that is the load state
func NewDemon(isInherited bool) (*Demon) {
	return nil
}

//terminate this instance ASAP - ungracefully but clean
func (*Demon) Terminate() {}

//shutdown this instance - node finish pending work befor terminating
func (*Demon) Shutdown() {}

//reload configuration - new configuration apply only to new work
func (*Demon) Reload() {}

//restart this instance as a new process - node finish pending work befor stopping
func (*Demon) Restart() {}

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

//Signals and their meaning
const (
	TERMINATE = syscall.SIGTERM //Terminating ungracefully but clean
	SHUTDOWN = syscall.SIGQUIT //Gracyfully shuting down
	RELOAD = syscall.SIGHUP //Reloading Configuration, Starting new Worker, gracefully Shutdown old worker
	RESTART = syscall.SIGINT //Restarting gracefully //
)

//starts a go routine to listen to signals about start, stop, reload and restart
func (d *Demon) handleSignals() {
    //connect channel with signal
    ch := make(chan os.Signal, 2)
    signal.Notify(ch, TERMINATE,SHUTDOWN,RELOAD,RESTART)

    //listen for signals and handle them
    go func(c chan os.Signal) {
        sig := <-c
        switch sig {
        	case TERMINATE:
        		log.Printf("Caught signal %s: Terminating ungracefully but clean",sig)
        		d.Terminate()
        	case SHUTDOWN:
        		log.Printf("Caught signal %s: Gracyfully shuting down",sig)
        		d.Shutdown()
        	case RELOAD:
        		log.Printf("Caught signal %s: Reloading Configuration, Starting new Worker, gracefully Shutdown old worker",sig)
        		d.Reload()
        	case RESTART:
        		log.Printf("Caught signal %s: Restarting gracefully",sig)
        		d.Restart()
        	default:
        		log.Printf("Can't handle signal %s: just living on",sig)
        }
    }(ch)
}

//kills parent process after restart
func (*Demon) terminateParent() {}

//waits untill all workers are done
func (*Demon) beGracefull() {}

//spawns a child process and inherits the current state
func (*Demon) spawnChild() {}