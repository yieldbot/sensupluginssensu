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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// AcquireLocalChecks will retrieve the currently running configuration and
// return a list of all checks it knows about
func AcquireLocalChecks() {
	// ?	var jsonOut Message
	var jsonOut interface{}
	// localChecks, err := exec.Command("/opt/sensu/embedded/bin/sensu-client", "-d", "/etc/sensu/conf.d", "-P")

	// out, err := localChecks.Output()
	// if err != nil {
	// 	panic(err)
	// }

	// func main() {
	c1 := exec.Command("/opt/sensu/embedded/bin/sensu-client", "-d", "/etc/sensu/conf.d", "-P")
	c2 := exec.Command("grep", "-v", "timestamp")

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	var b2 bytes.Buffer
	c2.Stdout = &b2

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()
	io.Copy(os.Stdout, &b2)

	b2.WriteTo(os.Stdout)

	err := json.Unmarshal(b2.Bytes(), &jsonOut)
	if err != nil {
		panic(err)
	}

	fmt.Println("test")
	fmt.Println(jsonOut)

	// fmt.Println(jsonOut.redis.port)
}
