package main

import (
    "fmt"
)

func next_character(candidates []int, source, value []byte, index int) ([]int){
    if (index != len(value)) {
        letter := value[index]
        next_l := make([]byte, len(candidates))
        for i := range(candidates) {
            next_l[i] = source[i + index]
        }

        for i := range(candidates) {
            if (next_l[i] == letter) {
                candidates[i] = i
            }
        }

        //TODO increment index
        candidates = next_character(candidates, source, value, index + 1)
    }
    return candidates
}


func string_search(source, value []byte)([]int) {
    size := len(source) - len(value) + 1
    p := make([]int, len(source))

    for i := 0; i < size; i++ {
        p[i] = i;
    }

    return next_character(p, source, value, 0)
}

func main() {
    str1 := "the fireflies blew and drifted in soft random"
    str2 := "and"

    s1 := []byte(str1)
    s2 := []byte(str2)

    fmt.Println(s1, s2)
}
