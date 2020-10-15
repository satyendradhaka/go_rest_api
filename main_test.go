package main

import "testing"

func TestSearchCase( t *testing.T){
	searchResult :=caseInsensitive("aBcd", "abC") //for 2 string search eg: abC in aBcd
	if searchResult==false{
		t.Errorf("case insensitive search failed, expected %v, got %v", "true", "false")
	}
}