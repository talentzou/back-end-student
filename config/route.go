package config

type Meta struct {
	Title string `json:"title" yaml:"title"`
	Icon  string `json:"icon" yaml:"icon"`
}
type Common struct {
	Path      string   `json:"path" yaml:"path"`
	Name      string   `json:"name" yaml:"name"`
	Hidden    bool     `json:"hidden" yaml:"hidden"`
	Component string   `json:"component" yaml:"component"`
	Meta      Meta     `json:"meta" yaml:"meta"`
	Children  []Common `json:"children" yaml:"children"`
}

type RouterConfig struct {
	Admin   []Common
	Student []Common
}
