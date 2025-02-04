package commons

import "testing"

func TestToFilename(t *testing.T) {
	testcases := []struct {
		Name     string
		Args     []string
		Expected string
	}{
		{
			Name:     "MyFileName",
			Expected: "my_file_name.go",
		},
		{
			Name:     "myFileName",
			Expected: "my_file_name.go",
		},
		{
			Name:     "my_file_name",
			Expected: "my_file_name.go",
		},
		{
			Name:     "myFileName",
			Args:     []string{"_repository"},
			Expected: "my_file_name_repository.go",
		},
	}

	for _, tc := range testcases {
		result := ToFilename(tc.Name, tc.Args...)
		if result != tc.Expected {
			t.Errorf("expected %v got: %v", tc.Expected, result)
		}
	}
}
