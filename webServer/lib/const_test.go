package lib

import (
	"testing"
	//"fmt"
)

func TestFbCredentials(t *testing.T) {
	cred := FbCredentials()
	if cred == nil {
		t.Error("Return from New should not be nil")
	}
	//TODO
}
