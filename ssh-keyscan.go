package main

import (
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"strings"
	"strconv"

	"golang.org/x/crypto/ssh"
)

func KeyPrint(dialAddr string, addr net.Addr, key ssh.PublicKey) error {
	fmt.Printf("%s %s %s\n", strings.Split(dialAddr, ":")[0], key.Type(), base64.StdEncoding.EncodeToString(key.Marshal()))
	return nil
}

func invalid_usage() {
	fmt.Fprintf(os.Stderr, "Error: Invalid usage.\n");
	fmt.Fprintf(os.Stderr, "\n");
	fmt.Fprintf(os.Stderr, "Usage: %s [-p port] <host>\n", os.Args[0]);
	os.Exit(1)
}

func main() {
	
	var err error
	
	port := 22
	
	i := 1
	for i+1 < len(os.Args) {
		
		arg := os.Args[i]
		
		if strings.HasPrefix(arg, "-") {
			
			if arg == "-p" {
				if i + 1 >= len(os.Args) {
					invalid_usage()
					return
				}
				port, err = strconv.Atoi(os.Args[i+1])
				if err != nil {
					invalid_usage()
					return
				}
				i++
			} else if arg == "--" {
				i++
				break
			} else {
				break
			}
		}
		
		i++
	}
	
	if i >= len(os.Args) {
		invalid_usage()
		return
	}
	
	host := os.Args[i]
	
	algorithms := []string{
		ssh.KeyAlgoRSA,
		ssh.KeyAlgoDSA,
		ssh.KeyAlgoECDSA256,
		ssh.KeyAlgoSKECDSA256,
		ssh.KeyAlgoECDSA384,
		ssh.KeyAlgoECDSA521,
		ssh.KeyAlgoED25519,
		ssh.KeyAlgoSKED25519,
	}
	
	for _, a := range algorithms {
		
		sshConfig := &ssh.ClientConfig{
			HostKeyAlgorithms: []string{a}, // KeyAlgoXxxx CertAlgoXxxx
			HostKeyCallback: KeyPrint,
		}
		
		client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), sshConfig)
		if err == nil {
			client.Close()
		}
	}
}
