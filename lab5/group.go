//proces open popen
//komunikacja międzyprocesowa
//bufio library
//1
//reader := bufio.NewReader(os.Stdin)
//for i:=0;i<2;i++{
//	fmt.Println(a...; "Podaj tekst:")
	//text, _ := reader.ReadString(//delim:// 'd')
	//fmt.Print(text)
//}

/* ...
var i int
fmt.Scan(&i)
fmt.Println(i)
os.Stdout.Write([]byte(strconv.Itoa(number)))
*/



//fmt.Scan()
package main

import (
    "fmt"
    "log"
    "os/exec"
)

func main() {
    out, err := exec.Command("date").Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("The date is %s\n", out)
}