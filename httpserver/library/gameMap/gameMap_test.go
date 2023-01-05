package gameMap

import (
	"fmt"
	"testing"
)
 


func TestGameMap(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got, err := gameMap(100)         // 程序输出的结果


	if err != nil{
		fmt.Println(got)
	}
	fmt.Println(got)
}