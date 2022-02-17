package binary

/*
var intComparator primitives.Comparator[int] = func(left, right int) value.Compared {
	if left < right {
		return value.Less
	}
	if left > right {
		return value.Larger
	}
	return value.Equal
}

func TestCase1(t *testing.T) {
	expected := 16

	data := []int{expected, 11, 9, 10, 5, 6, 8, 1, 2, 4}
	binaryHeap := Init(data, intComparator)

	actual := binaryHeap.Peek()

	if actual != expected {
		t.Errorf("Actual: %d, expected: %d.", actual, expected)
	}
}

func TestCase2(t *testing.T) {
	type composite struct {
		intValue    int
		boolValue   bool
		stringValue string
	}

	var comparator primitives.Comparator[composite] = func(left, right composite) value.Compared {
		leftHash, err := hashstructure.Hash(left, hashstructure.FormatV2, nil)
		if err != nil {
			t.Error(err)
		}
		rightHash, err := hashstructure.Hash(right, hashstructure.FormatV2, nil)
		if err != nil {
			t.Error(err)
		}

		if leftHash < rightHash {
			return value.Less
		}
		if leftHash > rightHash {
			return value.Larger
		}
		return value.Equal
	}

	expected := composite{
		intValue:    1,
		boolValue:   true,
		stringValue: "test1",
	}

	data := []composite{
		expected,
		{
			intValue:    2,
			boolValue:   false,
			stringValue: "test2",
		},
		{
			intValue:    3,
			boolValue:   true,
			stringValue: "test3",
		},
	}

	binaryHeap := Init(data, comparator)

	actual := binaryHeap.Peek()

	if actual != expected {
		t.Errorf("Actual: %d, expected: %d.", actual.intValue, expected.intValue)
	}
}
*/
