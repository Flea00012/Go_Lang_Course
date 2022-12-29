package fullapp

import (
	"testing"
)

func TestIPFS(t *testing.T)  {
	s := connectToIPFS()
	s.addToIPFS()
	s.getDataFromIPFS()
}
