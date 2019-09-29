package main

import "fmt"
import "time"
import "os"

const (
	MAX_REQUEST_NUM = 10
	CMD_USER_POS    = 1
)

var (
	save chan bool
	quit chan bool
	req  chan *Request
)

type Request struct {
	CmdID int16
	Data  interface{}
}

type UserPos struct {
	X int16
	Y int16
}

func init() {
	req = make(chan *Request, MAX_REQUEST_NUM)
	save = make(chan bool)
	quit = make(chan bool)
}

func saveGame() {
	fmt.Printf("Do Something With Save Game.\n")
	quit <- true
}

func onReq(r *Request) {
	pos := r.Data.(UserPos) // assert
	fmt.Println(r.CmdID, pos)
}

func handler() {
	for {
		select {
		case <-save:
			saveGame()
		case r, ok := <-req:
			if ok {
				onReq(r)
			} else {
				fmt.Println("req chan closed.")
				os.Exit(0)
			}
		}
	}

}

func main() {
	newReq := Request{
		CmdID: CMD_USER_POS,
		Data: UserPos{
			X: 10,
			Y: 20,
		},
	}
	go handler()

	req <- &newReq

	time.Sleep(2000 * time.Millisecond)

	save <- true
	close(req)

	<-quit
}
