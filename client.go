package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func Notify(ip string, port int, n *Notification) error {
	b, err := json.Marshal(n)
	if err != nil {
		return errors.Wrapf(err, "Cannot serialize Notification: %s", err)
	}
	log.Printf("notify to %s: %s\n", ip, string(b))
	res, err := http.DefaultClient.Post(fmt.Sprintf("http://%s:%d/notifications", ip, port), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return errors.Wrapf(err, "Cannot post Notification: %s", err)
	}

	if res.StatusCode != http.StatusCreated {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return errors.Wrapf(err, "Cannot create Notification: %s", string(body))
		}
		return errors.Errorf("Cannot create Notification: %d %s", res.StatusCode, res.Status)
	}
	return nil
}
