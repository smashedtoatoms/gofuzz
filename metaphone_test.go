package gofuzz

import (
	"fmt"
	"testing"
)

func TestMetaphoneMetric(t *testing.T) {
	errorMessage := "Unable to Metaphone compare the two values."
	errorMessage2 := "String is empty or non-alphabetic."
	_, e1 := MetaphoneMetric("", "")
	if e1.Error() != errorMessage {
		t.Errorf("MetaphoneMetric didn't compare emtpy strings.")
	}
	_, e2 := MetaphoneMetric("abc", "")
	if e2.Error() != errorMessage2 {
		t.Errorf("MetaphoneMetric didn't compare emtpy string to letters.")
	}
	_, e3 := MetaphoneMetric("", "xys")
	if e3.Error() != errorMessage {
		t.Errorf("MetaphoneMetric didn't compare letters to emtpy string.")
	}
	_, e4 := MetaphoneMetric("123", "123")
	if e4.Error() != errorMessage {
		t.Errorf("MetaphoneMetric didn't compare numeric strings.")
	}
	_, e5 := MetaphoneMetric("123", "")
	if e5.Error() != errorMessage {
		t.Errorf("MetaphoneMetric didn't compare numeric strings to empty string.")
	}
	_, e6 := MetaphoneMetric("", "123")
	if e6.Error() != errorMessage {
		t.Errorf("MetaphoneMetric didn't compare empty string to numeric strings.")
	}
	v1, _ := MetaphoneMetric("dumb", "dum")
	if v1 != true {
		t.Errorf("MetaphoneMetric didn't compare dumb and dum.")
	}
	v2, _ := MetaphoneMetric("smith", "smeth")
	if v2 != true {
		t.Errorf("MetaphoneMetric didn't compare smith and smeth.")
	}
	v3, _ := MetaphoneMetric("merci", "mercy")
	if v3 != true {
		t.Errorf("MetaphoneMetric didn't compare merci and mercy.")
	}
	v4, _ := MetaphoneMetric("dumb", "gum")
	if v4 != false {
		t.Errorf("MetaphoneMetric didn't compare dumb and gum.")
	}
	v5, _ := MetaphoneMetric("smith", "kiss")
	if v5 != false {
		t.Errorf("MetaphoneMetric didn't compare smith and kiss.")
	}
	v6, _ := MetaphoneMetric("merci", "burpy")
	if v6 != false {
		t.Errorf("MetaphoneMetric didn't compare merci and burpy.")
	}
}
func ExampleMetaphoneMetric() {
	result, err := MetaphoneMetric("Colorado", "Kolorado")
	if err != nil {
		fmt.Print("an unexpected error was encountered")
	}
	fmt.Print(result)
	// Output: true
}

