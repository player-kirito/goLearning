package main
/**
 * 我能理解现在的现状，却不会放弃最后的挣扎
 * @author player-kirito
 * @date 18:02 2023/2/10
 **/
import (
	"fmt"
	"sync"
)

var (
	lock = &sync.Mutex{}
	s *singleton
)

type singleton struct {}

func getSingleton() *singleton {
	if(s == nil) {
		lock.Lock()
		defer lock.Unlock()
		if(s == nil) {
			fmt.Println("是哦")
			s = &singleton{}
		}
	}
	return s
}
func main() {
	c1 := getSingleton()
	c2 := getSingleton()
	c3 := getSingleton()
	fmt.Println(c1 == c2 && c2 == c3)
}