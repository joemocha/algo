/*
Sorts a sequence of strings from standard input using heap sort.


% more tiny.txt

S O R T E X A M P L E

% go run cmd/heap/main.go < tiny.txt

A E E L M O P R S T X                 [ one string per line ]



% more words3.txt

bed bug dad yes zoo ... all bad yet


% go run cmd/heap/main.go < words3.txt

all bad bed bug dad ... yes yet zoo    [ one string per line ]
*/

package main

import (
	"fmt"

	"github.com/shellfly/algo/algs4"
	"github.com/shellfly/algo/stdin"
)

func main() {
	items := stdin.ReadAllStrings()
	algs4.HeapSort(algs4.StringSlice(items))
	if !algs4.IsSorted(algs4.StringSlice(items)) {
		fmt.Println("Not Sorted")
		fmt.Println(items)
	} else {
		fmt.Println("sorted items: ", items)
	}
}
