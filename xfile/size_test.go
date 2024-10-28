package xfile

import "testing"

func TestSize(t *testing.T) {
	type args struct {
		bytes int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-1",
			args: args{bytes: 1024},
			want: "1.0KB",
		},
		{
			name: "test-2",
			args: args{bytes: 1024 * 1024},
			want: "1.0MB",
		},
		{
			name: "test-3",
			args: args{bytes: 1024 * 1024 * 1024},
			want: "1.0GB",
		},
		{
			name: "test-4",
			args: args{bytes: 1024 * 1024 * 1024 * 1024},
			want: "1.0TB",
		},
		{
			name: "test-5",
			args: args{bytes: 1024 * 1024 * 1024 * 1024 * 1024},
			want: "1.0PB",
		},
		{
			name: "test-6",
			args: args{bytes: 2048},
			want: "2.0KB",
		},
		{
			name: "test-6",
			args: args{bytes: 2560},
			want: "2.5KB",
		},
		{
			name: "test-7",
			args: args{bytes: 2565}, // rounding
			want: "2.5KB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Size(tt.args.bytes); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBytes(t *testing.T) {
	type args struct {
		sizeStr string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "test-1",
			args:    args{sizeStr: "1KB"},
			want:    1024,
			wantErr: false,
		},
		{
			name:    "test-2",
			args:    args{sizeStr: "1MB"},
			want:    1024 * 1024,
			wantErr: false,
		},
		{
			name:    "test-3",
			args:    args{sizeStr: "1GB"},
			want:    1024 * 1024 * 1024,
			wantErr: false,
		},
		{
			name:    "test-4",
			args:    args{sizeStr: "1TB"},
			want:    1024 * 1024 * 1024 * 1024,
			wantErr: false,
		},
		{
			name:    "test-5",
			args:    args{sizeStr: "1K"},
			want:    1024,
			wantErr: false,
		},
		{
			name:    "test-6",
			args:    args{sizeStr: "1M"},
			want:    1024 * 1024,
			wantErr: false,
		},
		{
			name:    "test-7",
			args:    args{sizeStr: "1G"},
			want:    1024 * 1024 * 1024,
			wantErr: false,
		},
		{
			name:    "test-8",
			args:    args{sizeStr: "1T"},
			want:    1024 * 1024 * 1024 * 1024,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBytes(tt.args.sizeStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
