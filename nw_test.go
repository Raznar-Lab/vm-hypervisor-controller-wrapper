package main

import (
	"testing"
)

type NWTest struct {
	BaseTest
}

func (b NWTest) Start() (err error) {
	nwService := b.wrapper.NW()

	b.t.Log("resetting the network")
	success, err := nwService.Reset()
	if !success || err != nil {
		b.t.Log("failed to reset the network")
		return
	}
	b.t.Log("successfully reset the network")

	b.t.Log("creating a rule")
	success, err = nwService.Create("139.99.4.25", "02:00:00:be:c4:6c")
	if !success || err != nil {
		b.t.Log("failed to create the network rule")
		return
	}

	b.t.Log("deleting a rule")
	success, err = nwService.Delete("139.99.4.25", "02:00:00:be:c4:6c")
	if !success || err != nil {
		b.t.Log("failed to delete the network rule")
		return
	}
	b.t.Log("deleted a rule")

	b.t.Log("creating a multiple rule")
	success, err = nwService.CreateMultiple([]string{"139.99.4.25", "139.99.4.21", "139.99.4.23"}, "02:00:00:be:c4:6c")
	if !success || err != nil {
		b.t.Log("failed to create multiple rule")
		return
	}
	b.t.Log("successfully created a multiple rule")

	b.t.Log("resetting the network")
	success, err = nwService.Reset()
	if !success || err != nil {
		b.t.Log("failed to reset the network")
		return
	}
	b.t.Log("successfully reset the network")
	return
}

func TestNW(t *testing.T) {
	wrapper := New("http://15.235.160.20:8080", "cikoHNJOA2", "Caiktjhoi@tha@ithai@tha@izz")
	nwTest := NWTest{}
	nwTest.wrapper = wrapper
	nwTest.t = t

	err := nwTest.Start()
	if err != nil {
		panic(err)
	}
}
