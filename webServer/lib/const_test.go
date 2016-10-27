package lib

import (
	"testing"
	//"fmt"
)

func TestFbCredentials(t *testing.T) {
	cred, _ := FbCredentials()
	if cred == "" {
		t.Error("Return from New should not be nil")
	}
	//TODO
}
