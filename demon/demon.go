package demon

import (
	"log"
	"os"
	_ "os/exec"
	"os/signal"
	"syscall"
	"net"

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

//reload state - if gracefull = true, new configuration apply only to new work
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

//listen to os and demon controller signals and follow their commands 
func (d *Demon) handleSignals() {

//listen to osSignals
	//connect channel with signal
	osch := make(chan os.Signal, 2)
	signal.Notify(osch, os.Interrupt, os.Kill, syscall.SIGTERM)

	//listen for signals and handle them
	go func(c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: terminate clean, but without grace.", sig)
		d.Stop(false)
		os.Exit(0)
	}(osch)

//listen to demon controler unix socket
	//start listener
    l, err := net.Listen("unix", "/tmp/demonctrl.sock")
    if err != nil {
        log.Printf("Could not listen on controler socket: %s", err)
        os.Exit(0)
    }

    //listen on socket
    go func(l net.Listener) {
    	defer l.Close()

	   	for {
	        c, err := l.Accept()
	        if err != nil {
	            log.Printf("Could not accept on controller socket: %s", err)
	            os.Exit(0)
	        }

	        //handle message
	        go func(c net.Conn) {
	        	defer c.Close()

	        	var JO = []byte("JO")
	        	var NO = []byte("NO")
      	
			    for {
			        buf := make([]byte, 512)
			        nr, readerr := c.Read(buf)
			        if readerr != nil {
			            log.Printf("Could not read from controller socket: %s", readerr)
			            os.Exit(0)
			        }

			        command :=  string(buf[0:nr])
			        log.Printf("Controller Command received:",command)

			        switch command {
			        	case "stop":
							c.Write(JO)
			        		d.Stop(true)
							return 
			        	case "restart":
			        		c.Write(JO)
			        		d.Restart(true)
			        		return
			        	case "reload":
			        		c.Write(JO)
			        		d.Reload(true)
			        		return
			        	default:
			        		c.Write(NO)
			        		log.Printf("Invalid Controller Command received, just living on: %s", command)
			        }
			    }
			    
	       	}(c)
	    }
   	}(l) 

}

//kills parent process after restart - or dosnet.Conn the parent exit if all workers are done?
func (*Demon) terminateParent() {}

//start go routine to wait until all workers are done
func (*Demon) beGracefull() {}

//spawns a child process and inherits the current state
func (*Demon) spawnChild() {}