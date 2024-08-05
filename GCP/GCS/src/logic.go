package main

// func performSFTPRequest() error {
// 	hostKeyCallback, err := knownhosts.New("/path/to/known_hosts")
// 	if err != nil {
// 		return fmt.Errorf("failed to create host key callback: %v", err)
// 	}

// 	config := &ssh.ClientConfig{
// 		User: "your-username",
// 		Auth: []ssh.AuthMethod{
// 			ssh.Password("your-password"),
// 		},
// 		HostKeyCallback: hostKeyCallback,
// 	}

// 	conn, err := ssh.Dial("tcp", sftpServerAddr, config)
// 	if err != nil {
// 		return fmt.Errorf("SFTP connection failed: %v", err)
// 	}
// 	defer conn.Close()

// 	client, err := sftp.NewClient(conn)
// 	if err != nil {
// 		return fmt.Errorf("failed to create SFTP client: %v", err)
// 	}
// 	defer client.Close()

// 	// Perform an example operation to verify the connection
// 	_, err = client.ReadDir("/")
// 	if err != nil {
// 		return fmt.Errorf("failed to read directory: %v", err)
// 	}

// 	return nil
// }

func performHTTPSRequest() {
}
