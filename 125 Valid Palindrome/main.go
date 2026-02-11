package main
import "fmt"

func isPalindrome(s string) bool {
    
    if s == ""{
        return true
    }

    word := ""
    var word2 []rune
    word3 := ""
    result := false

    for _,v := range s{
        if v < 48 || v > 57 && v < 65 || v > 90 && v < 97 || v > 122{
            continue
        }
        if v >= 65 && v <= 90{
            word += string(v + 32)
            word2 = append(word2, v + 32)
        }else{
            word += string(v)
            word2 = append(word2, v)
        }
    }
    
    for i := len(word2) - 1; i >= 0; i--{
        sc := string(word2[i])
        word3 += sc
    }
    
    if word == word3{
        result = true
    }

    return result
}

func main() { 
	result := isPalindrome("A man, a plan, a canal: Panama")
  	fmt.Println(result)
}