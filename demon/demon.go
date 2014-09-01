package demon

import (
	"log"

)

type Demon struct {

}


//start this instance as a new process - wtf?
func (*Demon) Start() {}

//stop this instance - if gracefull = true, node finish pending work befor stopping
func (*Demon) Stop(gracefull bool) {}

//reload configuration - if gracefull = true, new configuration apply only to new work
func (*Demon) Reload(gracefull bool) {}

//restart this instance as a new process - if gracefull = true, node finish pending work befor stopping
func (*Demon) Restart(gracefull bool) {}

//register a ressource to be inherited if demon restarts
func (*Demon) RegisterRessource() {}

//unregister a ressource to be inherited if demon restarts
func (*Demon) UnregisterRessource() {}
