package types

type Ec2Instance struct {
	Name              string `yaml:"name"`
	VirtualServerType string `yaml:"virtual_server_type"`
	ImageId           string `yaml:"image_id"`
	TagName           string `yaml:"tag_name"`
}
