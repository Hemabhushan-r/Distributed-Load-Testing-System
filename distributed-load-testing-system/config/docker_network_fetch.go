package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func main() {
	// Execute the docker network inspect command
	cmd := exec.Command("docker", "network", "inspect", "host")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// Parse the JSON output
	var network []map[string]interface{}
	err = json.Unmarshal(output, &network)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	if len(network) == 0 {
		fmt.Println("Network data is empty")
		return
	}

	// Extract the network information
	networkData := network[0]
	containers := networkData["Containers"].(map[string]interface{})

	// Read the config.json file
	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading config.json:", err)
		return
	}

	// Unmarshal the config.json data
	var config map[string]interface{}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		fmt.Println("Error parsing config.json:", err)
		return
	}

	// Initialize containers field
	if config["containers"] == nil {
		config["containers"] = []interface{}{}
	}
	containersList := config["containers"].([]interface{})

	// Update the config based on container names
	for _, v := range containers {
		container := v.(map[string]interface{})
		name := container["Name"].(string)
		if strings.HasSuffix(name, "1") {
			config["namenode"] = container
		} else if strings.HasSuffix(name, "2") {
			config["client"] = container
		} else {
			containersList = append(containersList, container)
		}
	}
	config["containers"] = containersList

	// Write the updated config back to config.json
	updatedConfig, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling config:", err)
		return
	}
	err = ioutil.WriteFile("config.json", updatedConfig, 0644)
	if err != nil {
		fmt.Println("Error writing config.json:", err)
		return
	}
}
