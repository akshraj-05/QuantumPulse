package ping

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

// Function to extract round-trip time (RTT) from ping output
func extractRTT(pingOutput string) float64 {
	rttPattern := regexp.MustCompile(`time=([0-9]+(\.[0-9]+)?)`)
	match := rttPattern.FindStringSubmatch(pingOutput)
	if len(match) > 1 {
		rtt, err := strconv.ParseFloat(match[1], 64)
		if err == nil {
			return rtt
		}
	}
	return -1 // Return -1 if RTT not found
}

func Ping(str string) (float64, error) {
	// Command to execute ping (replace "example.com" with your desired destination)
	cmd := exec.Command("ping", str)

	// Run the ping command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing ping command:", err)
		fmt.Println("Output:", string(output))
		return -1, err
	}

	// Convert the output bytes to string
	pingOutput := string(output)
	//	fmt.Println("Ping output:", pingOutput)

	// Extract RTT from ping output
	rtt := extractRTT(pingOutput)
	if rtt == -1 {
		return -1, fmt.Errorf("RTT not found in ping output")
	}

	return rtt, nil
}
