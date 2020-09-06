package utils

// Conf : a struct for conf.yml
type Conf struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

// Data : a struct for sensor.yml
type Data struct {
	Computer struct {
		Name string `yaml:"Name"`
	} `yaml:"Computer"`

	Mainboard struct {
		Name        string  `yaml:"Name"`
		Tempratures float32 `yaml:"Temperatures"` // Average
		Fans        struct {
			Fan2 int `yaml:"Fan #2"`
			Fan3 int `yaml:"Fan #3"`
		}
	} `yaml:"Mainboard"`

	CPU struct {
		Name        string  `yaml:"Name"`
		Tempratures float32 `yaml:"Temperatures"` // Average
		Load        struct {
			Total float32 `yaml:"Total"`
		}
	} `yaml:"CPU"`

	Memory struct {
		Load float32 `yaml:"Load"`
	} `yaml:"Memory"`

	GPU struct {
		Name        string  `yaml:"Name"`
		Tempratures float32 `yaml:"Temperatures"`
		Load        struct {
			Core   float32 `yaml:"GPU Core"`
			Memory float32 `yaml:"GPU Memory"`
		}
		Fans int `yaml:"Fans"`
	} `yaml:"GPU"`

	Disks struct {
		Disk0 Disk `yaml:"Disk0"`
		Disk1 Disk `yaml:"Disk1"`
	}
}

// Disk : a struct for a disk's sensors
type Disk struct {
	Name        string  `yaml:"Name"`
	Tempratures float32 `yaml:"Temperatures"` // Packages
	Load        float32 `yaml:"Used Space"`
}
