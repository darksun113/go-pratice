// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const sample = 1000
	var r [sample]int
	for i := 0; i < sample; i++ {
		r[i] = rand.Intn(6) + 1
	}

	var static1 [7]int
	var sum = 0.0
	for i := 0; i < sample; i++ {
		static1[r[i]]++
		sum += (float64)(r[i])
	}

	var mean = sum / sample

	fmt.Printf("1: %.2f%%\n", (float32)(static1[1])/(float32)(sample)*100.0)
	fmt.Printf("2: %.2f%%\n", (float32)(static1[2])/(float32)(sample)*100.0)
	fmt.Printf("3: %.2f%%\n", (float32)(static1[3])/(float32)(sample)*100.0)
	fmt.Printf("4: %.2f%%\n", (float32)(static1[4])/(float32)(sample)*100.0)
	fmt.Printf("5: %.2f%%\n", (float32)(static1[5])/(float32)(sample)*100.0)
	fmt.Printf("6: %.2f%%\n", (float32)(static1[6])/(float32)(sample)*100.0)

	fmt.Println(mean)

	rand.Seed(time.Now().UnixNano())
	var p = rand.Perm(1000)
	var means [100]float64
	for i := 0; i < 1000; i += 10 {
		tempMean := 0.0
		for _, r1 := range p[i : i+10] {
			tempMean += (float64)(r[r1])
		}
		tempMean = tempMean / 10.0
		means[i/10] = tempMean
	}

	static2 := make(map[string]int)
	static2["0-0.5"] = 0
	static2["0.5-1"] = 0
	static2["1-1.5"] = 0
	static2["1.5-2"] = 0
	static2["2-2.5"] = 0
	static2["2.5-3"] = 0
	static2["3-3.5"] = 0
	static2["3.5-4"] = 0
	static2["4-4.5"] = 0
	static2["4.5-5"] = 0
	static2["5-5.5"] = 0
	static2["5.5-6"] = 0

	for i := 0; i < 100; i++ {
		if means[i] >= 0 && means[i] <= 0.5 {
			static2["0-0.5"]++
		} else if means[i] > 0.5 && means[i] <= 1 {
			static2["0.5-1"]++
		} else if means[i] > 1 && means[i] <= 1.5 {
			static2["1-1.5"]++
		} else if means[i] > 1.5 && means[i] <= 2 {
			static2["1.5-2"]++
		} else if means[i] > 2 && means[i] <= 2.5 {
			static2["2-2.5"]++
		} else if means[i] > 2.5 && means[i] <= 3 {
			static2["2.5-3"]++
		} else if means[i] > 3 && means[i] <= 3.5 {
			static2["3-3.5"]++
		} else if means[i] > 3.5 && means[i] <= 4 {
			static2["3.5-4"]++
		} else if means[i] > 4 && means[i] <= 4.5 {
			static2["4-4.5"]++
		} else if means[i] > 4.5 && means[i] <= 5 {
			static2["4.5-5"]++
		} else if means[i] > 5 && means[i] <= 5.5 {
			static2["5-5.5"]++
		} else if means[i] > 5.5 && means[i] <= 6 {
			static2["5.5-6"]++
		} else {
			static2["others"]++
		}
	}

	max := 0

	//keys := make([]string, 0, len(static2))
	for col := range static2 {
		if static2[col] > max {
			max = static2[col]
		}
	}

	for i := max; i > 0; i-- {
		fmt.Print(i)
		fmt.Print("\t")
		if static2["0-0.5"] >= i {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
		fmt.Print("\t")
		if static2["0.5-1"] >= i {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
		fmt.Print("\t")
		if static2["1-1.5"] >= i {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
		fmt.Print("\t")
		if static2["1.5-2"] >= i {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
		fmt.Print("\t")
		if static2["2-2.5"] >= i {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
		fmt.Print("\t")
		if static2["2.5-3"] >= i {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
		fmt.Print("\t")
		if static2["3-3.5"] >= i {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
		fmt.Print("\t")
		if static2["3.5-4"] >= i {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
		fmt.Print("\t")
		if static2["4-4.5"] >= i {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
		fmt.Print("\t")
		if static2["4.5-5"] >= i {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
		fmt.Print("\t")
		if static2["5-5.5"] >= i {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
		fmt.Print("\t")
		if static2["5.5-6"] >= i {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
		fmt.Print("\t")
		if static2["others"] >= i {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
		fmt.Print("\t")
		fmt.Println()
	}

	fmt.Println("0\t0.5\t1\t1.5\t2\t2.5\t3\t3.5\t4\t4.5\t5\t5.5\t6\tothers\t")

}
