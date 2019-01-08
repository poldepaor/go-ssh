## Go-SSH
A Linux/Windows command line SSH tool to SSH to remote Linux hosts run your command and print the result in the Windows terminal.

## Build
Linux(if running on Windows): Clone the repo and build the Linux binary by running `GOOS=Linux go build ssh.go` inside the parent directory of the repo. Else just run `go build ssh.go`
Windows: Clone the repo and build the Windows executable by running `go build ssh.go` inside the parent directory of the repo.

## Config
Enter the username and password for the Linux host in the `config.json` file.

## Run
Run the tool as follows `ssh 10.10.10.10 "ls /home"`

## Logging
The tool will create a `ssh.log` file with the command run, the host and a timestamp.
