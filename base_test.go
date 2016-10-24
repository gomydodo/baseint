package baseint

import "testing"

func TestNormalConvert(t *testing.T) {
	var arr = []uint64{1, 200, 300, 18731, 3843483491, 1118391}

	for _, a := range arr {
		s := Itoa(a)
		i, err := Atoi(s)
		if err != nil {
			t.Fatal("Atoi should not be error:", err)
		}

		if i != uint64(a) {
			t.Fatal("Atoi's result is not correct")
		}
	}

	_, err := Atoi("fwir")
	if err != nil {
		t.Fatal("Atoi should not be error:", err)
	}

	_, err = Atoi("xxxf$")
	if err != ErrChar {
		t.Fatal("Atoi with error char should return ErrChar error")
	}

}

func TestSetBaseString(t *testing.T) {
	var i uint64 = 2
	s1 := Itoa(i)
	err := SetBaseString("1234567890")
	if err != nil {
		t.Fatal("SetBaseString should not return err")
	}

	s2 := Itoa(i)

	if s1 == s2 {
		t.Fatal("s1 should not equal to s2", s1, s2)
	}

	err = SetBaseString("1239992212")
	if err != ErrRepeatChar {
		t.Fatal("should return repeat char error")
	}

}
