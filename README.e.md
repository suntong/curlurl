
# {{.Name}}

{{render "license/shields" . "License" "MIT"}}
{{template "badge/godoc" .}}
{{template "badge/goreport" .}}
{{template "badge/travis" .}}

## {{.Name}} - curl like url glob handling in Go

For input like `site.{one,two,three}[1-3].com`, it gives

	site.one1.com
	site.one2.com
	site.one3.com
	site.two1.com
	site.two2.com
	site.two3.com
	site.three1.com
	site.three2.com
	site.three3.com


For input like `https://site.{one,two,three}.com/prj[1-3]-{one,two,three}/[8-10]`, it gives

	https://site.one.com/prj1-one/8
	https://site.one.com/prj1-one/9
	https://site.one.com/prj1-one/10
	https://site.one.com/prj1-two/8
	https://site.one.com/prj1-two/9
	https://site.one.com/prj1-two/10
	https://site.one.com/prj1-three/8
	https://site.one.com/prj1-three/9
	https://site.one.com/prj1-three/10
	https://site.one.com/prj2-one/8
	https://site.one.com/prj2-one/9
	https://site.one.com/prj2-one/10
	https://site.one.com/prj2-two/8
	https://site.one.com/prj2-two/9
	https://site.one.com/prj2-two/10
	https://site.one.com/prj2-three/8
	https://site.one.com/prj2-three/9
	https://site.one.com/prj2-three/10
	https://site.one.com/prj3-one/8
	https://site.one.com/prj3-one/9
	https://site.one.com/prj3-one/10
	https://site.one.com/prj3-two/8
	https://site.one.com/prj3-two/9
	https://site.one.com/prj3-two/10
	https://site.one.com/prj3-three/8
	https://site.one.com/prj3-three/9
	https://site.one.com/prj3-three/10
	https://site.two.com/prj1-one/8
	https://site.two.com/prj1-one/9
	https://site.two.com/prj1-one/10
	https://site.two.com/prj1-two/8
	https://site.two.com/prj1-two/9
	https://site.two.com/prj1-two/10
	https://site.two.com/prj1-three/8
	https://site.two.com/prj1-three/9
	https://site.two.com/prj1-three/10
	https://site.two.com/prj2-one/8
	https://site.two.com/prj2-one/9
	https://site.two.com/prj2-one/10
	https://site.two.com/prj2-two/8
	https://site.two.com/prj2-two/9
	https://site.two.com/prj2-two/10
	https://site.two.com/prj2-three/8
	https://site.two.com/prj2-three/9
	https://site.two.com/prj2-three/10
	https://site.two.com/prj3-one/8
	https://site.two.com/prj3-one/9
	https://site.two.com/prj3-one/10
	https://site.two.com/prj3-two/8
	https://site.two.com/prj3-two/9
	https://site.two.com/prj3-two/10
	https://site.two.com/prj3-three/8
	https://site.two.com/prj3-three/9
	https://site.two.com/prj3-three/10
	https://site.three.com/prj1-one/8
	https://site.three.com/prj1-one/9
	https://site.three.com/prj1-one/10
	https://site.three.com/prj1-two/8
	https://site.three.com/prj1-two/9
	https://site.three.com/prj1-two/10
	https://site.three.com/prj1-three/8
	https://site.three.com/prj1-three/9
	https://site.three.com/prj1-three/10
	https://site.three.com/prj2-one/8
	https://site.three.com/prj2-one/9
	https://site.three.com/prj2-one/10
	https://site.three.com/prj2-two/8
	https://site.three.com/prj2-two/9
	https://site.three.com/prj2-two/10
	https://site.three.com/prj2-three/8
	https://site.three.com/prj2-three/9
	https://site.three.com/prj2-three/10
	https://site.three.com/prj3-one/8
	https://site.three.com/prj3-one/9
	https://site.three.com/prj3-one/10
	https://site.three.com/prj3-two/8
	https://site.three.com/prj3-two/9
	https://site.three.com/prj3-two/10
	https://site.three.com/prj3-three/8
	https://site.three.com/prj3-three/9
	https://site.three.com/prj3-three/10

## Author(s) & Contributor(s)

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

All patches welcome. 
