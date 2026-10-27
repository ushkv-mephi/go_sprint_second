package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive size",
			args: args{size: 10},
		},
		{
			name: "Zero size", args: args{size: 0}, wantErr: true,
		},
		{
			name:    "Negative size",
			args:    args{size: -1},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateRandomElements(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Fatalf("generateRandomElements() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.args.size > 0 && len(got) != tt.args.size {
				t.Errorf("expected %d elements, got %d", tt.args.size, len(got))
			}

			if tt.wantErr {
				if got != nil {
					t.Errorf("expected nil slice, got %v", got)
				}
				return
			}

			for i, value := range got {
				if value < 1 || value > tt.args.size {
					t.Errorf("element %d: value %d is out of range", i, value)
				}
			}

		})
	}
}

func TestMaximum(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Maximum in the middle",
			args: args{data: []int{1, 2, 3, 10, 4}},
			want: 10,
		},
		{
			name:    "Empty slice",
			args:    args{data: []int{}},
			want:    0,
			wantErr: true,
		},
		{
			name: "Single element",
			args: args{data: []int{1}},
			want: 1,
		},
		{
			name: "Equal elements",
			args: args{data: []int{5, 5, 5}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := maximum(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Fatalf("maximum() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Maximum in the last chunk with remainder",
			args:    args{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 100, 16}},
			want:    100,
			wantErr: false,
		},
		{
			name:    "Empty slice",
			args:    args{data: []int{}},
			want:    0,
			wantErr: true,
		},
		{
			name:    "Slice smaller than chunk count",
			args:    args{data: []int{1, 2, 3, 4, 5}},
			want:    5,
			wantErr: false,
		},
		{
			name:    "Maximum in the middle",
			args:    args{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 50, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}},
			want:    50,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := maxChunks(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Fatalf("maxChunks() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("maxChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
