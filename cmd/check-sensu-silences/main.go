// This will get a list of current stashes in Sensu and make sure that any that have
// no expiration date are allowed via the setting of a tag in the message.
//
// LICENSE:
//   Copyright 2016 Yieldbot. <devops@yieldbot.com>
//   Released under the MIT License; see LICENSE
//   for details.

package main

import (
	"crypto/tls"
	//"fmt"
	"github.com/codegangsta/cli"
	"github.com/yieldbot/sensuplugin/sensuutil"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

/*

TODO:

- need to document stuff as needed
- need to add error handling
- need to add debugging
- need to parse json to make check more effective
*/

func main() {
	var host string
	var user string
	var pass string
	var bodyString string
	var bodyBytes []byte

	app := cli.NewApp()
	app.Name = "check-sensu-silences"
	app.Usage = "Confirm that all checks silenced with no expiration are accounted for"
	app.Action = func(c *cli.Context) {

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client := &http.Client{Transport: tr}

		req, _ := http.NewRequest("GET", "https://"+host+":4567/stashes", nil)
		req.SetBasicAuth(user, pass)
		resp, _ := client.Do(req)
		if resp.StatusCode == 200 {
			bodyBytes, _ = ioutil.ReadAll(resp.Body)
			bodyString = string(bodyBytes)
		}
		defer resp.Body.Close()

		if strings.Contains(bodyString, "\"expire\":-1") {

			//fmt.Printf("body: %v\n", bodyString)
			sensuutil.Exit("critical", bodyString)
		} else {
			sensuutil.Exit("ok", "All is well\n")
			//fmt.Printf("All is well\n")
		}
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "user, u",
			Value:       "",
			Usage:       "Username for the sensu api",
			EnvVar:      "API_USER",
			Destination: &user,
		},

		cli.StringFlag{
			Name:        "password, p",
			Value:       "",
			Usage:       "Password for the sensu api",
			EnvVar:      "API_PASS",
			Destination: &pass,
		},
		cli.StringFlag{
			Name:        "host",
			Value:       "localhost",
			Usage:       "Sensu API host",
			Destination: &host,
		},
	}

	app.Run(os.Args)
}
