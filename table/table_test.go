package table

import (
	"os"
)

func ExampleTable_Print() {
	table := NewTable(
		Header{
			Name:  "C1",
			Align: AligmentLeft,
		},
		Header{
			Name:  "C2",
		},
		Header{
			Name:  "Column-3",
			Align: AligmentLeft,
		},
		Header{
			Name:  "Column-4",
			Align: AligmentRight,
		},
	)

	table.Append("val1", "val2222", "val3", "val7")
	table.Append("val44", "val5", "val6", "val8")

	table.Print(os.Stdout)

	// Output:
	//| C1    |      C2 | Column-3 | Column-4 |
	//|:------|---------|:---------|---------:|
	//| val1  | val2222 | val3     |     val7 |
	//| val44 |    val5 | val6     |     val8 |
}
