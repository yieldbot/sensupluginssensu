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

package cmd

import "github.com/spf13/cobra"

// var apiHost string
// var apiUser string
// var apiPassword string
// var bodyString string
// var bodyBytes []byte

// Tech Debt
// YELLOW needs clean refactor
// YELLOW needs unit tests

// handlerRemoveExpiredChecksCmd represents the handlerRemoveExpiredChecks command
var handlerRemoveExpiredChecksCmd = &cobra.Command{
	Use:   "handlerRemoveExpiredChecks",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,

	Run: func(cmd *cobra.Command, args []string) {

		AcquireLocalChecks()

		// if apiHost == "" {
		//   apiHost = viper.GetString("sensupluginssensu.handlerRemoveExpiredChecks.apiHost")
		// }
		// if apiUser == "" {
		//   apiUser = viper.GetString("sensupluginssensu.handlerRemoveExpiredChecks.apiUser")
		// }
		// if apiPassword == "" {
		//   apiPassword = viper.GetString("sensupluginssensu.handlerRemoveExpiredChecks.apiPassword")
		// }
		//
		// tr := &http.Transport{
		//   TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		// }
		//
		// client := &http.Client{Transport: tr}
		//
		// req, _ := http.NewRequest("GET", "https://"+apiHost+":4567/stashes", nil)
		// req.SetBasicAuth(apiUser, apiPassword)
		// resp, _ := client.Do(req)
		// if resp.StatusCode == 200 {
		//   bodyBytes, _ = ioutil.ReadAll(resp.Body)
		//   bodyString = string(bodyBytes)
		// }
		// defer resp.Body.Close()
		//
		// if strings.Contains(bodyString, "\"expire\":-1") {
		//
		//   sensuutil.Exit("critical", bodyString)
		// } else {
		//   sensuutil.Exit("ok")
		// }
		//
	},
}

func init() {
	RootCmd.AddCommand(handlerRemoveExpiredChecksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// handlerRemoveExpiredChecksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// handlerRemoveExpiredChecksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
