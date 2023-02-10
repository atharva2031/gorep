//Exercise for loops in Go

package exercise2
import (
	"fmt"
)
func main(){
	i:=1
	for i<=100{
		if i%3==0{
			fmt.Println(i)
		}
		i++
	}

	var j int64 = 1
	for j<=100{
		if(j%20==0){
			fmt.Println(j)
		}
		j++
	}
}

