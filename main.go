package main

import (
	"debug/elf"
	"fmt"
	"os"
)

func main() {
	const payload_elf_section_name = "payload"

	selfPath, _ := os.Executable()
	selfFile, _ := os.Open(selfPath)
	selfElf, _ := elf.NewFile(selfFile)
	var data []byte
	found := false
	for _, s := range selfElf.Sections {
		if s.Name == payload_elf_section_name {
			data, _ = s.Data()
			found = true
		}
	}
	if found {
		fmt.Println("found data: ", string(data))
	} else {
		fmt.Printf("did not find elf section \"%s\"\n", payload_elf_section_name)
	}

}