func TestMetaphone(t *testing.T) {
	_, er := Metaphone("")
	if er == nil {
		t.Errorf("Metaphone failed to throw an error on a blank string.")
	}
	_, er = Metaphone("123")
	if er == nil {
		t.Errorf("Metaphone failed to throw an error on a numeric string.")
	}
	_, er = Metaphone("abc1")
	if er == nil {
		t.Errorf("Metaphone failed to throw an error on a alpha-numeric string.")
	}

	// z
	z, _ := Metaphone("z")
	if z != "s" {
		t.Errorf("Metaphone 'z' expected 's' and got '%s'.", z)
	}
	z, _ = Metaphone("zz")
	if z != "s" {
		t.Errorf("Metaphone 'zz' expected 's' and got '%s'.", z)
	}

	// y
	y, _ := Metaphone("y")
	if y != "" {
		t.Errorf("Metaphone 'y' expected '' and got '%s'.", y)
	}
	y, _ = Metaphone("zy")
	if y != "s" {
		t.Errorf("Metaphone 'zy' expected 's' and got '%s'.", y)
	}
	y, _ = Metaphone("zyz")
	if y != "ss" {
		t.Errorf("Metaphone 'zyz' expected 'ss' and got '%s'.", y)
	}
	y, _ = Metaphone("zya")
	if y != "sy" {
		t.Errorf("Metaphone 'zya' expected 'sy' and got '%s'.", y)
	}

	// x
	x, _ := Metaphone("x")
	if x != "s" {
		t.Errorf("Metaphone 'x' expected 's' and got '%s'.", x)
	}
	x, _ = Metaphone("zx")
	if x != "sks" {
		t.Errorf("Metaphone 'zx' expected 'sks' and got '%s'.", x)
	}
	x, _ = Metaphone("zxz")
	if x != "skss" {
		t.Errorf("Metaphone 'zxz' expected 'skss' and got '%s'.", x)
	}

	// w
	w, _ := Metaphone("w")
	if w != "" {
		t.Errorf("Metaphone 'w' expected '' and got '%s'.", w)
	}
	w, _ = Metaphone("zw")
	if w != "s" {
		t.Errorf("Metaphone 'zw' expected 's' and got '%s'.", w)
	}
	w, _ = Metaphone("zwz")
	if w != "ss" {
		t.Errorf("Metaphone 'zwz' expected 'ss' and got '%s'.", w)
	}
	w, _ = Metaphone("zwa")
	if w != "sw" {
		t.Errorf("Metaphone 'zwa' expected 'sw' and got '%s'.", w)
	}

	// v
	v, _ := Metaphone("v")
	if v != "f" {
		t.Errorf("Metaphone 'v' expected 'f' and got '%s'.", v)
	}
	v, _ = Metaphone("zv")
	if v != "sf" {
		t.Errorf("Metaphone 'zv' expected 'sf' and got '%s'.", v)
	}
	v, _ = Metaphone("zvz")
	if v != "sfs" {
		t.Errorf("Metaphone 'zvz' expected 'sfs' and got '%s'.", v)
	}

	// u
	u, _ := Metaphone("u")
	if u != "u" {
		t.Errorf("Metaphone 'u' expected 'u' and got '%s'.", u)
	}
	u, _ = Metaphone("zu")
	if u != "s" {
		t.Errorf("Metaphone 'zu' expected 's' and got '%s'.", u)
	}

	// t
	letterT, _ := Metaphone("t")
	if letterT != "t" {
		t.Errorf("Metaphone 't' expected 't' and got '%s'.", letterT)
	}
	letterT, _ = Metaphone("ztiaz")
	if letterT != "sxs" {
		t.Errorf("Metaphone 'ztiaz' expected 'sxs' and got '%s'.", letterT)
	}
	letterT, _ = Metaphone("ztioz")
	if letterT != "sxs" {
		t.Errorf("Metaphone 'ztioz' expected 'sxs' and got '%s'.", letterT)
	}
	letterT, _ = Metaphone("zthz")
	if letterT != "s0s" {
		t.Errorf("Metaphone 'zthz' expected 's0s' and got '%s'.", letterT)
	}
	letterT, _ = Metaphone("ztchz")
	if letterT != "sxs" {
		t.Errorf("Metaphone 'ztchz' expected 'sxs' and got '%s'.", letterT)
	}
	letterT, _ = Metaphone("ztz")
	if letterT != "sts" {
		t.Errorf("Metaphone 'ztz' expected 'sts' and got '%s'.", letterT)
	}

	// s
	s, _ := Metaphone("s")
	if s != "s" {
		t.Errorf("Metaphone 's' expected 's' and got '%s'.", s)
	}
	s, _ = Metaphone("zshz")
	if s != "sxs" {
		t.Errorf("Metaphone 'zshz' expected 'sxs' and got '%s'.", s)
	}
	s, _ = Metaphone("zsioz")
	if s != "sxs" {
		t.Errorf("Metaphone 'zsioz' expected 'sxs' and got '%s'.", s)
	}
	s, _ = Metaphone("zsiaz")
	if s != "sxs" {
		t.Errorf("Metaphone 'zsiaz' expected 'sxs' and got '%s'.", s)
	}
	s, _ = Metaphone("zs")
	if s != "ss" {
		t.Errorf("Metaphone 'zs' expected 'ss' and got '%s'.", s)
	}
	s, _ = Metaphone("zsz")
	if s != "sss" {
		t.Errorf("Metaphone 'zsz' expected 'sss' and got '%s'.", s)
	}

	// r
	r, _ := Metaphone("r")
	if r != "r" {
		t.Errorf("Metaphone 'r' expected 'r' and got '%s'.", r)
	}
	r, _ = Metaphone("zr")
	if r != "sr" {
		t.Errorf("Metaphone 'zr' expected 'sr' and got '%s'.", r)
	}
	r, _ = Metaphone("zrz")
	if r != "srs" {
		t.Errorf("Metaphone 'zrz' expected 'srs' and got '%s'.", r)
	}

	// q
	q, _ := Metaphone("q")
	if q != "k" {
		t.Errorf("Metaphone 'q' expected 'k' and got '%s'.", q)
	}
	q, _ = Metaphone("zq")
	if q != "sk" {
		t.Errorf("Metaphone 'zq' expected 'sk' and got '%s'.", q)
	}
	q, _ = Metaphone("zqz")
	if q != "sks" {
		t.Errorf("Metaphone 'zqz' expected 'sks' and got '%s'.", q)
	}

	// p
	p, _ := Metaphone("p")
	if p != "p" {
		t.Errorf("Metaphone 'p' expected 'p' and got '%s'.", p)
	}
	p, _ = Metaphone("zp")
	if p != "sp" {
		t.Errorf("Metaphone 'zp' expected 'sp' and got '%s'.", p)
	}
	p, _ = Metaphone("zph")
	if p != "sf" {
		t.Errorf("Metaphone 'zph' expected 'sf' and got '%s'.", p)
	}
	p, _ = Metaphone("zpz")
	if p != "sps" {
		t.Errorf("Metaphone 'zpz' expected 'sps' and got '%s'.", p)
	}

	// o
	o, _ := Metaphone("o")
	if o != "o" {
		t.Errorf("Metaphone 'o' expected 'o' and got '%s'.", o)
	}
	o, _ = Metaphone("zo")
	if o != "s" {
		t.Errorf("Metaphone 'zo' expected 's' and got '%s'.", o)
	}

	// n
	n, _ := Metaphone("n")
	if n != "n" {
		t.Errorf("Metaphone 'n' expected 'n' and got '%s'.", n)
	}
	n, _ = Metaphone("zn")
	if n != "sn" {
		t.Errorf("Metaphone 'zn' expected 'sn' and got '%s'.", n)
	}
	n, _ = Metaphone("znz")
	if n != "sns" {
		t.Errorf("Metaphone 'znz' expected 'sns' and got '%s'.", n)
	}

	// m
	m, _ := Metaphone("m")
	if m != "m" {
		t.Errorf("Metaphone 'm' expected 'm' and got '%s'.", m)
	}
	m, _ = Metaphone("zm")
	if m != "sm" {
		t.Errorf("Metaphone 'zm' expected 'sm' and got '%s'.", m)
	}
	m, _ = Metaphone("zmz")
	if m != "sms" {
		t.Errorf("Metaphone 'zmz' expected 'sms' and got '%s'.", m)
	}

	// l
	l, _ := Metaphone("l")
	if l != "l" {
		t.Errorf("Metaphone 'l' expected 'l' and got '%s'.", l)
	}
	l, _ = Metaphone("zl")
	if l != "sl" {
		t.Errorf("Metaphone 'zl' expected 'sl' and got '%s'.", l)
	}
	l, _ = Metaphone("zlz")
	if l != "sls" {
		t.Errorf("Metaphone 'zlz' expected 'sls' and got '%s'.", l)
	}

	// k
	k, _ := Metaphone("k")
	if k != "k" {
		t.Errorf("Metaphone 'k' expected 'k' and got '%s'.", k)
	}
	k, _ = Metaphone("zk")
	if k != "sk" {
		t.Errorf("Metaphone 'zk' expected 'sk' and got '%s'.", k)
	}
	k, _ = Metaphone("zck")
	if k != "sk" {
		t.Errorf("Metaphone 'zck' expected 'sk' and got '%s'.", k)
	}

	// j
	j, _ := Metaphone("j")
	if j != "j" {
		t.Errorf("Metaphone 'j' expected 'j' and got '%s'.", j)
	}
	j, _ = Metaphone("zj")
	if j != "sj" {
		t.Errorf("Metaphone 'zj' expected 'sj' and got '%s'.", j)
	}
	j, _ = Metaphone("zjz")
	if j != "sjs" {
		t.Errorf("Metaphone 'zjz' expected 'sjs' and got '%s'.", j)
	}

	// i
	i, _ := Metaphone("i")
	if i != "i" {
		t.Errorf("Metaphone 'i' expected 'i' and got '%s'.", i)
	}
	i, _ = Metaphone("zi")
	if i != "s" {
		t.Errorf("Metaphone 'zi' expected 's' and got '%s'.", i)
	}

	// h
	h, _ := Metaphone("h") // php wrongly says nil
	if h != "h" {
		t.Errorf("Metaphone 'h' expected 'h' and got '%s'.", h)
	}
	h, _ = Metaphone("zh") // php wrongly says s
	if h != "sh" {
		t.Errorf("Metaphone 'zh' expected 'sh' and got '%s'.", h)
	}
	h, _ = Metaphone("zah")
	if h != "s" {
		t.Errorf("Metaphone 'zah' expected 's' and got '%s'.", h)
	}
	h, _ = Metaphone("zchh")
	if h != "sx" {
		t.Errorf("Metaphone 'zchh' expected 'sx' and got '%s'.", h)
	}
	h, _ = Metaphone("ha")
	if h != "h" {
		t.Errorf("Metaphone 'ha' expected 'h' and got '%s'.", h)
	}

	// g
	g, _ := Metaphone("g")
	if g != "k" {
		t.Errorf("Metaphone 'g' expected 'k' and got '%s'.", g)
	}
	g, _ = Metaphone("zg")
	if g != "sk" {
		t.Errorf("Metaphone 'zg' expected 'sk' and got '%s'.", g)
	}
	g, _ = Metaphone("zgh") // php wrongly says sf
	if g != "skh" {
		t.Errorf("Metaphone 'zgh' expected 'skh' and got '%s'.", g)
	}
	g, _ = Metaphone("zghz") // php wrongly says sfs
	if g != "shs" {
		t.Errorf("Metaphone 'zghz' expected 'shs' and got '%s'.", g)
	}
	g, _ = Metaphone("zgha") // php wrongly says sf
	if g != "sh" {           // others wrongly say skh
		t.Errorf("Metaphone 'zgha' expected 'sh' and got '%s'.", g)
	}
	g, _ = Metaphone("zgn")
	if g != "sn" {
		t.Errorf("Metaphone 'zgn' expected 'sn' and got '%s'.", g)
	}
	g, _ = Metaphone("zgns")
	if g != "skns" {
		t.Errorf("Metaphone 'zgns' expected 'skns' and got '%s'.", g)
	}
	g, _ = Metaphone("zgned") // php wrongly says sknt
	if g != "snt" {
		t.Errorf("Metaphone 'zgned' expected 'snt' and got '%s'.", g)
	}
	g, _ = Metaphone("zgneds") // php wrongly says snts
	if g != "sknts" {
		t.Errorf("Metaphone 'zgneds' expected 'sknts' and got '%s'.", g)
	}
	g, _ = Metaphone("zgi")
	if g != "sj" {
		t.Errorf("Metaphone 'zgi' expected 'sj' and got '%s'.", g)
	}
	g, _ = Metaphone("zgiz")
	if g != "sjs" {
		t.Errorf("Metaphone 'zgiz' expected 'sjs' and got '%s'.", g)
	}
	g, _ = Metaphone("zge")
	if g != "sj" {
		t.Errorf("Metaphone 'zge' expected 'sj' and got '%s'.", g)
	}
	g, _ = Metaphone("zgez")
	if g != "sjs" {
		t.Errorf("Metaphone 'zgez' expected 'sjs' and got '%s'.", g)
	}
	g, _ = Metaphone("zgy")
	if g != "sj" {
		t.Errorf("Metaphone 'zgy' expected 'sj' and got '%s'.", g)
	}
	g, _ = Metaphone("zgyz")
	if g != "sjs" {
		t.Errorf("Metaphone 'zgyz' expected 'sjs' and got '%s'.", g)
	}
	g, _ = Metaphone("zgz")
	if g != "sks" {
		t.Errorf("Metaphone 'zgz' expected 'sks' and got '%s'.", g)
	}

	// f
	f, _ := Metaphone("f")
	if f != "f" {
		t.Errorf("Metaphone 'f' expected 'f' and got '%s'.", f)
	}
	f, _ = Metaphone("zf")
	if f != "sf" {
		t.Errorf("Metaphone 'zf' expected 'sf' and got '%s'.", f)
	}
	f, _ = Metaphone("zfz")
	if f != "sfs" {
		t.Errorf("Metaphone 'zfz' expected 'sfs' and got '%s'.", f)
	}

	// e
	e, _ := Metaphone("e")
	if e != "e" {
		t.Errorf("Metaphone 'e' expected 'e' and got '%s'.", e)
	}
	e, _ = Metaphone("ze")
	if e != "s" {
		t.Errorf("Metaphone 'ze' expected 's' and got '%s'.", e)
	}

	// d
	d, _ := Metaphone("d")
	if d != "t" {
		t.Errorf("Metaphone 'd' expected 't' and got '%s'.", d)
	}
	d, _ = Metaphone("fudge")
	if d != "fjj" { // php wrongly says fj
		t.Errorf("Metaphone 'fudge' expected 'fjj' and got '%s'.", d)
	}
	d, _ = Metaphone("dodgy") // php wrongly says tj
	if d != "tjj" {           // others wrongly say tjjy
		t.Errorf("Metaphone 'dodgy' expected 'tjj' and got '%s'.", d)
	}
	d, _ = Metaphone("dodgi") // php wrongly says tj
	if d != "tjj" {
		t.Errorf("Metaphone 'dodgi' expected 'tjj' and got '%s'.", d)
	}
	d, _ = Metaphone("zd")
	if d != "st" {
		t.Errorf("Metaphone 'zd' expected 'st' and got '%s'.", d)
	}
	d, _ = Metaphone("zdz")
	if d != "sts" {
		t.Errorf("Metaphone 'zdz' expected 'sts' and got '%s'.", d)
	}

	// c
	c, _ := Metaphone("c")
	if c != "k" {
		t.Errorf("Metaphone 'c' expected 'k' and got '%s'.", c)
	}
	c, _ = Metaphone("zcia")
	if c != "sx" {
		t.Errorf("Metaphone 'zcia' expected 'sx' and got '%s'.", c)
	}
	c, _ = Metaphone("zciaz")
	if c != "sxs" {
		t.Errorf("Metaphone 'zciaz' expected 'sxs' and got '%s'.", c)
	}
	c, _ = Metaphone("zch")
	if c != "sx" {
		t.Errorf("Metaphone 'zch' expected 'sx' and got '%s'.", c)
	}
	c, _ = Metaphone("zchz")
	if c != "sxs" {
		t.Errorf("Metaphone 'zchz' expected 'sxs' and got '%s'.", c)
	}
	c, _ = Metaphone("zci")
	if c != "ss" {
		t.Errorf("Metaphone 'zci' expected 'ss' and got '%s'.", c)
	}
	c, _ = Metaphone("zciz")
	if c != "sss" {
		t.Errorf("Metaphone 'zciz' expected 'sss' and got '%s'.", c)
	}
	c, _ = Metaphone("zce")
	if c != "ss" {
		t.Errorf("Metaphone 'zce' expected 'ss' and got '%s'.", c)
	}
	c, _ = Metaphone("zcez")
	if c != "sss" {
		t.Errorf("Metaphone 'zcez' expected 'sss' and got '%s'.", c)
	}
	c, _ = Metaphone("zcy")
	if c != "ss" {
		t.Errorf("Metaphone 'zcy' expected 'ss' and got '%s'.", c)
	}
	c, _ = Metaphone("zcyz")
	if c != "sss" {
		t.Errorf("Metaphone 'zcyz' expected 'sss' and got '%s'.", c)
	}
	c, _ = Metaphone("zsci")
	if c != "ss" {
		t.Errorf("Metaphone 'zsci' expected 'ss' and got '%s'.", c)
	}
	c, _ = Metaphone("zsciz")
	if c != "sss" {
		t.Errorf("Metaphone 'zsciz' expected 'sss' and got '%s'.", c)
	}
	c, _ = Metaphone("zsce")
	if c != "ss" {
		t.Errorf("Metaphone 'zsce' expected 'ss' and got '%s'.", c)
	}
	c, _ = Metaphone("zscez")
	if c != "sss" {
		t.Errorf("Metaphone 'zscez' expected 'sss' and got '%s'.", c)
	}
	c, _ = Metaphone("zscy")
	if c != "ss" {
		t.Errorf("Metaphone 'zscy' expected 'ss' and got '%s'.", c)
	}
	c, _ = Metaphone("zscyz")
	if c != "sss" {
		t.Errorf("Metaphone 'zscyz' expected 'sss' and got '%s'.", c)
	}
	c, _ = Metaphone("zsch") // php wrongly says ssx
	if c != "sskh" {
		t.Errorf("Metaphone 'zsch' expected 'sskh' and got '%s'.", c)
	}
	c, _ = Metaphone("zc")
	if c != "sk" {
		t.Errorf("Metaphone 'zc' expected 'sk' and got '%s'.", c)
	}
	c, _ = Metaphone("zcz")
	if c != "sks" {
		t.Errorf("Metaphone 'zcz' expected 'sks' and got '%s'.", c)
	}

	// b
	b, _ := Metaphone("b")
	if b != "b" {
		t.Errorf("Metaphone 'b' expected 'b' and got '%s'.", b)
	}
	b, _ = Metaphone("zb")
	if b != "sb" {
		t.Errorf("Metaphone 'zb' expected 'sb' and got '%s'.", b)
	}
	b, _ = Metaphone("zbz")
	if b != "sbs" {
		t.Errorf("Metaphone 'zbz' expected 'sbs' and got '%s'.", b)
	}
	b, _ = Metaphone("zmb")
	if b != "sm" {
		t.Errorf("Metaphone 'zmb' expected 'sm' and got '%s'.", b)
	}

	// a
	a, _ := Metaphone("a")
	if a != "a" {
		t.Errorf("Metaphone 'a' expected 'a' and got '%s'.", a)
	}
	a, _ = Metaphone("za")
	if a != "s" {
		t.Errorf("Metaphone 'za' expected 's' and got '%s'.", a)
	}

	// Miscellaneous
	misc, _ := Metaphone("dumb")
	if misc != "tm" {
		t.Errorf("Metaphone 'dumb' expected 'tm' and got '%s'.", misc)
	}
	misc, _ = Metaphone("smith")
	if misc != "sm0" {
		t.Errorf("Metaphone 'smith' expected 'sm0' and got '%s'.", misc)
	}
	misc, _ = Metaphone("school") // php wrongly says sxl
	if misc != "skhl" {
		t.Errorf("Metaphone 'school' expected 'skhl' and got '%s'.", misc)
	}
	misc, _ = Metaphone("merci")
	if misc != "mrs" {
		t.Errorf("Metaphone 'merci' expected 'mrs' and got '%s'.", misc)
	}
	misc, _ = Metaphone("cool")
	if misc != "kl" {
		t.Errorf("Metaphone 'cool' expected 'kl' and got '%s'.", misc)
	}
	misc, _ = Metaphone("aebersold")
	if misc != "ebrslt" {
		t.Errorf("Metaphone 'aebersold' expected 'ebrslt' and got '%s'.", misc)
	}
	misc, _ = Metaphone("gnagy")
	if misc != "nj" {
		t.Errorf("Metaphone 'gnagy' expected 'nj' and got '%s'.", misc)
	}
	misc, _ = Metaphone("knuth")
	if misc != "n0" {
		t.Errorf("Metaphone 'knuth' expected 'n0' and got '%s'.", misc)
	}
	misc, _ = Metaphone("pniewski")
	if misc != "nsk" {
		t.Errorf("Metaphone 'pniewski' expected 'nsk' and got '%s'.", misc)
	}
	misc, _ = Metaphone("wright") // php wrongly says rft
	if misc != "rht" {
		t.Errorf("Metaphone 'wright' expected 'rht' and got '%s'.", misc)
	}
	misc, _ = Metaphone("phone")
	if misc != "fn" {
		t.Errorf("Metaphone 'phone' expected 'fn' and got '%s'.", misc)
	}
	misc, _ = Metaphone("aggregate")
	if misc != "akrkt" {
		t.Errorf("Metaphone 'aggregate' expected 'akrkt' and got '%s'.", misc)
	}
	misc, _ = Metaphone("accuracy")
	if misc != "akkrs" {
		t.Errorf("Metaphone 'accuracy' expected 'akkrs' and got '%s'.", misc)
	}
	misc, _ = Metaphone("encyclopedia")
	if misc != "ensklpt" {
		t.Errorf("Metaphone 'encyclopedia' expected 'ensklpt' and got '%s'.", misc)
	}
	misc, _ = Metaphone("honorificabilitudinitatibus")
	if misc != "hnrfkblttnttbs" {
		t.Errorf("Metaphone 'honorificabilitudinitatibus' expected "+
			"'hnrfkblttnttbs' and got '%s'.", misc)
	}
	misc, _ = Metaphone("antidisestablishmentarianism")
	if misc != "anttsstblxmntrnsm" {
		t.Errorf("Metaphone 'antidisestablishmentarianism' expected "+
			"'anttsstblxmntrnsm' and got '%s'.", misc)
	}
}
func ExampleMetaphone() {
	result, err := Metaphone("Colorado")
	if err != nil {
		fmt.Print("an unexpected error was encountered")
	}
	fmt.Print(result)
	// Output: klrt
}
