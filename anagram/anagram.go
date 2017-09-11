package anagram

import (
  "strings"
  "sort"
)

const testVersion = 2

func Detect(s string, sl []string) []string {
  anagrams := make([]string, 0)
  strSlice := strings.Split(strings.ToLower(s), "")
  sort.Strings(strSlice)
  for _, word := range sl {
    wordSlice := strings.Split(strings.ToLower(word), "")
    sort.Strings(wordSlice)
    if strings.Join(strSlice, "") == strings.Join(wordSlice,"") {
      anagrams = append(anagrams, word)
    }
  }
  return anagrams
}

