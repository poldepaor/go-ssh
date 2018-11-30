# Go-SSH
A Windows command line SSH tool to SSH to remote Linux hosts run your command and print the result in the Windows terminal.

# Build
Clone the repo and build the Windows executable by running `go build ssh.go` inside the parent directory of the repo.

# Config
Enter the username and password for the Linux host in the `config.json` file.

# Run
Run the tool as follows `ssh 10.10.10.10 "ls /home"`
