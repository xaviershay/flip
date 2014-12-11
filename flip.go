package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
)

func main() {
	var flips = map[string]string{
		"\u0021": "\u00A1",
		"\u0022": "\u201E",
		"\u0026": "\u214B",
		"\u0027": "\u002C",
		"\u0028": "\u0029",
		"\u002E": "\u02D9",
		"\u0033": "\u0190",
		"\u0034": "\u152D",
		"\u0036": "\u0039",
		"\u0037": "\u2C62",
		"\u003B": "\u061B",
		"\u003C": "\u003E",
		"\u003F": "\u00BF",
		"\u0041": "\u2200",
		"\u0042": "\u1041",
		"\u0043": "\u2183",
		"\u0044": "\u25D6",
		"\u0045": "\u018E",
		"\u0046": "\u2132",
		"\u0047": "\u2141",
		"\u004A": "\u017F",
		"\u004B": "\u22CA",
		"\u004C": "\u2142",
		"\u004D": "\u0057",
		"\u004E": "\u1D0E",
		"\u0050": "\u0500",
		"\u0051": "\u038C",
		"\u0052": "\u1D1A",
		"\u0054": "\u22A5",
		"\u0055": "\u2229",
		"\u0056": "\u1D27",
		"\u0059": "\u2144",
		"\u005B": "\u005D",
		"\u005F": "\u203E",
		"\u0061": "\u0250",
		"\u0062": "\u0071",
		"\u0063": "\u0254",
		"\u0064": "\u0070",
		"\u0065": "\u01DD",
		"\u0066": "\u025F",
		"\u0067": "\u0183",
		"\u0068": "\u0265",
		"\u0069": "\u0131",
		"\u006A": "\u027E",
		"\u006B": "\u029E",
		"\u006C": "\u0283",
		"\u006D": "\u026F",
		"\u006E": "\u0075",
		"\u0072": "\u0279",
		"\u0074": "\u0287",
		"\u0076": "\u028C",
		"\u0077": "\u028D",
		"\u0079": "\u028E",
		"\u007B": "\u007D",
		"\u203F": "\u2040",
		"\u2045": "\u2046",
		"\u2234": "\u2235",
		" ":      " ",
		"x":      "x",
		"o":      "o",
		"p":      "d",
		"u":      "n",
		"s":      "s",
		"q":      "b",
	}

	tablePtr := flag.Bool("table", false, "")
	putbackPtr := flag.Bool("putback", false, "")
	flag.Parse()

	if *putbackPtr {
		if *tablePtr {
			fmt.Println("┬─┬ ノ( ゜-゜ノ)")
		} else {
			fmt.Printf("%s ノ( ゜-゜ノ)\n", strings.Join(flag.Args(), " "))
		}
		return
	}

	if *tablePtr {
		fmt.Println("(╯°□°)╯︵ ┻━┻ ")
		return
	}

	// fmt.Printf("%s -> %s", "\u0021", flips["\u0021"])
	buffer := bytes.Buffer{}
	for _, arg := range flag.Args() {
		for _, ch := range arg {
			str := flips[string(ch)]
			if str == "" {
				str = string(ch)
			}
			buffer.WriteString(str)
		}
		buffer.WriteString(" ")
	}
	fmt.Printf("(╯°□°）╯︵ %s\n", Reverse(buffer.String()))
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
