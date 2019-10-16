package comment

import (
	"context"
	"testing"
)

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
