package config

type Route struct {
	Path     string  `json:"path" yaml:"path"`
	Name     string  `json:"name" yaml:"name"`
	Children []Route `json:"children" yaml:"children"`
}

type RouterConfig struct {
	Admin   []Route
	Student []Route
	Dorm    []Route
}
