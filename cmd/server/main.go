package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	list, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	defer list.Close()

	for {
		conn, err := list.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		go handle(conn)

	}
}

func handle(con net.Conn) {

	scanner := bufio.NewScanner(con)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Println(con, "I heard you say: %s\n", ln)
	}

	defer con.Close()
	fmt.Println(" Code got here !!")

}



// func handleConnection(conn net.Conn) {

// 	defer conn.Close()

// 	buf := make([]byte, 1024)

// 	_, err := conn.Read(buf)

// 	if err != nil {
// 		return
// 	}

// 	// taking request

// 	// helper method http.Request !!

// 	// log.Println(string(buf[:n]))

// 	// first read hello , second read  second , third read hello
// 	// print garda hellod at 3rd time

// 	// data := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: "+string(length())"\r\n<p>Hello World!</p>"
// 	// 1. Define your HTML content
// 	// data := Data{
// 	// 	Name: "Ram",
// 	// 	Age:  28,
// 	// }

// 	// body, _ := json.Marshal(data)

// 	userDetailPage := pages.UserDetails("1")

// 	var resBuf bytes.Buffer

// 	userDetailPage.Render(context.Background(), &resBuf)

// 	// 2. Format the full HTTP response string
// 	// Every line must end with \r\n (CRLF)

// 	response := fmt.Sprintf(
// 		"HTTP/1.1 200 OK\r\n"+
// 			"Content-Type: text/html; charset=utf-8\r\n"+
// 			"Content-Length: %d\r\n"+
// 			"Connection: close\r\n"+
// 			"\r\n"+ // Mandatory blank line between headers and body
// 			"%s",
// 		resBuf.Len(), resBuf.String())

// 	// 3. Write the raw string to the connection
// 	conn.Write([]byte(response))

// 	// userDetailPage.Render(context.Background())

// 	// n, err := conn.Write([]byte(data))
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// }
// 	// log.Println("Response send successfully!!", n)

// }

// // HTTP/1.1 301 Moved Permanently
// // Location: http://www.google.com/
// // Content-Type: text/html; charset=UTF-8
// // Content-Security-Policy-Report-Only: object-src 'none';base-uri 'self';script-src 'nonce-Xd4_VH6gBgQid4CfbmRYWA' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp
// // Date: Mon, 04 May 2026 09:33:11 GMT
// // Expires: Wed, 03 Jun 2026 09:33:11 GMT
// // Cache-Control: public, max-age=2592000
// // Server: gws
// // Content-Length: 219
// // X-XSS-Protection: 0
// // X-Frame-Options: SAMEORIGIN

// // <HTML><HEAD><meta http-equiv="content-type" content="text/html;charset=utf-8">
// // <TITLE>301 Moved</TITLE></HEAD><BODY>
// // <H1>301 Moved</H1>
// // The document has moved
// // <A HREF="http://www.google.com/">here</A>.
// // </BODY></HTML>
