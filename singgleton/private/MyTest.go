package main

import "singgleton/private"

/**
 * 我能理解现在的现状，却不会放弃最后的挣扎
 * @author player-kirito
 * @date 20:13 2023/2/10
 **/

/**
 * @author player-kirito
 * @description 公有的是大写开头的 如果在其他包可以通过包名调用 如果在本包可以直接调用 私有的是小写开头的 可以在包内直接调用 但是在包外无法调用
 * @date 20:16 2023/2/10
 * @param nil
 * @return nil
 **/

func main() {
	private.PrivateTest()
	//private.privateTest() // 检测不到
}