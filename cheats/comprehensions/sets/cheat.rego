package cheat

set_match_vowels := {match |
  some letter in letters
  some vowel in vowels
  letter == vowel
  match := letter
}
