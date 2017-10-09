////////////////////////////////////////////////////////////////////////////
// Program: dnstools
// Purpose: DNS Tools
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package curlurl

import (
	"fmt"
)

func ExampleURLGlob_GetURLs_output() {
	getURLsTest("google.ca")
	getURLsTest("site.{one,two,three}.com")
	getURLsTest("site.{one,two,three}")
	getURLsTest("site[1-3].com")
	getURLsTest("site.{one,two,three}[1-3].com")
	getURLsTest("site.{one,two,three}[1-3]")
	getURLsTest("{one,two,three}[1-3]")
	getURLsTest("https://site.{one,two,three}.com/prj[1-3]-{one,two,three}/[8-10]")
	// Output:
	// -
	// google.ca
	// [{0 [google.ca]}]
	// [google.ca]
	// -
	// site.{one,two,three}.com
	// [{0 [site.]} {1 [one two three]} {0 [.com]}]
	// [site.one.com site.two.com site.three.com]
	// -
	// site.{one,two,three}
	// [{0 [site.]} {1 [one two three]}]
	// [site.one site.two site.three]
	// -
	// site[1-3].com
	// [{0 [site]} {2 [1 2 3]} {0 [.com]}]
	// [site1.com site2.com site3.com]
	// -
	// site.{one,two,three}[1-3].com
	// [{0 [site.]} {1 [one two three]} {2 [1 2 3]} {0 [.com]}]
	// [site.one1.com site.one2.com site.one3.com site.two1.com site.two2.com site.two3.com site.three1.com site.three2.com site.three3.com]
	// -
	// site.{one,two,three}[1-3]
	// [{0 [site.]} {1 [one two three]} {2 [1 2 3]}]
	// [site.one1 site.one2 site.one3 site.two1 site.two2 site.two3 site.three1 site.three2 site.three3]
	// -
	// {one,two,three}[1-3]
	// [{1 [one two three]} {2 [1 2 3]}]
	// [one1 one2 one3 two1 two2 two3 three1 three2 three3]
	// -
	// https://site.{one,two,three}.com/prj[1-3]-{one,two,three}/[8-10]
	// [{0 [https://site.]} {1 [one two three]} {0 [.com/prj]} {2 [1 2 3]} {0 [-]} {1 [one two three]} {0 [/]} {2 [8 9 10]}]
	// [https://site.one.com/prj1-one/8 https://site.one.com/prj1-one/9 https://site.one.com/prj1-one/10 https://site.one.com/prj1-two/8 https://site.one.com/prj1-two/9 https://site.one.com/prj1-two/10 https://site.one.com/prj1-three/8 https://site.one.com/prj1-three/9 https://site.one.com/prj1-three/10 https://site.one.com/prj2-one/8 https://site.one.com/prj2-one/9 https://site.one.com/prj2-one/10 https://site.one.com/prj2-two/8 https://site.one.com/prj2-two/9 https://site.one.com/prj2-two/10 https://site.one.com/prj2-three/8 https://site.one.com/prj2-three/9 https://site.one.com/prj2-three/10 https://site.one.com/prj3-one/8 https://site.one.com/prj3-one/9 https://site.one.com/prj3-one/10 https://site.one.com/prj3-two/8 https://site.one.com/prj3-two/9 https://site.one.com/prj3-two/10 https://site.one.com/prj3-three/8 https://site.one.com/prj3-three/9 https://site.one.com/prj3-three/10 https://site.two.com/prj1-one/8 https://site.two.com/prj1-one/9 https://site.two.com/prj1-one/10 https://site.two.com/prj1-two/8 https://site.two.com/prj1-two/9 https://site.two.com/prj1-two/10 https://site.two.com/prj1-three/8 https://site.two.com/prj1-three/9 https://site.two.com/prj1-three/10 https://site.two.com/prj2-one/8 https://site.two.com/prj2-one/9 https://site.two.com/prj2-one/10 https://site.two.com/prj2-two/8 https://site.two.com/prj2-two/9 https://site.two.com/prj2-two/10 https://site.two.com/prj2-three/8 https://site.two.com/prj2-three/9 https://site.two.com/prj2-three/10 https://site.two.com/prj3-one/8 https://site.two.com/prj3-one/9 https://site.two.com/prj3-one/10 https://site.two.com/prj3-two/8 https://site.two.com/prj3-two/9 https://site.two.com/prj3-two/10 https://site.two.com/prj3-three/8 https://site.two.com/prj3-three/9 https://site.two.com/prj3-three/10 https://site.three.com/prj1-one/8 https://site.three.com/prj1-one/9 https://site.three.com/prj1-one/10 https://site.three.com/prj1-two/8 https://site.three.com/prj1-two/9 https://site.three.com/prj1-two/10 https://site.three.com/prj1-three/8 https://site.three.com/prj1-three/9 https://site.three.com/prj1-three/10 https://site.three.com/prj2-one/8 https://site.three.com/prj2-one/9 https://site.three.com/prj2-one/10 https://site.three.com/prj2-two/8 https://site.three.com/prj2-two/9 https://site.three.com/prj2-two/10 https://site.three.com/prj2-three/8 https://site.three.com/prj2-three/9 https://site.three.com/prj2-three/10 https://site.three.com/prj3-one/8 https://site.three.com/prj3-one/9 https://site.three.com/prj3-one/10 https://site.three.com/prj3-two/8 https://site.three.com/prj3-two/9 https://site.three.com/prj3-two/10 https://site.three.com/prj3-three/8 https://site.three.com/prj3-three/9 https://site.three.com/prj3-three/10]

}

func ExampleURLGlob_GetURLs_errors() {
	getURLsTest("site[8].com")
	getURLsTest("site[8-].com")
	getURLsTest("site[8-a].com")
	// Output:
	// -
	// site[8].com
	// Error, Range expression parsing: Wrong range format
	// -
	// site[8-].com
	// Error, Range expression num range not match: strconv.Atoi: parsing "": invalid syntax
	// [{0 [site]} {2 []} {0 [.com]}]
	// []
	// -
	// site[8-a].com
	// Error, Range expression num range not match: strconv.Atoi: parsing "a": invalid syntax
	// [{0 [site]} {2 []} {0 [.com]}]
	// []

}

func getURLsTest(hostn string) {
	// https://github.com/golang/go/wiki/PanicAndRecover
	defer func() {
		if r := recover(); r != nil {
			_, ok := r.(error)
			if !ok {
				fmt.Printf("pkg: %v", r)
			}
		}
	}()

	fmt.Printf("-\n%s\n", hostn)
	g := NewURLGlob(hostn).Parse(abortOn)
	fmt.Println(g.urlGlob)
	fmt.Println(g.GetURLs(0))
}

func abortOn(errCase string, e error) {
	if e != nil {
		fmt.Printf("Error, %s: %v\n", errCase, e)
	}
}
