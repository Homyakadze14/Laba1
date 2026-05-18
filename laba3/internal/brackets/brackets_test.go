package brackets

import "testing"

func TestGoodCheck(t *testing.T) {
	ch := NewChecker()
	_, ok := ch.Check("()[]{}")
	if !ok {
		t.Errorf("want true, got %v", ok)
	}
}

func TestBadCheck(t *testing.T) {
	ch := NewChecker()
	_, ok := ch.Check("([)]")
	if ok {
		t.Errorf("want false, got %v", ok)
	}
}

func TestNotEmptyCheck(t *testing.T) {
	ch := NewChecker()
	_, ok := ch.Check("([)][")
	if ok {
		t.Errorf("want false, got %v", ok)
	}
}

func TestGoodArrowCheck(t *testing.T) {
	ch := NewChecker()
	_, ok := ch.Check("([]<>)")
	if !ok {
		t.Errorf("want false, got %v", ok)
	}
}

func TestBadArrowCheck(t *testing.T) {
	ch := NewChecker()
	_, ok := ch.Check("([]<<)>>")
	if ok {
		t.Errorf("want false, got %v", ok)
	}
}

func TestIdxGoodCheck(t *testing.T) {
	ch := NewChecker()
	idx, ok := ch.Check("()[]{}")
	if !ok && idx == -1 {
		t.Errorf("want true and -1, got %v and %v", ok, idx)
	}
}

func TestIdxBadCheck(t *testing.T) {
	ch := NewChecker()
	idx, ok := ch.Check("([)]")
	if ok && idx == 2 {
		t.Errorf("want false and 2, got %v and %v", ok, idx)
	}
}

func TestGoodCheckAndCount(t *testing.T) {
	ch := NewChecker()
	c, ok := ch.CheckAndCount("()[]{}")
	if !ok && c == 3 {
		t.Errorf("want true and 3, got %v and %v", ok, c)
	}
}

func TestIdxBadCheckAndCount(t *testing.T) {
	ch := NewChecker()
	c, ok := ch.CheckAndCount("([)]")
	if ok && c == 0 {
		t.Errorf("want false and 0, got %v and %v", ok, c)
	}
}

func TestGoodCheckOnly(t *testing.T) {
	ch := NewChecker()
	ok := ch.CheckOnly("(()[[]{}})")
	if !ok {
		t.Errorf("want true, got %v", ok)
	}
}

func TestBadCheckOnly(t *testing.T) {
	ch := NewChecker()
	ok := ch.CheckOnly("([)[)]")
	if ok {
		t.Errorf("want false, got %v", ok)
	}
}

func TestGoodCheckLine(t *testing.T) {
	ch := NewChecker()
	_, ok := ch.IgnoreCheck("a + (b * c)", "<>")
	if !ok {
		t.Errorf("want true, got %v", ok)
	}
}

func TestGood2CheckLine(t *testing.T) {
	ch := NewChecker()
	_, ok := ch.IgnoreCheck("if (a[2] > b)", "<>")
	if !ok {
		t.Errorf("want true, got %v", ok)
	}
}

func TestBadCheckLine(t *testing.T) {
	ch := NewChecker()
	_, ok := ch.IgnoreCheck("(a + (b * c)", "<>")
	if ok {
		t.Errorf("want false, got %v", ok)
	}
}

func TestBad2CheckLine(t *testing.T) {
	ch := NewChecker()
	_, ok := ch.IgnoreCheck("if (a[2] > b) else (a[2] < c", "<>")
	if ok {
		t.Errorf("want false, got %v", ok)
	}
}

func TestBiggerCountBrackets(t *testing.T) {
	ch := NewChecker()
	c := ch.CountBrackets("(a + ((b * c)")
	if c < 0 {
		t.Errorf("want count > 0, got %v", c)
	}
	t.Log("Открывающих больше")
}

func TestLowerCountBrackets(t *testing.T) {
	ch := NewChecker()
	c := ch.CountBrackets("(a + (b * c))")
	if c > 0 {
		t.Errorf("want count < 0, got %v", c)
	}
	t.Log("Закрывающих больше")
}

func TestEqualCountBrackets(t *testing.T) {
	ch := NewChecker()
	c := ch.CountBrackets("a + (b * c)")
	if c != 0 {
		t.Errorf("want count = 0, got %v", c)
	}
	t.Log("Равно скобок")
}

func TestCheckAndCountDepth1(t *testing.T) {
	ch := NewChecker()
	d, ok := ch.CheckAndCountDepth("(((()))")
	if ok && d != 4 {
		t.Errorf("want depth = 4, got %v", d)
	}
}

func TestCheckAndCountDepth2(t *testing.T) {
	ch := NewChecker()
	d, ok := ch.CheckAndCountDepth("(((([][](()))))")
	if ok && d != 6 {
		t.Errorf("want depth = 6, got %v", d)
	}
}

func TestCheckWithError1(t *testing.T) {
	ch := NewChecker()
	ok, err := ch.CheckWithError("())")
	if ok {
		t.Errorf("want ok = false, got %v", ok)
	}
	t.Log(err)
}

func TestCheckWithError2(t *testing.T) {
	ch := NewChecker()
	ok, err := ch.CheckWithError("(()")
	if ok {
		t.Errorf("want ok = false, got %v", ok)
	}
	t.Log(err)
}

func TestCheckWithError3(t *testing.T) {
	ch := NewChecker()
	ok, err := ch.CheckWithError("(()]")
	if ok {
		t.Errorf("want ok = false, got %v", ok)
	}
	t.Log(err)
}

func TestCheckWithPrintStack(t *testing.T) {
	ch := NewChecker()
	_, ok := ch.CheckWithPrintStack("(([]))", t.Log)
	if !ok {
		t.Errorf("want ok = true, got %v", ok)
	}
}

func TestGoodMultiplyCheck(t *testing.T) {
	ch := NewChecker()
	res := ch.MultiplyCheck([]string{"()[]{}", "{(<>)}"})
	if !res[0] && !res[1] {
		t.Errorf("want true and true, got %v and %v", res[0], res[1])
	}
}

func TestGoodMultiplyCheck1(t *testing.T) {
	ch := NewChecker()
	res := ch.MultiplyCheck([]string{"()", "([)]", "{[]}"})
	if !res[0] && res[1] && !res[2] {
		t.Errorf("want true, false, true, got %v, %v, %v", res[0], res[1], res[2])
	}
}
