package config

type Meta struct {
	Title string `yaml:"title"`
	Icon  string `yaml:"icon"`
}
type Common struct {
	Path      string   `yaml:"path"`
	Name      string   `yaml:"name"`
	Hidden    bool     `yaml:"hidden"`
	Component string   `yaml:"component"`
	Meta      Meta     `yaml:"meta"`
	Children  []Common `yaml:"children"`
}
// 管理员
type admin struct {
	Common
}
//学生
type student struct {
	Common
}
type RouterConfig struct {
	Admin   admin
	Student student
}
