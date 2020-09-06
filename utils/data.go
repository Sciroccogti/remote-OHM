package utils

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"

	"gopkg.in/yaml.v2"
)

// SaveData : save data of sensors to data.yml
func SaveData(ohmdata OhmData) {
	fmt.Println("Saving data.yml ...")
	data := new(Data)

	diskcount := 0

	for _, computer := range ohmdata.Children {
		if computer.ImageURL != "images_icon/computer.png" {
			fmt.Println("No computer found!")
			return
		}
		data.Computer.Name = computer.Name

		for _, component := range computer.Children {
			switch component.ImageURL {
			case "images_icon/mainboard.png":
				data.Mainboard.Name = component.Name

				for _, sensors := range component.Children[0].Children {
					switch sensors.Name {
					case "Temperatures":
						data.Mainboard.Tempratures = extractTemperature(sensors)

					case "Fans":
						for _, fan := range sensors.Children {
							re := regexp.MustCompile(`(\d+) RPM`)
							match := re.FindStringSubmatch(fan.Value)
							if match != nil {
								tmp, err := strconv.ParseInt(match[1], 10, 32)
								if err == nil && tmp > 0 {
									if fan.Name == "Fan #2" {
										data.Mainboard.Fans.Fan2 = int(tmp)
									} else if fan.Name == "Fan #3" {
										data.Mainboard.Fans.Fan3 = int(tmp)
									}
								}
							}
						}
					}
				}

			case "images_icon/cpu.png":
				data.CPU.Name = component.Name

				for _, sensors := range component.Children {
					switch sensors.Name {
					case "Temperatures":
						data.CPU.Tempratures = extractTemperature(sensors)

					case "Load":
						for _, load := range sensors.Children {
							if load.Name == "CPU Total" {
								data.CPU.Load.Total = extractLoad(load)
							}
						}
					}

				}
			case "images_icon/ram.png":
				for _, sensors := range component.Children {
					if sensors.Name == "Load" {
						data.Memory.Load = extractLoad(sensors.Children[0])
					}
				}

			case "images_icon/nvidia.png":
				data.GPU.Name = component.Name

				for _, sensors := range component.Children {
					switch sensors.Name {
					case "Temperatures":
						data.GPU.Tempratures = extractTemperature(sensors)

					case "Load":
						for _, load := range sensors.Children {
							switch load.Name {
							case "GPU Core":
								data.GPU.Load.Core = extractLoad(load)
							case "GPU Memory":
								data.GPU.Load.Memory = extractLoad(load)
							}

						}

					case "Fans":
						for _, fan := range sensors.Children {
							re := regexp.MustCompile(`(\d+) RPM`)
							match := re.FindStringSubmatch(fan.Value)
							if match != nil {
								tmp, err := strconv.ParseInt(match[1], 10, 32)
								if err == nil && tmp > 0 {
									data.GPU.Fans = int(tmp)
								}
							}
						}
					}
				}

			case "images_icon/hdd.png":
				var disk Disk
				disk.Name = component.Name

				for _, sensors := range component.Children {
					switch sensors.Name {
					case "Temperatures":
						disk.Tempratures = extractTemperature(sensors)

					case "Load":
						for _, load := range sensors.Children {
							if load.Name == "Used Space" {
								disk.Load = extractLoad(load)
							}
						}
					}

				}

				switch diskcount {
				case 0:
					data.Disks.Disk0 = disk
				case 1:
					data.Disks.Disk1 = disk
				}
				diskcount++
			}
		}
	}
	// tmp := ohmdata
	// data :=
	dataYaml, err := yaml.Marshal(data)
	if err != nil {
		fmt.Printf("Error while saving data.yml :\n")
		fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
	}
	err = ioutil.WriteFile("data.yml", dataYaml, 0644)
}

// extractTemperature : extract a temperature data from OhmData
func extractTemperature(sensors OhmData) float32 {
	tempsum := 0.0
	count := 0
	for _, temperature := range sensors.Children {
		re := regexp.MustCompile(`(-?\d+\.\d{1,2}) °C`)
		match := re.FindStringSubmatch(temperature.Value)
		if match != nil {
			tmp, err := strconv.ParseFloat(match[1], 32)
			if err == nil && tmp > 0 {
				tempsum += tmp
				count++
			}
		}
	}
	if count > 0 {
		return float32(tempsum) / float32(count)
	}
	return 0.0
}

// extractLoad : extract a load data from OhmData
func extractLoad(load OhmData) float32 {
	re := regexp.MustCompile(`(-?\d+\.\d{1,2}) %`)
	match := re.FindStringSubmatch(load.Value)
	if match != nil {
		tmp, err := strconv.ParseFloat(match[1], 32)
		if err == nil && tmp > 0 {
			return float32(tmp)
		}
	}

	return 0.0
}
