package work

import (
	"fmt"
	"io"
	"net/http"
)

type HttpSink struct {
	Url           string
	Method        string
	AuthRefresher AuthRefresher
	//ClientPool    ClientPool
	Client    *http.Client
	BodyMaker func([]interface{}) io.Reader
}

func (hs *HttpSink) Do(p []interface{}) error {
	req, err := http.NewRequest(hs.Method, hs.Url, hs.BodyMaker(p))
	if err != nil {
		return err
	}
	//client := hs.ClientPool.GetClient()
	res, err := hs.Client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			fmt.Printf("Error while closing %v", err)
		}
	}()
	if res.StatusCode != 200 {
		fmt.Println(res)
	}
	return nil
}
