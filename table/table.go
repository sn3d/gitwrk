package table

import "io"

type Aligment int

const (
	AligmentNone  = Aligment(0)
	AligmentLeft  = Aligment(1)
	AligmentRight = Aligment(2)
)

type Header struct {
	Name  string
	Size  int
	Align Aligment
}

type Column struct {
	Value string
}

type Row []*Column

type Table struct {
	Header []*Header
	Rows   []Row
}

func NewTable(h ...Header) *Table {
	table := &Table{
		Header: make([]*Header, len(h)),
		Rows:   make([]Row, 0),
	}

	// set headers
	for i, _ := range h {
		table.Header[i] = &h[i]
		table.Header[i].Size = len(table.Header[i].Name)
	}

	return table
}

func (t *Table) Append(vals ...string) {
	// create row
	row := make([]*Column, len(vals))
	for i, val := range vals {

		// update size of column, header need to know max. size
		// for formatting
		size := len(val)
		if size > t.Header[i].Size {
			t.Header[i].Size = size
		}

		// append column to row
		row[i] = &Column{
			Value: val,
		}
	}

	//append row to table
	t.Rows = append(t.Rows, row)
}

func (t *Table) Print(out io.Writer) {
	// print header
	for i, hRow := range t.Header {
		if i == 0 {
			io.WriteString(out, "| ")
		} else {
			io.WriteString(out, " ")
		}

		io.WriteString(out, align(hRow.Name, hRow))
		io.WriteString(out, " |")
	}
	io.WriteString(out, "\n")

	t.printSplitLine(out)

	// print rows
	for _, row := range t.Rows {

		for i, col := range row {
			head := t.Header[i]
			if i == 0 {
				io.WriteString(out, "| ")
			} else {
				io.WriteString(out, " ")
			}

			io.WriteString(out, align(col.Value, head))
			io.WriteString(out, " |")
		}

		io.WriteString(out, "\n")
	}
}

func (t *Table) printSplitLine(out io.Writer) {
	for i, hRow := range t.Header {
		if i == 0 {
			if hRow.Align == AligmentLeft {
				io.WriteString(out, "|:")
			} else {
				io.WriteString(out, "|-")
			}
		} else {
			if hRow.Align == AligmentLeft {
				io.WriteString(out, ":")
			} else {
				io.WriteString(out, "-")
			}
		}

		for x := 0; x < hRow.Size; x++ {
			io.WriteString(out, "-")
		}

		if hRow.Align == AligmentRight {
			io.WriteString(out, ":|")
		} else {
			io.WriteString(out, "-|")
		}
	}
	io.WriteString(out, "\n")
}

func align(txt string, h *Header) string {
	if h.Align == AligmentLeft {
		return alignLeft(txt, h.Size)
	} else {
		return alignRight(txt, h.Size)
	}
}

func alignLeft(txt string, size int) string {
	res := txt
	if len(txt) <= size {
		for i := len(txt); i < size; i++ {
			res = res + " "
		}
	} else {
		res = res[:size]
	}
	return res
}

func alignRight(txt string, size int) string {
	res := txt
	if len(txt) <= size {
		for i := len(txt); i < size; i++ {
			res = " " + res
		}
	} else {
		res = res[:size]
	}

	return res
}
