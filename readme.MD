# Go Directory Finder

This is a simple Go script that checks for common directories on a given website using a wordlist. It helps in finding hidden directories on web servers.

## Features
- Scans a target website for common directories.
- Uses a wordlist for directory enumeration.
- Prints found directories with `200 OK` status.

## Requirements
- Go installed on your system.
- A wordlist file with common directory names.

## Installation
1. Clone this repository:
   ```sh
   git clone https://github.com/your-username/go-directory-finder.git
   cd go-directory-finder
   ```
2. Build the project (optional):
   ```sh
   go build -o dirfinder main.go
   ```

## Usage
Run the script using:
```sh
go run main.go <URL> <wordlist>
```
Example:
```sh
go run main.go http://example.com wordlist.txt
```
Or if built:
```sh
./dirfinder http://example.com wordlist.txt
```

## Wordlist Example
Create a `wordlist.txt` file with common directories:
```
admin
administrator
api
assets
backup
config
dashboard
database
docs
downloads
files
images
includes
js
logs
media
private
public
secure
server
static
storage
styles
temp
test	uploads
user
wp-admin
wp-content
wp-includes
```

## Example Output
```
Found: http://example.com/admin
Found: http://example.com/wp-content
```

## License
This project is licensed under the MIT License.

## Contributing
Feel free to submit pull requests or open issues to improve this project.

## Author
Z3RO_COOL - [GitHub Profile](https://github.com/shabbirtoha/)

