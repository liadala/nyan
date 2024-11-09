// Copyright (c) 2024 liadala
//
// This work is licensed under the Creative Commons Attribution-ShareAlike 4.0 International License (CC BY-SA 4.0).
// You are free to share and adapt this work, as long as you provide appropriate attribution and distribute any
// derivative works under the same license.
//
// A copy of the license is available at:
//
//     https://creativecommons.org/licenses/by-sa/4.0/
//
// UNLESS REQUIRED BY APPLICABLE LAW OR AGREED TO IN WRITING, THIS SOFTWARE IS PROVIDED "AS IS", WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	basePath := flag.String("path", "./", "start path")
	flag.Parse()

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("📁%s\n", *basePath))

	if err := dirstruct(&builder, *basePath, " "); err != nil {
		fmt.Println("cant create view:", err)
	} else {
		fmt.Print(builder.String())
	}
}

func dirstruct(builder *strings.Builder, path string, prefix string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("cant read dir %s: %v", path, err)
	}

	for i, entry := range entries {
		last := i == len(entries)-1
		connector := "┣"
		if last {
			connector = "┗"
		}

		if entry.IsDir() {
			builder.WriteString(fmt.Sprintf("%s%s📁 %s\n", prefix, connector, entry.Name()))
			newPrefix := prefix + "┃ "
			if last {
				newPrefix = prefix + "  "
			}
			if err := dirstruct(builder, filepath.Join(path, entry.Name()), newPrefix); err != nil {
				return err
			}
		} else {
			builder.WriteString(fmt.Sprintf("%s%s%s %s\n", prefix, connector, "📄", entry.Name()))
		}
	}
	return nil
}
