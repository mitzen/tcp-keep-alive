// package main

// import (
// 	"fmt"
// 	"net"
// 	"time"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "iam.countdown.co.nz:443")
// 	if err != nil {
// 		fmt.Println("Error dialing:", err)
// 		return
// 	}
// 	defer conn.Close()

// 	tcpConn := conn.(*net.TCPConn)

// 	// Enable keepalives
// 	err = tcpConn.SetKeepAlive(true)
// 	if err != nil {
// 		fmt.Println("Error setting keepalives:", err)
// 		return
// 	}

// 	// Simulate and set keepalive period to more than 5 min
// 	// 16 minutes
// 	err = tcpConn.SetKeepAlivePeriod(1000 * time.Second)
// 	if err != nil {
// 		fmt.Println("Error setting keepalive period:", err)
// 		return
// 	}

// 	// Send data to the server
// 	fmt.Fprintf(conn, "Connected to iam!\n")

// 	// Keep the connection alive for 5 minutes
// 	// without sending anything
// 	for i := 0; i < 10; i++ {
// 		//fmt.Fprintf(conn, "#")
// 		fmt.Println("#")
// 		time.Sleep(60 * time.Second)
// 	}
// }

package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	targetURL := "http://localhost:8080"
	//targetURL := "https://www.countdown.co.nz" // Replace with the actual target URL

	// Create an HTTP client with keep-alive enabled
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false, // Enable keep-alives
			DialContext: (&net.Dialer{
				//Timeout:   30 * time.Second,
				Timeout:   60 * time.Second,
				KeepAlive: 120 * time.Second,
				DualStack: true,
			}).DialContext,
		},
	}

	// Define the keep-alive duration in minutes
	//keepAliveDuration := 10 * time.Minute // Keep connection alive for 5 minutes

	for {
		// Send an HTTP request to the target URL
		resp, err := client.Get(targetURL)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close() // Close the response body to avoid leaks

		// Check if the request was successful
		if resp.StatusCode != http.StatusOK {
			fmt.Println("Request failed:", resp.Status)
			break // Stop the loop if the request fails
		}

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			break // Stop the loop if reading fails
		}

		fmt.Println("Response body:", string(body))

		// // Sleep for the keep-alive duration
		// fmt.Println("Entering sleep mode")
		// time.Sleep(keepAliveDuration)
		// fmt.Println("Completed")
	}
}
