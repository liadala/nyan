package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDirStruct(t *testing.T) {
	// create a temporary directory
	tempDir := t.TempDir()
	subDir := filepath.Join(tempDir, "subdir")
	err := os.Mkdir(subDir, 0755)
	if err != nil {
		t.Fatalf("failed to create subdir: %v", err)
	}

	// create files files
	files := []string{
		filepath.Join(tempDir, "file1.txt"),
		filepath.Join(tempDir, "file2.log"),
		filepath.Join(subDir, "file3.md"),
	}
	for _, file := range files {
		err := os.WriteFile(file, []byte("test content"), 0644)
		if err != nil {
			t.Fatalf("failed to create test file %s: %v", file, err)
		}
	}

	// Tests
	tests := []struct {
		name      string
		recursive bool
		expected  []string
	}{
		{
			name:      "Non-recursive",
			recursive: false,
			expected: []string{
				"📁",
				"┣📄 file1.txt",
				"┣📄 file2.log",
				"┗📁 subdir",
			},
		},
		{
			name:      "Recursive",
			recursive: true,
			expected: []string{
				"📁",
				"┣📄 file1.txt",
				"┣📄 file2.log",
				"┗📁 subdir",
				"  ┗📄 file3.md",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var builder strings.Builder
			builder.WriteString(fmt.Sprintf("📁%s\n", tempDir))
			err := dirstruct(&builder, tempDir, " ", test.recursive)
			if err != nil {
				t.Fatalf("dirstruct failed: %v", err)
			}

			output := builder.String()
			for _, expected := range test.expected {
				if !strings.Contains(output, expected) {
					t.Errorf("expected output to contain: %s\ngot:\n%s", expected, output)
				}
			}
		})
	}
}

func BenchmarkDirStruct(b *testing.B) {
	tempDir := b.TempDir()
	subDir := filepath.Join(tempDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		b.Fatalf("failed to create subdir: %v", err)
	}

	numFiles := 100
	for i := 0; i < numFiles; i++ {
		fileName := filepath.Join(tempDir, fmt.Sprintf("file%d.txt", i))
		if err := os.WriteFile(fileName, []byte("test content"), 0644); err != nil {
			b.Fatalf("failed to create test file %s: %v", fileName, err)
		}
	}

	for i := 0; i < numFiles; i++ {
		fileName := filepath.Join(subDir, fmt.Sprintf("file%d.txt", i))
		if err := os.WriteFile(fileName, []byte("test content"), 0644); err != nil {
			b.Fatalf("failed to create test file %s: %v", fileName, err)
		}
	}

	benchmarks := []struct {
		name      string
		recursive bool
	}{
		{name: "Non-recursive", recursive: false},
		{name: "Recursive", recursive: true},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var builder strings.Builder
				builder.WriteString(fmt.Sprintf("📁%s\n", tempDir))
				err := dirstruct(&builder, tempDir, " ", bm.recursive)
				if err != nil {
					b.Fatalf("dirstruct failed: %v", err)
				}
			}
		})
	}
}
