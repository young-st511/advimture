package game

import "testing"

func TestCalcGrade_S(t *testing.T) {
	// effKeys <= optimalKeys
	if g := CalcGrade(8, 8); g != "S" {
		t.Errorf("expected S, got %s", g)
	}
	if g := CalcGrade(1, 8); g != "S" {
		t.Errorf("expected S, got %s", g)
	}
}

func TestCalcGrade_A(t *testing.T) {
	// effKeys <= optimalKeys * 1.5 (opt=8 → ≤12)
	if g := CalcGrade(9, 8); g != "A" {
		t.Errorf("expected A, got %s (9 vs opt=8)", g)
	}
	if g := CalcGrade(12, 8); g != "A" {
		t.Errorf("expected A, got %s (12 vs opt=8)", g)
	}
}

func TestCalcGrade_B(t *testing.T) {
	// effKeys <= optimalKeys * 2.5 (opt=8 → ≤20)
	if g := CalcGrade(13, 8); g != "B" {
		t.Errorf("expected B, got %s (13 vs opt=8)", g)
	}
	if g := CalcGrade(20, 8); g != "B" {
		t.Errorf("expected B, got %s (20 vs opt=8)", g)
	}
}

func TestCalcGrade_C(t *testing.T) {
	// effKeys > optimalKeys * 2.5 (opt=8 → >20)
	if g := CalcGrade(21, 8); g != "C" {
		t.Errorf("expected C, got %s (21 vs opt=8)", g)
	}
	if g := CalcGrade(100, 8); g != "C" {
		t.Errorf("expected C, got %s (100 vs opt=8)", g)
	}
}

func TestCalcGrade_ZeroOptimal(t *testing.T) {
	if g := CalcGrade(5, 0); g != "C" {
		t.Errorf("expected C for zero optimal, got %s", g)
	}
}

func TestCalcAccuracy_Full(t *testing.T) {
	if a := CalcAccuracy(10, 10); a != 100.0 {
		t.Errorf("expected 100.0, got %f", a)
	}
}

func TestCalcAccuracy_Half(t *testing.T) {
	if a := CalcAccuracy(5, 10); a != 50.0 {
		t.Errorf("expected 50.0, got %f", a)
	}
}

func TestCalcAccuracy_Zero(t *testing.T) {
	if a := CalcAccuracy(0, 0); a != 100.0 {
		t.Errorf("expected 100.0 for zero total, got %f", a)
	}
}

func TestCalcAccuracy_Capped(t *testing.T) {
	// effKeys > totalKeys edge case → capped at 100
	if a := CalcAccuracy(15, 10); a != 100.0 {
		t.Errorf("expected 100.0 (capped), got %f", a)
	}
}

func TestIsBetterGrade(t *testing.T) {
	if !IsBetterGrade("S", "A") {
		t.Error("S should be better than A")
	}
	if !IsBetterGrade("A", "B") {
		t.Error("A should be better than B")
	}
	if IsBetterGrade("B", "A") {
		t.Error("B should not be better than A")
	}
	if IsBetterGrade("C", "C") {
		t.Error("same grade should not be considered better")
	}
}

func TestMentorMessage(t *testing.T) {
	grades := []string{"S", "A", "B", "C"}
	for _, g := range grades {
		msg := MentorMessage(g)
		if msg == "" {
			t.Errorf("MentorMessage(%s) returned empty string", g)
		}
	}
}
