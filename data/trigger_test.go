package data_test

import (
	"os"
	"testing"

	"github.com/danesparza/fxtrigger/data"
)

func TestFile_AddFile_ValidFile_Successful(t *testing.T) {

	//	Arrange
	systemdb := getTestFiles()

	db, err := data.NewManager(systemdb)
	if err != nil {
		t.Fatalf("NewManager failed: %s", err)
	}
	defer func() {
		db.Close()
		os.RemoveAll(systemdb)
	}()

	testTrigger := data.Trigger{Name: "Unit test trigger", Description: "Unit test trigger desc", GPIOPin: "23", WebHooks: []data.WebHook{}}

	//	Act
	newFile, err := db.AddTrigger(testTrigger.Name, testTrigger.Description, testTrigger.GPIOPin, testTrigger.WebHooks, testTrigger.MinimumSleepBeforeRetrigger)

	//	Assert
	if err != nil {
		t.Errorf("AddTrigger - Should add trigger without error, but got: %s", err)
	}

	if newFile.Created.IsZero() {
		t.Errorf("AddTrigger failed: Should have set an item with the correct datetime: %+v", newFile)
	}

}

func TestFile_GetTrigger_ValidTrigger_Successful(t *testing.T) {

	//	Arrange
	systemdb := getTestFiles()

	db, err := data.NewManager(systemdb)
	if err != nil {
		t.Fatalf("NewManager failed: %s", err)
	}
	defer func() {
		db.Close()
		os.RemoveAll(systemdb)
	}()

	testTrigger1 := data.Trigger{Name: "Trigger 1", Description: "Unit test 1", GPIOPin: "11"}
	testTrigger2 := data.Trigger{Name: "Trigger 2", Description: "Unit test 2", GPIOPin: "12"}
	testTrigger3 := data.Trigger{Name: "Trigger 3", Description: "Unit test 3", GPIOPin: "13"}

	//	Act
	db.AddTrigger(testTrigger1.Name, testTrigger1.Description, testTrigger1.GPIOPin, testTrigger1.WebHooks, testTrigger1.MinimumSleepBeforeRetrigger)
	newTrigger2, _ := db.AddTrigger(testTrigger2.Name, testTrigger2.Description, testTrigger2.GPIOPin, testTrigger2.WebHooks, testTrigger2.MinimumSleepBeforeRetrigger)
	db.AddTrigger(testTrigger3.Name, testTrigger3.Description, testTrigger3.GPIOPin, testTrigger3.WebHooks, testTrigger3.MinimumSleepBeforeRetrigger)

	gotTrigger, err := db.GetTrigger(newTrigger2.ID)

	//	Log the file details:
	t.Logf("Trigger: %+v", gotTrigger)

	//	Assert
	if err != nil {
		t.Errorf("GetTrigger - Should get trigger without error, but got: %s", err)
	}

	if len(gotTrigger.ID) < 2 {
		t.Errorf("GetTrigger failed: Should get valid id but got: %v", gotTrigger.ID)
	}
}
