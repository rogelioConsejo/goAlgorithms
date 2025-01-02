package calculation

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNewCalculation(t *testing.T) {
	t.Run("an empty calculation should return a root with a nilNode", func(t *testing.T) {
		calc := NewCalculation("")
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()
		if _, ok := root.(nilNode); !ok {
			t.Error("Expected a nilNode, but got something else")
		}
	})
	t.Run("a calculation with a single numeric value should return a root with a numeric value Node", func(t *testing.T) {
		value := rand.Int()
		calc := NewCalculation(fmt.Sprintf("%d", value))
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()
		if root == nil {
			t.Error("Expected a value, but got nil")
		}
		if root.Value() != value {
			t.Errorf("Expected %d, but got %d", value, root.Value())
		}
	})
	t.Run("a calculation with two values and a plus sign should return a root with a plus Node", func(t *testing.T) {
		randValue1 := rand.Int()
		randValue2 := rand.Int()
		calc := NewCalculation(fmt.Sprintf("%d+%d", randValue1, randValue2))
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()
		if root == nil {
			t.Error("Expected a value, but got nil")
		}
		expected := randValue1 + randValue2
		if root.Value() != expected {
			t.Errorf("Expected %d, but got %d", expected, root.Value())
		}
	})
	t.Run("a calculation with two values and a minus sign should return a root with a minus Node", func(t *testing.T) {
		randValue1 := rand.Int()
		randValue2 := rand.Int()
		calc := NewCalculation(fmt.Sprintf("%d-%d", randValue1, randValue2))
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()

		if root == nil {
			t.Fatal("Expected a value, but got nil")
		}
		expected := randValue1 - randValue2
		if root.Value() != expected {
			t.Errorf("Expected %d, but got %d", expected, root.Value())
		}
	})
	t.Run("a calculation with a multiplication operation should return a root with a multiplication Node", func(t *testing.T) {
		randValue1 := rand.Int()
		randValue2 := rand.Int()
		calc := NewCalculation(fmt.Sprintf("%d*%d", randValue1, randValue2))
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()
		if root == nil {
			t.Error("Expected a value, but got nil")
		}
		expected := randValue1 * randValue2
		if root.Value() != expected {
			t.Errorf("Expected %d, but got %d", expected, root.Value())
		}
	})
	t.Run("a calculation with a division operation should return a root with a division Node", func(t *testing.T) {
		randValue1 := rand.Int()
		randValue2 := rand.Int()
		calc := NewCalculation(fmt.Sprintf("%d/%d", randValue1, randValue2))
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()
		if root == nil {
			t.Error("Expected a value, but got nil")
		}
		expected := randValue1 / randValue2
		if root.Value() != expected {
			t.Errorf("Expected %d, but got %d", expected, root.Value())
		}
	})
	t.Run("44*60+57 should return 2697", func(t *testing.T) {
		calc := NewCalculation("44*60+57")
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()
		if root == nil {
			t.Error("Expected a value, but got nil")
		}
		expected := 2697
		if root.Value() != expected {
			t.Errorf("Expected %d, but got %d", expected, root.Value())
		}
	})
	t.Run("10+20*30 should return 610", func(t *testing.T) {
		calc := NewCalculation("10+20*30")
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()
		if root == nil {
			t.Error("Expected a value, but got nil")
		}
		expected := 610
		if root.Value() != expected {
			t.Errorf("Expected %d, but got %d", expected, root.Value())
		}
	})
	t.Run("10*20+30 should return 230", func(t *testing.T) {
		calc := NewCalculation("10*20+30")
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()
		if root == nil {
			t.Error("Expected a value, but got nil")
		}
		expected := 230
		if root.Value() != expected {
			t.Errorf("Expected %d, but got %d", expected, root.Value())
		}
	})
	t.Run("10*20+30*40 should return 1400", func(t *testing.T) {
		calc := NewCalculation("10*20+30*40")
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()
		if root == nil {
			t.Error("Expected a value, but got nil")
		}
		expected := 1400
		if root.Value() != expected {
			t.Errorf("Expected %d, but got %d", expected, root.Value())
		}
	})
	t.Run("10*20+30*40+50 should return 1450", func(t *testing.T) {
		calc := NewCalculation("10*20+30*40+50")
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()
		if root == nil {
			t.Error("Expected a value, but got nil")
		}
		expected := 1450
		if root.Value() != expected {
			t.Errorf("Expected %d, but got %d", expected, root.Value())
		}
	})
	t.Run("100+10/2 should return 105", func(t *testing.T) {
		calc := NewCalculation("100+10/2")
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()
		if root == nil {
			t.Error("Expected a value, but got nil")
		}
		expected := 105
		if root.Value() != expected {
			t.Errorf("Expected %d, but got %d", expected, root.Value())
		}
	})
	t.Run("100/10+2 should return 12", func(t *testing.T) {
		calc := NewCalculation("100/10+2")
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()
		if root == nil {
			t.Error("Expected a value, but got nil")
		}
		expected := 12
		if root.Value() != expected {
			t.Errorf("Expected %d, but got %d", expected, root.Value())
		}
	})
	t.Run("100/10+2*3 should return 16", func(t *testing.T) {
		calc := NewCalculation("100/10+2*3")
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()
		if root == nil {
			t.Error("Expected a value, but got nil")
		}
		expected := 16
		if root.Value() != expected {
			t.Errorf("Expected %d, but got %d", expected, root.Value())
		}
	})
	t.Run("(10+1)*2 should return 22", func(t *testing.T) {
		calc := NewCalculation("(10+1)*2")
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()
		if root == nil {
			t.Error("Expected a value, but got nil")
		}
		expected := 22
		if root.Value() != expected {
			t.Errorf("Expected %d, but got %d", expected, root.Value())
		}
	})
	t.Run("(10+1)*(2+3) should return 55", func(t *testing.T) {
		calc := NewCalculation("(10+1)*(2+3)")
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()
		if root == nil {
			t.Error("Expected a value, but got nil")
		}
		expected := 55
		if root.Value() != expected {
			t.Errorf("Expected %d, but got %d", expected, root.Value())
		}
	})
	t.Run("((10+1)-(2*3+1)) should return 4", func(t *testing.T) {
		calc := NewCalculation("((10+1)-(2*3+1))")
		if calc == nil {
			t.Error("NewCalculation should not return nil")
		}
		root := calc.Root()
		if root == nil {
			t.Error("Expected a value, but got nil")
		}
		expected := 4
		if root.Value() != expected {
			t.Errorf("Expected %d, but got %d", expected, root.Value())
		}
		fmt.Printf("%+v\n", calc)
	})
}

func TestParseAndCalculate(t *testing.T) {
	t.Run("(100-10)*2+5*(10+5) should return 255", func(t *testing.T) {
		result := ParseAndCalculate("(100-10)*2+5*(10+5)")
		expected := 255
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})
}
