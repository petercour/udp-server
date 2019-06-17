// udp socket server
// https://golangr.com
package main
 
import (
    "fmt"
    "net"
    "os"
)
 
/* Error checking */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}
 
func main() {
    /* Open at port 8000 */
    fmt.Println("Server started at port 8000")
    ServerAddr,err := net.ResolveUDPAddr("udp",":8000")
    CheckError(err)
 
    /* Now listen at selected port */
    fmt.Println("Listening..")
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close()
 
    buf := make([]byte, 1024)

    /* Read data */
    for {
        n,addr,err := ServerConn.ReadFromUDP(buf)
        fmt.Println("Received ",string(buf[0:n]), " from ",addr)
 
        if err != nil {
            fmt.Println("Error: ",err)
        } 
    }
}