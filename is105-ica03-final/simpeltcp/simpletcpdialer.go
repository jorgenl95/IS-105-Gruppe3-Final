
package main

import (
	"fmt"
	"io/ioutil"
	"net"

)

// Kobler seg opp mot en server.
func main () {

	// conn = connection.
	conn, err := net.Dial("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Leser data fra server og konverter denne dataen til string og printer den. bs = bytes.
	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs))
}



