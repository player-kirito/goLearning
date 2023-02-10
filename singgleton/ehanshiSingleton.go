package main

import "fmt"

/**
 * 我能理解现在的现状，却不会放弃最后的挣扎
 * @author player-kirito
 * @date 17:50 2023/2/10
 **/
type singleton struct {
}

var s *singleton

// init 会在main前执行
func init() {
	s = &singleton{}
}

func getSingleton() *singleton {
	return s
}

func main() {
	c1 := getSingleton()
	c2 := getSingleton()

	fmt.Println(c1 == c2)
}
