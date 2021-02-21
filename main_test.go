package main

import (
	"reflect"
	"testing"
)

func TestGetStringToFloat64(t *testing.T) {
	list := []string{"0.1", "0.2", "0.3", "0.4", "0.5", "0.6", "0.7", "0.8", "0.9"}
	expect := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}
	actual := stringToFloat64(list)

	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("actual: %v, expected: %v", actual, expect)
	}
}

func TestGetMaximum(t *testing.T) {
	list := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}
	expect := 0.9
	actual := getMaximum(list)

	if actual != expect {
		t.Errorf("actual: %v, expected: %v", actual, expect)
	}
}

func TestGetMinimum(t *testing.T) {
	list := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}
	expect := 0.1
	actual := getMinimum(list)

	if actual != expect {
		t.Errorf("actual: %v, expected: %v", actual, expect)
	}
}

func TestGetAverage(t *testing.T) {
	list := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}
	expect := 0.5
	actual := getAverage(list)

	if actual != expect {
		t.Errorf("actual: %v, expected: %v", actual, expect)
	}
}

func TestGetMedian(t *testing.T) {
	list := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}
	expect := 0.4
	actual := getMedian(list)

	if actual != expect {
		t.Errorf("actual: %v, expected: %v", actual, expect)
	}
}

func TestGetPercentile(t *testing.T) {
	list := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}
	expect := 0.7
	actual := getPercentile(list, 80)

	if actual != expect {
		t.Errorf("actual: %v, expected: %v", actual, expect)
	}
}

func TestGetPercentileWithZero(t *testing.T) {
	list := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}
	expect := 0.1
	actual := getPercentile(list, 0)

	if actual != expect {
		t.Errorf("actual: %v, expected: %v", actual, expect)
	}
}
