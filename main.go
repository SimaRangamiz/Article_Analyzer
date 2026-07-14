package main

import "fmt"

func main() {
	
	samplebody := "Artificial intelligence or ai is changing the world today. Many businesses use ai to automate tasks and improve efficiency. Because ai can learn from historical data, ai models become smarter over time. In healthcare, ai helps doctors diagnose diseases faster. Developers are now building new applications integrated with ai to solve complex problems every day."

	top_tags := extract(samplebody, 3)

	fmt.Println("\nTop Tags:")
	for i := 0; i < len(top_tags); i++ {
		fmt.Printf("%v. Tag: %v\n", i+1, top_tags[i])
	}
}