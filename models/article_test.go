package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// Comparing length of test returned value with global value in main.go
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// Comparing test value fields against global value fields.  
	for i,v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			
			t.Fail()
			break
			}
	}
}



