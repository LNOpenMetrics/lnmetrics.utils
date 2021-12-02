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

func TestIterateThrogh(t *testing.T) {
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
