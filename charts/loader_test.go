package charts

import (
	"os"
	"testing"
)

//
// func TestReadTemplate(t *testing.T) {
// 	v, err := ReadTemplate("templates/timeseries.tmpl")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(v)
// }

func TestWriteTemplate(t *testing.T) {
	v, err := writeTemplate(templateList["base.tmpl"])
	if err != nil {
		t.Fatal(err)
	}
	err = v.Close()
	if err != nil {
		t.Fatal("failed to close file manual removal may be required")
	}
	err = os.Remove(v.Name())
	if err != nil {
		t.Fatal("failed to remove file manual removal may be required")
	}
}
