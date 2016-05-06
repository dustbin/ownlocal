package main

type TestReader struct {
	contents []byte
}

func (t TestReader) Read(p []byte) (n int, err error) {
	length := len(p)
	if len(t.contents) < length {
		length = len(t.contents)
	}

	for ; n < length; n++ {
		p[n] = t.contents[n]
	}

	return
}

//
// func TestNewBusinessDb(t *testing.T) {
// 	reader := TestReader{contents: []byte{0, 1, 2, 3}}
//
// 	_, err := NewBusinessDB(reader)
// 	if err != nil {
// 		t.Errorf("error was not nil: %#v", err)
// 	}
// }
