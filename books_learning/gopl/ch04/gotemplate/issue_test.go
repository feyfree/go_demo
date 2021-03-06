// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub gotemplate matching the search terms.
package gotemplate

import (
	"fmt"
	"go_demo/books_learning/gopl/ch04/github"
	"log"
	"testing"
)

//!+
func TestIssues(t *testing.T) {
	//result, err := github.SearchIssues(os.Args[1:])
	terms := []string{"repo:golang/go", "is:open", "json", "decoder"}
	result, err := github.SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d gotemplate:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/gotemplate
$ ./gotemplate repo:golang/go is:open json decoder
13 gotemplate:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
