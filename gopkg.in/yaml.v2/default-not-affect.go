package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// 今更v2なうえにdefaultタグを付けてるコードがあって、効果ないことの確認

func main() {
	type Config struct {
		Name string `yaml:"name" default:"default-name"`
		Some string `yaml:"some"`
	}

	var c Config
	yamlData := []byte("some: val")
	err := yaml.Unmarshal(yamlData, &c)
	if err != nil {
		panic(err)
	}

	/*
		main.Config{Name:"", Some:"val"}
	*/
	fmt.Printf("%#v\n", c)
}
