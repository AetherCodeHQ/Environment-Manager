package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: environment-manager <.env> [get|set|list|diff]")
		os.Exit(1)
	}
	path := os.Args[1]
	data, _ := os.ReadFile(path)
	vars := map[string]string{}
	for _, line := range strings.Split(string(data), "\n") {
		t := strings.TrimSpace(line)
		if t == "" || strings.HasPrefix(t, "#") {
			continue
		}
		parts := strings.SplitN(t, "=", 2)
		if len(parts) == 2 {
			vars[parts[0]] = parts[1]
		}
	}
	cmd := "list"
	if len(os.Args) > 2 {
		cmd = os.Args[2]
	}
	switch cmd {
	case "list":
		keys := make([]string, 0, len(vars))
		for k := range vars {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			v := vars[k]
			if len(v) > 20 {
				v = v[:8] + "..."
			}
			fmt.Printf("%s=%s\n", k, v)
		}
		fmt.Printf("\n%d variables\n", len(vars))
	case "get":
		if len(os.Args) > 3 {
			fmt.Println(vars[os.Args[3]])
		}
	case "set":
		if len(os.Args) > 4 {
			vars[os.Args[3]] = os.Args[4]
			saveEnv(path, vars)
			fmt.Println("updated")
		}
	case "diff":
		if len(os.Args) > 3 {
			other, _ := os.ReadFile(os.Args[3])
			ov := map[string]string{}
			for _, line := range strings.Split(string(other), "\n") {
				parts := strings.SplitN(strings.TrimSpace(line), "=", 2)
				if len(parts) == 2 {
					ov[parts[0]] = parts[1]
				}
			}
			for k, v := range vars {
				ov2, ok := ov[k]
				if !ok {
					fmt.Printf("+ %s=%s\n", k, v)
				} else if ov2 != v {
					fmt.Printf("~ %s: %s -> %s\n", k, v, ov2)
				}
			}
		}
	}
}

func saveEnv(path string, vars map[string]string) {
	keys := make([]string, 0, len(vars))
	for k := range vars {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b strings.Builder
	for _, k := range keys {
		fmt.Fprintf(&b, "%s=%s\n", k, vars[k])
	}
	os.WriteFile(path, []byte(b.String()), 0644)
}
