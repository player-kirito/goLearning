package main

import (
	"fmt"
	"sync"
)

/**
 * 我能理解现在的现状，却不会放弃最后的挣扎
 * @author player-kirito
 * @date 17:50 2023/2/10
 **/

var (
	lock = &sync.Mutex{}
	s *singleton
)

type singleton struct {
}

func getSingleton() *singleton {
	lock.Lock()
	defer lock.Unlock()
	if(s == nil) {
		s = &singleton{}
	}



	return s
}

func main() {

	c1 := getSingleton()
	c2 := getSingleton()
	fmt.Println(c1 == c2)
}
