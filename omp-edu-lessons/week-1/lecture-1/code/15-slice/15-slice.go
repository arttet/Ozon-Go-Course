package main

import "fmt"

func main() {
	allNewsPosts := []string{
		"new title 1",
		"new title 2",
		"new title 3",
		"new title 4",
		"new title 5",
	}

	fmt.Println("<mainpage>")
	showMainPage(allNewsPosts[:3]) // [:3:3]
	fmt.Println("</mainpage>")

	fmt.Println("<all>")
	showPosts(allNewsPosts)
	fmt.Println("</all>")
}

func showMainPage(posts []string) {
	postsWithAds := append(posts, "CLICK HERE TO BUY THIS THING!!!")

	showPosts(postsWithAds)
}

func showPosts(posts []string) {
	for _, post := range posts {
		fmt.Println(post)
	}
}
