// Copyright © 2016 Yieldbot <devops@yieldbot.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package sensupluginssensu

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/yieldbot/sensuplugin/sensuutil"
)

var apiHost string
var apiUser string
var apiPassword string
var bodyString string
var bodyBytes []byte

// Tech Debt
// YELLOW needs clean refactor
// YELLOW needs unit tests

// checkSensuSilencesCmd represents the checkSensuSilences command
var checkSensuSilencesCmd = &cobra.Command{
	Use:   "checkSensuSilences",
	Short: "Confirm that all checks silenced with no expiration are accounted for",
	Long: `Checks that have been silenced with no expiration can cause problems and lead to
  unforeseen issues. This check will ensure that any checks that receive a expiration time
  of 'no expiration' are known to everyone and can be accounted for and not forgotten about.`,
	Run: func(sensupluginssensu *cobra.Command, args []string) {

		if apiHost == "" {
			apiHost = viper.GetString("sensupluginssensu.checkSensuSilences.apiHost")
		}
		if apiUser == "" {
			apiUser = viper.GetString("sensupluginssensu.checkSensuSilences.apiUser")
		}
		if apiPassword == "" {
			apiPassword = viper.GetString("sensupluginssensu.checkSensuSilences.apiPassword")
		}

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client := &http.Client{Transport: tr}

		req, _ := http.NewRequest("GET", "https://"+apiHost+":4567/stashes", nil)
		req.SetBasicAuth(apiUser, apiPassword)
		resp, _ := client.Do(req)
		if resp.StatusCode == 200 {
			bodyBytes, _ = ioutil.ReadAll(resp.Body)
			bodyString = string(bodyBytes)
		}
		defer resp.Body.Close()

		if strings.Contains(bodyString, "\"expire\":-1") {

			sensuutil.Exit("critical", bodyString)
		} else {
			sensuutil.Exit("ok")
		}

	},
}

func init() {
	RootCmd.AddCommand(checkSensuSilencesCmd)

	checkSensuSilencesCmd.Flags().StringVar(&apiUser, "user", "", "Username for the sensu api")
	checkSensuSilencesCmd.Flags().StringVar(&apiPassword, "password", "", "Password for the sensu api")
	checkSensuSilencesCmd.Flags().StringVar(&apiHost, "host", "", "Sensu Api host")
	checkSensuSilencesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
