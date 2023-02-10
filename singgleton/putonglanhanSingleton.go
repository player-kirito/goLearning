package main

import "fmt"

/**
 * 我能理解现在的现状，却不会放弃最后的挣扎
 * @author player-kirito
 * @date 17:50 2023/2/10
 **/
// 普通单例模式
var s *singleton

// 太懒太笨笨了 每次要才去创建它
type singleton struct {
}

//func init() {
//	s = &singleton{}
//}

func getSingleton() *singleton {
	if s == nil {
		s = &singleton{}
	}
	return s
}

func main() {
	c1 := getSingleton()
	c2 := getSingleton()
	fmt.Println(c1 == c2)
}
