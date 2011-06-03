package hello_test

import (
    . "launchpad.net/gocheck"
    "testing"
    "os"
    log "log4go.googlecode.com/svn/stable"
)

// Hook up gocheck into the gotest runner.
func Test(t *testing.T) { TestingT(t) }


type S struct{}
var _ = Suite(&S{})

func (s *S) TestHelloWorld(c *C) {
    c.Check(42, Equals, "42")
    c.Check(os.Errno(13), Matches, "perm.*accepted")
}

func (s *S) TestLog4goModule(c *C) {
    // Setup a log file 
    log.AddFilter("file", log.DEBUG, log.NewFileLogWriter("test.log", false))
  
    // Formatted logging can be done at any of the logging levels (Finest, Fine, Debug, Trace, Info, Warning, Error, Critical)
    log.Trace("Received message: %s (%d)", "Testing", 10)

// Warnings, Errors, and Criticals provide an os.Error that you can use for a return
    //return l4g.Error("Unable to open file: %s", err)
    
// The wrapper functions can also behave like Sprint if the first argument isn't a string
    portno := 1000
    clientid := "123456"
    client := "qbui"
    log.Debug(portno, clientid, client)

    // Use a closure so that if DEBUG isn't logged, it doesn't take any time
    // log.Debug(func()string{ decodeRaw(raw) })
}
