package main

import "fmt"

func isAnagram(s string, t string) bool {
   if len(s) != len(t) {
     return false
   } 


   sMap := make(map[rune]int, len(s)) 
   tMap := make(map[rune]int, len(t))

   for i := 0; i < len(s); i++ {
     rCs := rune(s[i]) 
     rCt := rune(t[i])

     sVal, _ := sMap[rCs]
     tVal, _ := tMap[rCt]

     sMap[rCs] = sVal+1
     tMap[rCt] = tVal+1
   }

   for sky, sVal := range sMap {
     tVal, inT := tMap[sky]

     if !inT || tVal != sVal {
       return false
     }
   }

   return true
}

func main() {
  fmt.Printf("is Anagram %v\n", isAnagram("dry", "rd"))
}