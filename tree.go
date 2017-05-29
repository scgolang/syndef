package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/scgolang/sc"
)

func (c *controller) writeTree(w io.Writer, d *sc.Synthdef) error {
	buf := &bytes.Buffer{}

	if err := d.WriteJSON(buf); err != nil {
		return err
	}
	s := synthdef{}

	if err := json.NewDecoder(buf).Decode(&s); err != nil {
		return err
	}
	return tree(s, s.root(), "")
}

func tree(s synthdef, ugenIndex int, prefix string) error {
	u := s.Ugens[ugenIndex]

	fmt.Printf("%s(%d)\n", u.Name, ugenIndex)

	for i, in := range u.Inputs {
		fmt.Printf(prefix + "|- ")
		if isConstant(in) {
			fmt.Printf("%f\n", s.Constants[in.OutputIndex])
			continue
		}
		if i == len(u.Inputs)-1 {
			tree(s, in.UgenIndex, prefix+"   ")
		} else {
			tree(s, in.UgenIndex, prefix+"|  ")
		}
	}
	return nil
}