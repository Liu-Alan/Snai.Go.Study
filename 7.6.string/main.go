package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		fmt.Println(strings.Contains("seafood", "foo"))
		fmt.Println(strings.Contains("seafood", "bar"))
		fmt.Println(strings.Contains("seafood", ""))
		fmt.Println(strings.Contains("", ""))
		//true
		//false
		//true
		//true
	*/
	/*
		s := []string{"foo", "bar", "baz"}
		fmt.Println(strings.Join(s, ","))
		//Output:foo, bar, baz
	*/
	/*
		fmt.Println(strings.Index("chicken", "ken"))
		fmt.Println(strings.Index("chicken", "dmr"))
		//Output:4
		//-1
	*/
	/*
		fmt.Println("ba" + strings.Repeat("na", 2))
		//Output:banana
	*/
	/*
		fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
		fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
		//Output:oinky oinky oink
		//moo moo moo
	*/
	/*
		fmt.Printf("%q\n", strings.Split("a,b,c", ","))
		fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
		fmt.Printf("%q\n", strings.Split(" xyz ", ""))
		fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
		//Output:["a" "b" "c"]
		//["" "man " "plan " "canal panama"]
		//[" " "x" "y" "z" " "]
		//[""]
	*/
	/*
		fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung !!! ", "! "))
		//Output:["Achtung"]
	*/
	/*
		fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz    "))
		//Output:Fields are: ["foo" "bar" "baz"]
	*/
	/*
		str := make([]byte, 0, 100)
		str = strconv.AppendInt(str, 4567, 10)
		str = strconv.AppendBool(str, false)
		str = strconv.AppendQuote(str, "abcdefg")
		str = strconv.AppendQuoteRune(str, '单')
		fmt.Println(string(str))
	*/
	/*
		a := strconv.FormatBool(false)
		b := strconv.FormatFloat(123.23, 'g', 12, 64)
		c := strconv.FormatInt(1234, 10)
		d := strconv.FormatUint(12345, 10)
		e := strconv.Itoa(1023)
		fmt.Println(a, b, c, d, e)
	*/

	a, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println(err)
	}
	b, err := strconv.ParseFloat("123.23", 64)
	if err != nil {
		fmt.Println(err)
	}
	c, err := strconv.ParseInt("1234", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	d, err := strconv.ParseUint("12345", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	e, err := strconv.Atoi("1023")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a, b, c, d, e)
}
