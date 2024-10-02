package fileresolver

import "testing"

func Test_windowsToPosix(t *testing.T) {
	type args struct {
		windowsPath string
	}
	tests := []struct {
		name          string
		args          args
		wantPosixPath string
	}{
		{
			name: "basic case",
			args: args{
				windowsPath: `C:\some\windows\place`,
			},
			wantPosixPath: "/c/some/windows/place",
		},
		{
			name: "escaped case",
			args: args{
				windowsPath: `C:\\some\\windows\\place`,
			},
			wantPosixPath: "/c/some/windows/place",
		},
		{
			name: "forward slash",
			args: args{
				windowsPath: `C:/foo/bar`,
			},
			wantPosixPath: "/c/foo/bar",
		},
		{
			name: "mix slash",
			args: args{
				windowsPath: `C:\foo/bar\`,
			},
			wantPosixPath: "/c/foo/bar",
		},
		{
			name: "case sensitive case",
			args: args{
				windowsPath: `C:\Foo/bAr\`,
			},
			wantPosixPath: "/c/Foo/bAr",
		},
		{
			name: "special char case",
			args: args{
				windowsPath: `C:\ふー\バー`,
			},
			wantPosixPath: "/c/ふー/バー",
		},
		{
			name: "basic relative case",
			args: args{
				windowsPath: `some\windows\place`,
			},
			wantPosixPath: "some/windows/place",
		},
		{
			name: "escaped relative case",
			args: args{
				windowsPath: `some\\windows\\place`,
			},
			wantPosixPath: "some/windows/place",
		},
		{
			name: "forward relative slash",
			args: args{
				windowsPath: `foo/bar`,
			},
			wantPosixPath: "foo/bar",
		},
		{
			name: "mix relative slash",
			args: args{
				windowsPath: `foo/bar\`,
			},
			wantPosixPath: "foo/bar",
		},
		{
			name: "case sensitive relative case",
			args: args{
				windowsPath: `Foo/bAr\`,
			},
			wantPosixPath: "Foo/bAr",
		},
		{
			name: "special char relative case",
			args: args{
				windowsPath: `ふー\バー`,
			},
			wantPosixPath: "ふー/バー",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPosixPath := windowsToPosix(tt.args.windowsPath); gotPosixPath != tt.wantPosixPath {
				t.Errorf("windowsToPosix() = %v, want %v", gotPosixPath, tt.wantPosixPath)
			}
		})
	}
}

func Test_posixToWindows(t *testing.T) {
	type args struct {
		posixPath string
	}
	tests := []struct {
		name            string
		args            args
		wantWindowsPath string
	}{
		{
			name: "basic case",
			args: args{
				posixPath: "/c/some/windows/place",
			},
			wantWindowsPath: `C:\some\windows\place`,
		},
		{
			name: "escaped case",
			args: args{
				posixPath: "/c/some/windows/place",
			},
			wantWindowsPath: `C:\some\windows\place`,
		},
		{
			name: "forward slash",
			args: args{
				posixPath: "/c/foo/bar",
			},
			wantWindowsPath: `C:\foo\bar`,
		},
		{
			name: "mix slash",
			args: args{
				posixPath: "/c/foo/bar",
			},
			wantWindowsPath: `C:\foo\bar`,
		},
		{
			name: "case sensitive case",
			args: args{
				posixPath: "/c/Foo/bAr",
			},
			wantWindowsPath: `C:\Foo\bAr`,
		},
		{
			name: "special char case",
			args: args{
				posixPath: "/c/ふー/バー",
			},
			wantWindowsPath: `C:\ふー\バー`,
		},
		{
			name: "basic relative case",
			args: args{
				posixPath: "some/windows/place",
			},
			wantWindowsPath: `some\windows\place`,
		},
		{
			name: "escaped relative case",
			args: args{
				posixPath: "some/windows/place",
			},
			wantWindowsPath: `some\windows\place`,
		},
		{
			name: "forward relative slash",
			args: args{
				posixPath: "foo/bar",
			},
			wantWindowsPath: `foo\bar`,
		},
		{
			name: "mix relative slash",
			args: args{
				posixPath: "foo/bar",
			},
			wantWindowsPath: `foo\bar`,
		},
		{
			name: "case sensitive relative case",
			args: args{
				posixPath: "Foo/bAr",
			},
			wantWindowsPath: `Foo\bAr`,
		},
		{
			name: "special char relative case",
			args: args{
				posixPath: "ふー/バー",
			},
			wantWindowsPath: `ふー\バー`,
		},
		{
			name: "leading slash relative case",
			args: args{
				posixPath: "/some/windows/place",
			},
			wantWindowsPath: `\some\windows\place`,
		},
		{
			name: "leading slash non-drive relative case",
			args: args{
				posixPath: "/1/some/windows/place",
			},
			wantWindowsPath: `\1\some\windows\place`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotWindowsPath := posixToWindows(tt.args.posixPath); gotWindowsPath != tt.wantWindowsPath {
				t.Errorf("posixToWindows() = %v, want %v", gotWindowsPath, tt.wantWindowsPath)
			}
		})
	}
}
