package config

type Route struct {
	Path     string  `json:"path" yaml:"path"`
	Name     string  `json:"name" yaml:"name"`
	Children []Route `json:"children" yaml:"children"`
	// Hidden    bool     `json:"hidden" yaml:"hidden"`
	// Component string   `json:"component" yaml:"component"`
	// Meta      Meta     `json:"meta" yaml:"meta"`
}

type RouterConfig struct {
	Admin   []Route
	Student []Route
	Dorm    []Route
}
