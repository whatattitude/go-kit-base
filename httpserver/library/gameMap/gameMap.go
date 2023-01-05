package gameMap

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)


const length  = 20
const width  = 20

func gameMap(tileMapItemMax int)( bool, error){

    var tileMap = [20][20]int{}
    rand.Seed(time.Now().UnixNano())
   
    for i:=0; i<length; i++ {
        for j:=0; j<width; j++ {
        tileMap[i][j] = (rand.Intn(tileMapItemMax))
        }
    }


    fmt.Println("======")

    for i:=0; i<length; i++ {
        for j:=0; j<width; j++ {
            fmt.Printf("%4d", tileMap[i][j])
        }
        fmt.Println("")
    }

	return true, nil
}


func FilePrint(path string, tileMap [][]int)  {
    // create file
    f, err := os.Create(path)
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file
    defer f.Close()
 
    for _, line := range tileMap {
        _, err := fmt.Fprintln(f, "*", line, "*")
        if err != nil {
            log.Fatal(err)
        }
    }
}