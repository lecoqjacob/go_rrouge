package main

import (
	"bytes"
	"encoding/gob"

	"github.com/anaseto/gruid"
)

type config struct {
	NormalModeKeys map[gruid.Key]actionType
	TargetModeKeys map[gruid.Key]actionType
	DarkLOS        bool
	Tiles          bool
	Version        string
	ShowNumbers    bool
}

func (c *config) ConfigSave() ([]byte, error) {
	data := bytes.Buffer{}
	enc := gob.NewEncoder(&data)
	err := enc.Encode(c)
	if err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}

func DecodeConfigSave(data []byte) (*config, error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	c := &config{}
	err := dec.Decode(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
