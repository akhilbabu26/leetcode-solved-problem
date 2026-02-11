package main
import "fmt"

// func maxNumberOfBalloons(text string) int {
// 	result := []string{}
// 	m := make(map[string]int)
// 	var count int

//     for _, v := range text{
// 		if string(v) == "b" ||
// 			string(v) == "a" ||
// 			string(v) == "l" ||
// 			string(v) == "o" ||
// 			string(v) == "n"{
// 				result = append(result, string(v))
// 			}
//     }

// 	for _,v := range result{
// 		m[v] = m[v] + 1
// 	}

// 	if m["b"] == 0{
// 		count = m["b"]
// 	}

// 	if m["b"] == 1 && m["a"] > 1 && m["l"] > 1 && m["o"] > 1 && m["n"] > 1{
// 		count = m["b"]
// 	}

// 	if m["b"] == m["a"] && m["b"] == m["n"] && m["o"] == m["l"]{
// 		count = m["b"] 
// 	}

// 	if m["b"] > m["a"] && m["b"] > m["n"] && m["l"] != m["b"] + m["b"] && m["n"] != m["b"] + m["b"]	{
		
// 	}
// 	return count	
// }
func maxNumberOfBalloons(text string) int {

	cnt := make([]int, 26)

	for i := 0; i < len(text); i++ {
		cnt[text[i]-'a']++
	}

	b := cnt['b'-'a']
	a := cnt['a'-'a']
	l := cnt['l'-'a'] / 2
	o := cnt['o'-'a'] / 2
	n := cnt['n'-'a']

	res := b

	if a < res {
		res = a
	}
	if l < res {
		res = l
	}
	if o < res {
		res = o
	}
	if n < res {
		res = n
	}

	return res
}


func main() {
    r := maxNumberOfBalloons("loonbalxallpoon")
    fmt.Println(r)
}
