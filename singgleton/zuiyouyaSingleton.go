package main

import (
	"fmt"
	"sync"
)

/**
 * 我能理解现在的现状，却不会放弃最后的挣扎
 * @author player-kirito
 * @date 18:03 2023/2/10
 **/

var (
	s *singleton
	once = &sync.Once{}
)

type singleton struct {}

func getSingleton() *singleton {
	once.Do(func() {
		s = &singleton{}
	})
	return s
}

func main() {
	c1 := getSingleton()
	c2 := getSingleton()
	fmt.Println(c1 == c2)
}