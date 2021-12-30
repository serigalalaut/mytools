## mytools CLI

Task for backend developer.

## Quick Start

- Clone this repo or download it's release archive and extract it somewhere
- You may delete `.git` folder if you get this code via `git clone`
- Run `go mod download`
- Run `go build mytools`
- Well Done. 

## Cli Command

- `./mytools /var/log/nginx/error.log` get file format plaintext
- `./mytools -t json /var/log/nginx/error.log` get file format json
- `./mytools -t text /var/log/nginx/error.log` get file format plaintext
- `./mytools output -o /var/log/nginx/error.log /User/johnmayer/Desktop/nginxlog.txt` get file format plaintext and save file in /User/johnmayer/Desktop
- `./mytools output -o json /var/log/nginx/error.log /User/johnmayer/Desktop/nginxlog.json` 
- `./mytools -h` display information about command

## License
