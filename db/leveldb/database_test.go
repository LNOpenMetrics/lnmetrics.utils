package db

import (
	"os"
	"testing"
)

func init() {
	rootDir, _ := os.Getwd()
	if err := GetInstance().InitDB(rootDir); err != nil {
		panic(err)
	}
}

// Teardown method
func cleanupDB(tb testing.TB) func(tb testing.TB) {
	return func(tb testing.TB) {
		if err := GetInstance().EraseDatabase(); err != nil {
			panic(err)
		}
	}
}

func TestIterateThrogh(t *testing.T) {
	defer cleanupDB(t)

	items := []string{"1", "2", "3", "4", "5", "6"}
	for _, item := range items {
		if err := GetInstance().PutValue(item, item); err != nil {
			t.Errorf("%s", err)
		}
	}

	getItems := make([]string, 0)
	err := GetInstance().IterateThrough("1", "7", func(value string) error {
		getItems = append(getItems, value)
		return nil
	})

	if err != nil {
		t.Errorf("%s", err)
	}

	if len(items) != len(getItems) {
		t.Errorf("Expected size is %d bu received %d", len(items), len(getItems))
	}

	for index := range items {
		if items[index] != getItems[index] {
			t.Errorf("Expected %s at %d but received %s", items[index], index, getItems[index])
		}
	}
}

func TestIterateThroghWithNoFoundVal(t *testing.T) {
	defer cleanupDB(t)

	items := []string{"alibaba/1/metric", "zalibas/2/metric", "tedgs/3/metric", "bibbo/4/metric", "citto/5/metric", "alibaba/2/metric"}
	for _, item := range items {
		if err := GetInstance().PutValue(item, item); err != nil {
			t.Errorf("%s", err)
		}
	}

	getItems := make([]string, 0)
	err := GetInstance().IterateThrough("alibaba/0/metric", "alibaba/7/metric", func(value string) error {
		getItems = append(getItems, value)
		return nil
	})

	if err != nil {
		t.Errorf("%s", err)
	}

	if len(getItems) != 2 {
		t.Errorf("Expected size is %d but received %d", 2, len(getItems))
	}

	for _, value := range getItems {
		switch value {
		case "alibaba/1/metric", "alibaba/2/metric":
			continue
		default:
			t.Errorf("Found %s not expected value", value)
		}
	}
}
