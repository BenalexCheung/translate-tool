package utils

import (
	"fmt"
	"testing"
)

func TestXLS_ParsingXLS(t *testing.T) {
	type fields struct {
		FileName string
		Rows     [][]string
		Cols     [][]string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "testXLS1",
			fields: fields{
				FileName: "Base.xlsx",
				Rows:     nil,
				Cols:     nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XLS{
				FileName: tt.fields.FileName,
				Rows:     tt.fields.Rows,
				Cols:     tt.fields.Cols,
			}
			if err := x.ParsingXLS(); (err != nil) != tt.wantErr {
				t.Errorf("XLS.ParsingXLS() error = %v, wantErr %v", err, tt.wantErr)
			}
			for _, col := range x.Cols {
				for _, rowCell := range col {
					fmt.Print(rowCell, "\n")
				}
				fmt.Println()
			}
		})
	}
}

func TestXLS_SaveXLS(t *testing.T) {
	type fields struct {
		FileName string
		Rows     [][]string
		Cols     [][]string
	}
	type args struct {
		axis   string
		newCol []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "testXLS2",
			fields: fields{
				FileName: "en.xlsx",
				Rows:     nil,
				Cols:     nil,
			},
			args: args{
				axis:   "B",
				newCol: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XLS{
				FileName: tt.fields.FileName,
				Rows:     tt.fields.Rows,
				Cols:     tt.fields.Cols,
			}
			if err := x.ParsingXLS(); (err != nil) != tt.wantErr {
				t.Errorf("XLS.ParsingXLS() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.args.newCol = x.Cols[0]
			if err := x.SaveXLS(tt.args.axis, tt.args.newCol); (err != nil) != tt.wantErr {
				t.Errorf("XLS.SaveXLS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
