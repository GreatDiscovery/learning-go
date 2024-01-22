package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"testing"
)

//github: https://github.com/deckarep/golang-set

func TestSetFrame(t *testing.T) {
	// Create a string-based set of required classes.
	required := mapset.NewSet[string]()
	required.Add("cooking")
	required.Add("english")
	required.Add("math")
	required.Add("biology")

	// Create a string-based set of science classes.
	sciences := mapset.NewSet[string]()
	sciences.Add("biology")
	sciences.Add("chemistry")

	// Create a string-based set of electives.
	electives := mapset.NewSet[string]()
	electives.Add("welding")
	electives.Add("music")
	electives.Add("automotive")

	// Create a string-based set of bonus programming classes.
	bonus := mapset.NewSet[string]()
	bonus.Add("beginner go")
	bonus.Add("python for dummies")

	//Create a set of all unique classes. Sets will automatically deduplicate the same data.
	all := required.Union(sciences).Union(electives).Union(bonus)
	fmt.Println(all)
	fmt.Println(all.Contains("cooking"))
}

func TestSetDifference(t *testing.T) {
	// Create a string-based set of required classes.
	set1 := mapset.NewSet[string]()
	set1.Add("cooking")
	set1.Add("english")
	set1.Add("math")
	set1.Add("biology")

	set2 := mapset.NewSet[string]()
	set2.Add("cooking")
	set2.Add("english")
	set2.Add("math")
	set2.Add("biology")
	set2.Add("diff1")
	set2.Add("diff2")

	difference := set2.Difference(set1)
	fmt.Println(difference)
	slice := difference.ToSlice()
	fmt.Println(slice)
}
