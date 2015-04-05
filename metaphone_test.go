package gofuzz

import (
	"testing"
)

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
		t.Errorf("Metaphone failed to evaluate 'z'.")
	}
	z, _ = Metaphone("zz")
	if z != "s" {
		t.Errorf("Metaphone failed to evaluate 'zz'.")
	}

	// y
	y, _ := Metaphone("y")
	if y != "" {
		t.Errorf("Metaphone failed to evaluate 'y'.")
	}
	y, _ = Metaphone("zy")
	if y != "s" {
		t.Errorf("Metaphone failed to evaluate 'zy'.")
	}
	y, _ = Metaphone("zyz")
	if y != "ss" {
		t.Errorf("Metaphone failed to evaluate 'zyz'.")
	}
	y, _ = Metaphone("zya")
	if y != "sy" {
		t.Errorf("Metaphone failed to evaluate 'zya'.")
	}

	// x
	x, _ := Metaphone("x")
	if x != "s" {
		t.Errorf("Metaphone failed to evaluate 'x'.")
	}
	x, _ = Metaphone("zx")
	if x != "sks" {
		t.Errorf("Metaphone failed to evaluate 'x'.")
	}
	x, _ = Metaphone("zxz")
	if x != "skss" {
		t.Errorf("Metaphone failed to evaluate 'x'.")
	}

	// w
	w, _ := Metaphone("w")
	if w != "" {
		t.Errorf("Metaphone failed to evaluate 'w'.")
	}
	w, _ = Metaphone("zw")
	if w != "s" {
		t.Errorf("Metaphone failed to evaluate 'zw'.")
	}
	w, _ = Metaphone("zwz")
	if w != "ss" {
		t.Errorf("Metaphone failed to evaluate 'zwz'.")
	}
	w, _ = Metaphone("zwa")
	if w != "sw" {
		t.Errorf("Metaphone failed to evaluate 'zwa'.")
	}

	// v
	v, _ := Metaphone("v")
	if v != "f" {
		t.Errorf("Metaphone failed to evaluate 'v'.")
	}
	v, _ = Metaphone("zv")
	if v != "sf" {
		t.Errorf("Metaphone failed to evaluate 'zv'.")
	}
	v, _ = Metaphone("zvz")
	if v != "sfs" {
		t.Errorf("Metaphone failed to evaluate 'zvz'.")
	}

	// u
	u, _ := Metaphone("u")
	if u != "u" {
		t.Errorf("Metaphone failed to evaluate 'u'.")
	}
	u, _ = Metaphone("zu")
	if u != "s" {
		t.Errorf("Metaphone failed to evaluate 'zu'.")
	}

	// t
	letterT, _ := Metaphone("t")
	if letterT != "t" {
		t.Errorf("Metaphone failed to evaluate 't'.")
	}
	letterT, _ = Metaphone("ztiaz")
	if letterT != "sxs" {
		t.Errorf("Metaphone failed to evaluate 'ztiaz'.")
	}
	letterT, _ = Metaphone("ztioz")
	if letterT != "sxs" {
		t.Errorf("Metaphone failed to evaluate 'ztioz'.")
	}
	letterT, _ = Metaphone("zthz")
	if letterT != "s0s" {
		t.Errorf("Metaphone failed to evaluate 'zthz'.")
	}
	letterT, _ = Metaphone("ztchz")
	if letterT != "sxs" {
		t.Errorf("Metaphone failed to evaluate 'ztchz'.")
	}
	letterT, _ = Metaphone("ztz")
	if letterT != "sts" {
		t.Errorf("Metaphone failed to evaluate 'ztz'.")
	}

	// s
	s, _ := Metaphone("s")
	if s != "s" {
		t.Errorf("Metaphone failed to evaluate 's'.")
	}
	s, _ = Metaphone("zshz")
	if s != "sxs" {
		t.Errorf("Metaphone failed to evaluate 'zshz'.")
	}
	s, _ = Metaphone("zsioz")
	if s != "sxs" {
		t.Errorf("Metaphone failed to evaluate 'zsioz'.")
	}
	s, _ = Metaphone("zsiaz")
	if s != "sxs" {
		t.Errorf("Metaphone failed to evaluate 'zsiaz'.")
	}
	s, _ = Metaphone("zs")
	if s != "ss" {
		t.Errorf("Metaphone failed to evaluate 'zs'.")
	}
	s, _ = Metaphone("zsz")
	if s != "sss" {
		t.Errorf("Metaphone failed to evaluate 'zsz'.")
	}

	// r
	r, _ := Metaphone("r")
	if r != "r" {
		t.Errorf("Metaphone failed to evaluate 'r'.")
	}
	r, _ = Metaphone("zr")
	if r != "sr" {
		t.Errorf("Metaphone failed to evaluate 'zr'.")
	}
	r, _ = Metaphone("zrz")
	if r != "srs" {
		t.Errorf("Metaphone failed to evaluate 'zrz'.")
	}

	// q
	q, _ := Metaphone("q")
	if q != "k" {
		t.Errorf("Metaphone failed to evaluate 'q'.")
	}
	q, _ = Metaphone("zq")
	if q != "sk" {
		t.Errorf("Metaphone failed to evaluate 'zq'.")
	}
	q, _ = Metaphone("zqz")
	if q != "sks" {
		t.Errorf("Metaphone failed to evaluate 'zqz'.")
	}

	// p
	p, _ := Metaphone("p")
	if p != "p" {
		t.Errorf("Metaphone failed to evaluate 'p'.")
	}
	p, _ = Metaphone("zp")
	if p != "sp" {
		t.Errorf("Metaphone failed to evaluate 'zp'.")
	}
	p, _ = Metaphone("zph")
	if p != "sf" {
		t.Errorf("Metaphone failed to evaluate 'zph'.")
	}
	p, _ = Metaphone("zpz")
	if p != "sps" {
		t.Errorf("Metaphone failed to evaluate 'zpz'.")
	}

	// o
	o, _ := Metaphone("o")
	if o != "o" {
		t.Errorf("Metaphone failed to evaluate 'o'.")
	}
	o, _ = Metaphone("zo")
	if o != "s" {
		t.Errorf("Metaphone failed to evaluate 'zo'.")
	}

	// n
	n, _ := Metaphone("n")
	if n != "n" {
		t.Errorf("Metaphone failed to evaluate 'n'.")
	}
	n, _ = Metaphone("zn")
	if n != "sn" {
		t.Errorf("Metaphone failed to evaluate 'zn'.")
	}
	n, _ = Metaphone("znz")
	if n != "sns" {
		t.Errorf("Metaphone failed to evaluate 'znz'.")
	}

	// m
	m, _ := Metaphone("m")
	if m != "m" {
		t.Errorf("Metaphone failed to evaluate 'm'.")
	}
	m, _ = Metaphone("zm")
	if m != "sm" {
		t.Errorf("Metaphone failed to evaluate 'zm'.")
	}
	m, _ = Metaphone("zmz")
	if m != "sms" {
		t.Errorf("Metaphone failed to evaluate 'zmz'.")
	}

	// l
	l, _ := Metaphone("l")
	if l != "l" {
		t.Errorf("Metaphone failed to evaluate 'l'.")
	}
	l, _ = Metaphone("zl")
	if l != "sl" {
		t.Errorf("Metaphone failed to evaluate 'zl'.")
	}
	l, _ = Metaphone("zlz")
	if l != "sls" {
		t.Errorf("Metaphone failed to evaluate 'zlz'.")
	}

	// k
	k, _ := Metaphone("k")
	if k != "k" {
		t.Errorf("Metaphone failed to evaluate 'k'.")
	}
	k, _ = Metaphone("zk")
	if k != "sk" {
		t.Errorf("Metaphone failed to evaluate 'zk'.")
	}
	k, _ = Metaphone("zck")
	if k != "sk" {
		t.Errorf("Metaphone failed to evaluate 'zck'.")
	}

	// j
	j, _ := Metaphone("j")
	if j != "j" {
		t.Errorf("Metaphone failed to evaluate 'j'.")
	}
	j, _ = Metaphone("zj")
	if j != "sj" {
		t.Errorf("Metaphone failed to evaluate 'zj'.")
	}
	j, _ = Metaphone("zjz")
	if j != "sjs" {
		t.Errorf("Metaphone failed to evaluate 'zjz'.")
	}

	// i
	i, _ := Metaphone("i")
	if i != "i" {
		t.Errorf("Metaphone failed to evaluate 'i'.")
	}
	i, _ = Metaphone("zi")
	if i != "s" {
		t.Errorf("Metaphone failed to evaluate 'zi'.")
	}

	// h
	h, _ := Metaphone("h") // php wrongly says nil
	if h != "h" {
		t.Errorf("Metaphone failed to evaluate 'h'.")
	}
	h, _ = Metaphone("zh") // php wrongly says s
	if h != "sh" {
		t.Errorf("Metaphone failed to evaluate 'zh'.")
	}
	h, _ = Metaphone("zah")
	if h != "s" {
		t.Errorf("Metaphone failed to evaluate 'zah'.")
	}
	h, _ = Metaphone("zchh")
	if h != "sx" {
		t.Errorf("Metaphone failed to evaluate 'zchh'.")
	}
	// h, _ = Metaphone("ha")
	// if h != "h" {
	// 	t.Errorf("Metaphone failed to evaluate 'ha'.")
	// }

	// g
	g, _ := Metaphone("g")
	if g != "k" {
		t.Errorf("Metaphone failed to evaluate 'g'.")
	}
	g, _ = Metaphone("zg")
	if g != "sk" {
		t.Errorf("Metaphone failed to evaluate 'zg'.")
	}
	g, _ = Metaphone("zgh") // php wrongly says sf
	if g != "skh" {
		t.Errorf("Metaphone failed to evaluate 'zgh'.")
	}
	g, _ = Metaphone("zghz") // php wrongly says sfs
	if g != "shs" {
		t.Errorf("Metaphone failed to evaluate 'zghz'.")
	}
	// g, _ = Metaphone("zgha") // php wrongly says sf
	// if g != "sh" {           // others wrongly say skh
	// 	t.Errorf("Metaphone failed to evaluate 'zgha'.")
	// }
	g, _ = Metaphone("zgn")
	if g != "sn" {
		t.Errorf("Metaphone failed to evaluate 'zgn'.")
	}
	g, _ = Metaphone("zgns")
	if g != "skns" {
		t.Errorf("Metaphone failed to evaluate 'zgns'.")
	}
	g, _ = Metaphone("zgned") // php wrongly says sknt
	if g != "snt" {
		t.Errorf("Metaphone failed to evaluate 'zgned'.")
	}
	g, _ = Metaphone("zgneds") // php wrongly says snts
	if g != "sknts" {
		t.Errorf("Metaphone failed to evaluate 'zgneds'.")
	}
	g, _ = Metaphone("zgi")
	if g != "sj" {
		t.Errorf("Metaphone failed to evaluate 'zgi'.")
	}
	g, _ = Metaphone("zgiz")
	if g != "sjs" {
		t.Errorf("Metaphone failed to evaluate 'zgiz'.")
	}
	g, _ = Metaphone("zge")
	if g != "sj" {
		t.Errorf("Metaphone failed to evaluate 'zge'.")
	}
	g, _ = Metaphone("zgez")
	if g != "sjs" {
		t.Errorf("Metaphone failed to evaluate 'zgez'.")
	}
	g, _ = Metaphone("zgy")
	if g != "sj" {
		t.Errorf("Metaphone failed to evaluate 'zgy'.")
	}
	g, _ = Metaphone("zgyz")
	if g != "sjs" {
		t.Errorf("Metaphone failed to evaluate 'zgyz'.")
	}
	g, _ = Metaphone("zgz")
	if g != "sks" {
		t.Errorf("Metaphone failed to evaluate 'zgz'.")
	}

	// f
	f, _ := Metaphone("f")
	if f != "f" {
		t.Errorf("Metaphone failed to evaluate 'f'.")
	}
	f, _ = Metaphone("zf")
	if f != "sf" {
		t.Errorf("Metaphone failed to evaluate 'zf'.")
	}
	f, _ = Metaphone("zfz")
	if f != "sfs" {
		t.Errorf("Metaphone failed to evaluate 'zfz'.")
	}

	// e
	e, _ := Metaphone("e")
	if e != "e" {
		t.Errorf("Metaphone failed to evaluate 'e'.")
	}
	e, _ = Metaphone("ze")
	if e != "s" {
		t.Errorf("Metaphone failed to evaluate 'ze'.")
	}

	// d
	d, _ := Metaphone("d")
	if d != "t" {
		t.Errorf("Metaphone failed to evaluate 'd'.")
	}
	d, _ = Metaphone("fudge")
	if d != "fjj" { // php wrongly says fj
		t.Errorf("Metaphone failed to evaluate 'fudge'.")
	}
	d, _ = Metaphone("dodgy") // php wrongly says tj
	if d != "tjj" {           // others wrongly say tjjy
		t.Errorf("Metaphone failed to evaluate 'dodgy'.")
	}
	d, _ = Metaphone("dodgi") // php wrongly says tj
	if d != "tjj" {
		t.Errorf("Metaphone failed to evaluate 'dodgi'.")
	}
	d, _ = Metaphone("zd")
	if d != "st" {
		t.Errorf("Metaphone failed to evaluate 'zd'.")
	}
	d, _ = Metaphone("zdz")
	if d != "sts" {
		t.Errorf("Metaphone failed to evaluate 'zdz'.")
	}

	// c
	c, _ := Metaphone("c")
	if c != "k" {
		t.Errorf("Metaphone failed to evaluate 'c'.")
	}
	c, _ = Metaphone("zcia")
	if c != "sx" {
		t.Errorf("Metaphone failed to evaluate 'zcia'.")
	}
	c, _ = Metaphone("zciaz")
	if c != "sxs" {
		t.Errorf("Metaphone failed to evaluate 'zciaz'.")
	}
	c, _ = Metaphone("zch")
	if c != "sx" {
		t.Errorf("Metaphone failed to evaluate 'zch'.")
	}
	c, _ = Metaphone("zchz")
	if c != "sxs" {
		t.Errorf("Metaphone failed to evaluate 'zchz'.")
	}
	c, _ = Metaphone("zci")
	if c != "ss" {
		t.Errorf("Metaphone failed to evaluate 'zci'.")
	}
	c, _ = Metaphone("zciz")
	if c != "sss" {
		t.Errorf("Metaphone failed to evaluate 'zciz'.")
	}
	c, _ = Metaphone("zce")
	if c != "ss" {
		t.Errorf("Metaphone failed to evaluate 'zce'.")
	}
	c, _ = Metaphone("zcez")
	if c != "sss" {
		t.Errorf("Metaphone failed to evaluate 'zcez'.")
	}
	c, _ = Metaphone("zcy")
	if c != "ss" {
		t.Errorf("Metaphone failed to evaluate 'zcy'.")
	}
	c, _ = Metaphone("zcyz")
	if c != "sss" {
		t.Errorf("Metaphone failed to evaluate 'zcyz'.")
	}
	c, _ = Metaphone("zsci")
	if c != "ss" {
		t.Errorf("Metaphone failed to evaluate 'zsci'.")
	}
	c, _ = Metaphone("zsciz")
	if c != "sss" {
		t.Errorf("Metaphone failed to evaluate 'zsciz'.")
	}
	c, _ = Metaphone("zsce")
	if c != "ss" {
		t.Errorf("Metaphone failed to evaluate 'zsce'.")
	}
	c, _ = Metaphone("zscez")
	if c != "sss" {
		t.Errorf("Metaphone failed to evaluate 'zscez'.")
	}
	c, _ = Metaphone("zscy")
	if c != "ss" {
		t.Errorf("Metaphone failed to evaluate 'zscy'.")
	}
	c, _ = Metaphone("zscyz")
	if c != "sss" {
		t.Errorf("Metaphone failed to evaluate 'zscyz'.")
	}
	c, _ = Metaphone("zsch") // php wrongly says ssx
	if c != "sskh" {
		t.Errorf("Metaphone failed to evaluate 'zsch'.")
	}
	c, _ = Metaphone("zc")
	if c != "sk" {
		t.Errorf("Metaphone failed to evaluate 'zc'.")
	}
	c, _ = Metaphone("zcz")
	if c != "sks" {
		t.Errorf("Metaphone failed to evaluate 'zcz'.")
	}

	// b
	b, _ := Metaphone("b")
	if b != "b" {
		t.Errorf("Metaphone failed to evaluate 'b'.")
	}
	b, _ = Metaphone("zb")
	if b != "sb" {
		t.Errorf("Metaphone failed to evaluate 'zb'.")
	}
	b, _ = Metaphone("zbz")
	if b != "sbs" {
		t.Errorf("Metaphone failed to evaluate 'zbz'.")
	}
	b, _ = Metaphone("zmb")
	if b != "sm" {
		t.Errorf("Metaphone failed to evaluate 'smb'.")
	}

	// a
	a, _ := Metaphone("a")
	if a != "a" {
		t.Errorf("Metaphone failed to evaluate 'a'.")
	}
	a, _ = Metaphone("za")
	if a != "s" {
		t.Errorf("Metaphone failed to evaluate 'za'.")
	}

	// Miscellaneous
	misc, _ := Metaphone("dumb")
	if misc != "tm" {
		t.Errorf("Metaphone failed to evaluate 'dumb'.")
	}
	misc, _ = Metaphone("smith")
	if misc != "sm0" {
		t.Errorf("Metaphone failed to evaluate 'smith'.")
	}
	// misc, _ = Metaphone("school")
	// if misc != "skhl" {
	// 	t.Errorf("Metaphone failed to evaluate 'school'.")
	// }
	misc, _ = Metaphone("merci")
	if misc != "mrs" {
		t.Errorf("Metaphone failed to evaluate 'merci'.")
	}
	// misc, _ = Metaphone("cool")
	// if misc != "kl" {
	// 	t.Errorf("Metaphone failed to evaluate 'cool'.")
	// }
	misc, _ = Metaphone("aebersold")
	if misc != "ebrslt" {
		t.Errorf("Metaphone failed to evaluate 'aebersold'.")
	}
	misc, _ = Metaphone("gnagy")
	if misc != "nj" {
		t.Errorf("Metaphone failed to evaluate 'gnagy'.")
	}
	misc, _ = Metaphone("knuth")
	if misc != "n0" {
		t.Errorf("Metaphone failed to evaluate 'knuth'.")
	}
	misc, _ = Metaphone("pniewski")
	if misc != "nsk" {
		t.Errorf("Metaphone failed to evaluate 'pniewski'.")
	}
	misc, _ = Metaphone("wright")
	if misc != "rht" {
		t.Errorf("Metaphone failed to evaluate 'wright'.")
	}
	misc, _ = Metaphone("phone")
	if misc != "fn" {
		t.Errorf("Metaphone failed to evaluate 'phone'.")
	}
	misc, _ = Metaphone("aggregate")
	if misc != "akrkt" {
		t.Errorf("Metaphone failed to evaluate 'aggregate'.")
	}
	misc, _ = Metaphone("accuracy")
	if misc != "akkrs" {
		t.Errorf("Metaphone failed to evaluate 'accuracy'.")
	}
	misc, _ = Metaphone("encyclopedia")
	if misc != "ensklpt" {
		t.Errorf("Metaphone failed to evaluate 'encyclopedia'.")
	}
	// misc, _ = Metaphone("honorificabilitudinitatibus")
	// if misc != "hnrfkblttnttbs" {
	// 	t.Errorf("Metaphone failed to evaluate 'honorificabilitudinitatibus'.")
	// }
	misc, _ = Metaphone("antidisestablishmentarianism")
	if misc != "anttsstblxmntrnsm" {
		t.Errorf("Metaphone failed to evaluate 'antidisestablishmentarianism'.")
	}
}
