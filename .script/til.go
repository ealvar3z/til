package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func main() {
	root, _ := filepath.Abs(filepath.Dir(filepath.Dir(os.Args[0])))
	target := filepath.Join(root, "README.md")
	putInfo("start update README.md")

	putInfo("write header")
	f, _ := os.Create(target)
	defer f.Close()
	fmt.Fprintln(f, "# TIL")
	fmt.Fprintln(f, "Today I Learned")
	fmt.Fprintln(f, "- - -")

	putInfo("pick categories")
	categories, _ := filepath.Glob(filepath.Join(root, "[a-zA-Z]*"))
	sort.Strings(categories)

	putInfo("start write categories TOC")
	fmt.Fprintln(f, "## Categories")
	for _, p := range categories {
		categoryName := filepath.Base(p)
		putInfo("write category: " + categoryName)
		fmt.Fprintf(f, "- [%s](#%s)\n", categoryName, categoryName)
	}
	fmt.Fprintln(f, "- - -")

	putInfo("write content TOC")
	for _, p := range categories {
		categoryName := filepath.Base(p)
		putInfo("write category: " + categoryName)
		fmt.Fprintf(f, "### %s\n\n", strings.Title(categoryName))

		files, _ := filepath.Glob(filepath.Join(p, "*"))
		sort.Strings(files)
		for _, file := range files {
			title, _ := os.Open(file)
			defer title.Close()
			b := make([]byte, 100)
			n, _ := title.Read(b)
			titleStr := string(b[:n])
			titleStr = strings.Trim(titleStr, "# ")
			putInfo("write content: " + titleStr)
			fmt.Fprintf(f, "- [%s](%s)\n", titleStr, strings.TrimPrefix(file, root+"/"))
		}
		fmt.Fprintln(f)
	}

	putInfo("complete")
}

func putInfo(s string) {
	fmt.Printf("[%s] %s\n", time.Now().Format("2006-01-02.15:04:05.000000000"), s)
}
