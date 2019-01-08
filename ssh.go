package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

type Config struct {
	Username string `json: "Username"`
	Password string `json: "Password"`
	SSH_Port string `json: "SSH_Port"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-address command\r\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	command := os.Args[2]
	addr := net.ParseIP(name)
	address := addr.String()
	if addr == nil {
		fmt.Printf("Invalid IP address %s", address)
		logger("Invalid IP address " + address + "\r\n")
	} else {
		logger("Running job on " + address + "\r\n")
		fmt.Printf("SSHing to %s and running command %s\r\n", address, command)
		logger("SSHing to " + address + " and running command " + command + "\r\n")
		execute(address, command)
	}
	os.Exit(0)
}

func logger(s string) {
	file, err := os.OpenFile("ssh.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("Error opening ssh log file")
		logger("Error opening ssh log file\r\n")
	}
	log.SetOutput(file)
	log.Print(s)
}

func loadConfig(filename string) Config {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		fmt.Println("Error opening config file")
		logger("Error opening config file\r\n")
		os.Exit(1)
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding json file")
		logger("Error decoding json file\r\n")
		os.Exit(1)
	}

	return config
}

func execute(ip, command string) {
	config := loadConfig("config.json")

	sshAddress := ip + config.SSH_Port
	sshConfig := &ssh.ClientConfig{
		User: config.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", sshAddress, sshConfig)
	if err != nil {
		fmt.Println("Error connecting to host")
		logger("Error connecting to host " + sshAddress + "\r\n")
		fmt.Println(err)
		os.Exit(1)
	}
	session, err := conn.NewSession()
	defer session.Close()
	if err != nil {
		fmt.Println("Error creating new session")
		logger("Error creating new session\r\n")
		fmt.Println(err)
		os.Exit(1)
	}

	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer
	stdoutBuf.Grow(1000000)
	stderrBuf.Grow(1000000)
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf
	session.Run(command)
	if strings.TrimSpace(stdoutBuf.String()) == "" {
		fmt.Println(stderrBuf.String())
		fmt.Print("You may need to use Sudo")
	} else {
		fmt.Print(stdoutBuf.String())
	}
}
