package mypackage

import "testing"

// MockFileReader implements the FileReader interface for testing purposes.
type MockFileReader struct {
	Data   []byte
	Err    error
	Closed bool
}

func (m *MockFileReader) ReadFile(filename string) ([]byte, error) {
	return m.Data, m.Err
}

func (m *MockFileReader) Close() error {
	m.Closed = true
	return nil
}

func TestProcessFile(t *testing.T) {
	mockReader := &MockFileReader{
		Data: []byte("Mocked data"),
		Err:  nil,
	}

	err := ProcessFile(mockReader)
	if err != nil {
		t.Errorf("Error processing file: %v", err)
	}

	if !mockReader.Closed {
		t.Error("Expected the file reader to be closed")
	}
}
