package comment

import (
	"context"
	"testing"
)

// Test add a comment
func testAddComment(t *testing.T) {
	testForms := []CForm{
		// Form valid
		{PostID: 1, Content: "message"},

		// Form not valid
		{PostID: 0, Content: ""},
	}

	for i, form := range testForms {
		_, err := testService.addComment(context.Background(), form)

		switch i {
		case 0:
			if err != nil {
				t.Error("Error should be nil but got:", err)
			}
			break
		case 1:
			if err != ErrorForm {
				t.Errorf("Error should be %s but got: %s", ErrorForm, err)
			}
			break
		}
	}

}

// Test update a comment
func testUpdateComment(t *testing.T) {

	testForms := []CForm{
		// Form valid
		{PostID: 1, Content: "updated message"},

		// Form not valid
		{PostID: 0, Content: ""},
	}

	for i, form := range testForms {

		_, err := testService.updateComment(context.Background(), 1, form)

		switch i {
		case 0:
			if err != nil {
				t.Error("Error should be nil but got:", err)
			}
			break
		case 1:
			if err != ErrorForm {
				t.Errorf("Error should be %s but got: %s", ErrorForm, err)
			}
			break
		}
	}
}

// Test delete a comment
func testDeleteComment(t *testing.T) {
	err := testService.deleteComment(context.Background(), 1)

	if err != nil {
		t.Error("Error should be nil but got:", err)
	}
}
