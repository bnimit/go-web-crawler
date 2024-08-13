# Go-Web-Crawler

### <u>Introduction</u>
This is a small web crawler application that has been written in `Go`.

The objective of this mall command line application is parse a list of URLs, crawl them using Go's parallel execution power and return a JSON object containing the top 10 words with the highest counts.

### <u>Requirements</u>
To compile & run this application you would require:
- Go binaries installed in your local machine, preferably the latest ones along with installation instructions and can be found at the following location:
  https://go.dev/doc/install

- Download / Clone this repository to a local folder, extract its contents and change directory into the location.

### <u>How to Build & Run</u>
- Once you have the repository download and extracted into a local folder:
- Open the `terminal` or any other command line application on your machine.

-  you can run the application using
  `go run main.go -wordFile="custom_words.txt" -urlFile="custom_urls.txt"`

- If the command line arguments are not provided for the file paths it defaults to local text file present within the repository.

- You could also build an executable of the code by using the build command as:
  `go build -o web-crawler` 