package skace

import "testing"

func TestIsPascal(t *testing.T) {
	sm := map[string]bool{
		"HelloWorld": true,
		"hEllo":      false,
	}

	for k, v := range sm {
		r := IsPascal(k)
		if r != v {
			t.Errorf("%v => expected: %v | got: %v", k, v, r)
		}
	}
}
func TestIsCamel(t *testing.T) {
	sm := map[string]bool{
		"helloWorld":  true,
		"HEllo":       false,
		"HEll_o":      false,
		"hello_World": false,
	}

	for k, v := range sm {
		r := IsCamel(k)
		if r != v {
			t.Errorf("%v => expected: %v | got: %v", k, v, r)
		}
	}
}

func TestIsKedab(t *testing.T) {
	sm := map[string]bool{
		"hello-World": true,
		"HEllo":       false,
		"HEll_o":      false,
	}

	for k, v := range sm {
		r := Iskedab(k)
		if r != v {
			t.Errorf("%v => expected: %v | got: %v", k, v, r)
		}
	}
}

func TestIsSnake(t *testing.T) {
	sm := map[string]bool{
		"hello_World": true,
		"hello":       false,
		"HEllo":       false,
		"HEll_o":      true,
	}

	for k, v := range sm {
		r := IsSnake(k)
		if r != v {
			t.Errorf("%v => expected: %v | got: %v", k, v, r)
		}
	}
}
func TestToCamel(t *testing.T) {
	sm := map[string]string{
		"HEll_o":     "hellO",
		"HelloWorld": "helloWorld",
		"hEllo":      "hEllo",
	}

	for k, v := range sm {
		r := ToCamel(k)
		if r != v {
			t.Errorf("%v => expected: %v | got: %v", k, v, r)
		}
	}
}
func TestToSnake(t *testing.T) {
	sm := map[string]string{
		"HelloWorld": "hello_world",
		"hEllo":      "h_ello",
		"HEll_o":     "hell_o",
	}

	for k, v := range sm {
		r := ToSnake(k)
		if r != v {
			t.Errorf("%v => expected: %v | got: %v", k, v, r)
		}
	}
}
