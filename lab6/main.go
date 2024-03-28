/* las. Tablica dwuwymiarowa w której sadzimy losowo drzewa. Losujemy uderzenie pioruna, które niszczy drzewo. Jeśli pole jest puste, nic się nie dzieje. Jeśli pole jest zajęte, drzewo się zapala. Jeśli obok niego jest inne drzewo, również się zapala i tak dalej. Najlepiej funkcja rekurencyjna 
Wypisujemy procent spalenia lasu i procent zalesienia*/
package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strings"
    "time"
)

var matrix [10][10]int

func piorun() {
    x := rand.Intn(10)
    y := rand.Intn(10)
    if matrix[x][y] == 1 {
        matrix[x][y] = 2
        spalDrzewa(x, y)
    }
}

func spalDrzewa(x, y int) {
    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            newX, newY := x+i, y+j
            if newX >= 0 && newX < 10 && newY >= 0 && newY < 10 && matrix[newX][newY] == 1 {
                matrix[newX][newY] = 2
                printMatrix()
                
                spalDrzewa(newX, newY)
                time.Sleep(1 * time.Second)
            }
        }
    }
}

func printMatrix() {
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            switch matrix[i][j] {
            case 0:
                fmt.Print(". ")
            case 1:
                fmt.Print("🌲")
            case 2:
                fmt.Print("🔥")
            }
        }
        fmt.Println()
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())

    for i := 0; i < 70; i++ {
        for {
            x := rand.Intn(10)
            y := rand.Intn(10)

            if matrix[x][y] == 0 {
                matrix[x][y] = 1
                break
            }
        }
    }

    reader := bufio.NewReader(os.Stdin)

    for {
        printMatrix()
        fmt.Print("Naciśnij enter, aby zainicjować uderzenie pioruna lub wpisz 'q', aby zakończyć: ")
        text, _ := reader.ReadString('\n')
        text = strings.Replace(text, "\n", "", -1)

        if strings.Compare("q", text) == 0 {
            break
        }

        piorun()
    }
}