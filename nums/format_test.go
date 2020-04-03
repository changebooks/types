package nums

import "testing"

func TestFormat(t *testing.T) {
	isNumeric, isFloat, isNegative, isBase16, isScientific := Format("")
	if isNumeric || isFloat || isNegative || isBase16 || isScientific {
		t.Errorf("%t %t %t %t %t", isNumeric, isFloat, isNegative, isBase16, isScientific)
	}

	isNumeric2, isFloat2, isNegative2, isBase162, isScientific2 := Format("a")
	if isNumeric2 || isFloat2 || isNegative2 || isBase162 || isScientific2 {
		t.Errorf("%t %t %t %t %t", isNumeric2, isFloat2, isNegative2, isBase162, isScientific2)
	}

	isNumeric3, isFloat3, isNegative3, isBase163, isScientific3 := Format("1")
	if !isNumeric3 || isFloat3 || isNegative3 || isBase163 || isScientific3 {
		t.Errorf("%t %t %t %t %t", isNumeric3, isFloat3, isNegative3, isBase163, isScientific3)
	}

	isNumeric4, isFloat4, isNegative4, isBase164, isScientific4 := Format("-1")
	if !isNumeric4 || isFloat4 || !isNegative4 || isBase164 || isScientific4 {
		t.Errorf("%t %t %t %t %t", isNumeric4, isFloat4, isNegative4, isBase164, isScientific4)
	}

	isNumeric5, isFloat5, isNegative5, isBase165, isScientific5 := Format("0.5")
	if !isNumeric5 || !isFloat5 || isNegative5 || isBase165 || isScientific5 {
		t.Errorf("%t %t %t %t %t", isNumeric5, isFloat5, isNegative5, isBase165, isScientific5)
	}

	isNumeric6, isFloat6, isNegative6, isBase166, isScientific6 := Format("0e1")
	if !isNumeric6 || isFloat6 || isNegative6 || isBase166 || !isScientific6 {
		t.Errorf("%t %t %t %t %t", isNumeric6, isFloat6, isNegative6, isBase166, isScientific6)
	}

	isNumeric7, isFloat7, isNegative7, isBase167, isScientific7 := Format("0x1")
	if !isNumeric7 || isFloat7 || isNegative7 || !isBase167 || isScientific7 {
		t.Errorf("%t %t %t %t %t", isNumeric7, isFloat7, isNegative7, isBase167, isScientific7)
	}
}
