package stringdivisor

import (
	"testing"
)

/*
Example 1:

Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
Example 2:

Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"
Example 3:

Input: str1 = "LEET", str2 = "CODE"
Output: ""
*/


func TestGcdCase1(t *testing.T) {
	s := gcdOfStrings("ABCABC", "ABC")
	if s != "ABC" {
		t.Fail()
	}
}

func TestGcdCase2(t *testing.T) {
	s := gcdOfStrings("ABABAB", "ABAB")
	if s != "AB" {
		t.Fail()
	}
}

func TestGcdCase3(t *testing.T) {
	s := gcdOfStrings("LEET", "CODE")
	if s != "" {
		t.Fail()
	}
}

func TestGcdCase4(t *testing.T) {
	s := gcdOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX")
	if s != "TAUXX" {
		t.Fail()
	}
}

func TestCutCase1(t *testing.T) {
	s := EnhancedString("ABCABC")
	s = s.Cut(0, 3)
	if string(s) != "ABC" {
		t.Fail()
	}
}	

func TestCutCase2(t *testing.T) {
	s := EnhancedString("ABCABC")
	s = s.Cut(0, 3).Cut(1, 3)
	if string(s) != "BC" {
		t.Fail()
	}
}	

func TestIsDividedByCase1(t *testing.T) {
	es := EnhancedString("ABCABC")
	if !es.isDividedBy("ABC") {
		t.Fail()
	}
}

