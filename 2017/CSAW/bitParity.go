package main
import (
    "net"
    "os"
//    "encoding/json"
//    "fmt"
    "strings"
    "strconv"
//    "time"
//    "math/rand"
)

type creditcard struct{
    CreditCard cc
}
type cc struct{
    IssuingNetwork string
    CardNumber uint64
}

func main2() {
	intthing, err := strconv.ParseInt("00001011", 2, 32)
	println("00001011", intthing, err)
}
func main() {
    
    strEcho := "-111\n"
    servAddr := "misc.chal.csaw.io:4239"

    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
        println("ResolveTCPAddr failed:", err.Error())
        os.Exit(1)
    }

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        println("Dial failed:", err.Error())
        os.Exit(1)
    }


    reply := make([]byte, 1024)
















	for i:=0; i<3000; i++ {
	    strEcho = "0\n"
	    _, err = conn.Read(reply)
	    if err != nil {   
		println("Write to server failed:", err.Error())
		os.Exit(1)
	    }
	    //println(i,":  reply from server=", string(reply))
	    number := "WARG!"
	    if strings.Contains(string(reply)," 8-1-1 even parity") {
		number = strings.Split(string(reply), "\n")[2]

	    } else {
		number = strings.Split(string(reply), "\n")[0]
	    }
		numOfOnes := len(strings.Split(string(number[:len(number)-1]), "1"))
		lastDigit := number[len(number)-1:]
		//intNumber, err := strconv.ParseInt( number[:len(number)-1], 2, 32)
		//println(".",number,".",numOfOnes, lastDigit,intNumber)
		//print(intNumber, " ")
		if(lastDigit == "1" && numOfOnes%2==1)  || (lastDigit=="0" && numOfOnes%2==0) {
			strEcho = "1\n"
		print( number[1:len(number)-2], " ")
		}
		//println("write to server = ", strEcho)

		_, err = conn.Write( []byte( strEcho ) )

		if err!=nil {
			println("error!", err)
		}
	}










    conn.Close()
}


