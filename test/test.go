package main

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)


 const test = `/home/dev/temp/bcmc0vp7qcnfj68okjgg.js:1
(function (exports, require, module, __filename, __dirname) { consolxe.log("Chao cac ban")
                                                              ^

ReferenceError: consolxe is not defined
    at Object.<anonymous> (/home/dev/temp/bcmc0vp7qcnfj68okjgg.js:1:63)
    at Module._compile (internal/modules/cjs/loader.js:702:30)
    at Object.Module._extensions..js (internal/modules/cjs/loader.js:713:10)
    at Module.load (internal/modules/cjs/loader.js:612:32)
    at tryModuleLoad (internal/modules/cjs/loader.js:551:12)
    at Function.Module._load (internal/modules/cjs/loader.js:543:3)
    at Function.Module.runMain (internal/modules/cjs/loader.js:744:10)
    at startup (internal/bootstrap/node.js:238:19)
    at bootstrapNodeJSCore (internal/bootstrap/node.js:572:3)`
func main() {
	fmt.Println(TestRegexp(test))
}

func CatchTimeOut() error {
	err := make(chan error, 1)
	fmt.Println(time.Now().Second())
	ticker := time.NewTicker(2 * time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("Over 2s ")
			fmt.Println(time.Now().Second())
			err <- errors.New(t.String())
		}
	}()
	go func() {
		for {
			fmt.Printf("\r %v", "asdasdas")
		}
	}()

	time.Sleep(3 * time.Second)
	defer ticker.Stop()
	fmt.Println("Ticker stopped")
	return <-err
}

func TestRegexp(input string) string {
	rex := regexp.MustCompile(".js:([[:graph:]]|[[:space:]])*(ReferenceError).{1,}")
	return rex.FindString(input)
}
