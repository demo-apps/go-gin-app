// models.article_test.go

package main

import "testing"

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// Check that the length of the list of articles returned is the
	// same as the length of the global variable holding the list
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {

			t.Fail()
			break
		}
	}
}
