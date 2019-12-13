package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"syscall"

	"github.com/DATA-DOG/godog"
)

var Response string

func theAppIsRunning(arg1 string) error {
	pidstr, err := ioutil.ReadFile("server.PID")
	if err != nil {
		return err
	}
	pid, err := strconv.Atoi(string(pidstr[:len(pidstr)-1]))
	if err != nil {
		return err
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return err
	}
	return nil
}

func aRequestIsMadeTo(arg1 string) error {
	resp, err := http.Get("http://127.0.0.1:5000" + arg1)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	Response = string(body)
	return nil
}

func theAppShouldRespond(arg1 string) error {
	expect := arg1
	actual := Response[:len(Response)-1]
	if expect != actual {
		return fmt.Errorf("Expected: '%s', Got: '%s'", expect, actual)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^the "([^"]*)" app is running$`, theAppIsRunning)
	s.Step(`^a request is made to: "([^"]*)"$`, aRequestIsMadeTo)
	s.Step(`^the app should respond: "([^"]*)"$`, theAppShouldRespond)

	s.BeforeScenario(func(interface{}) {
		Response = "" // clean the state before every scenario
	})
}
