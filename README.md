# crackcat

## FAQ
### Q: Why does it say "Started cracking at ..." but it's taking forever?
a: This could be because of a few reasons. The most common is the "max_time" argument is too high and the dictionary/password list is huge, causing the program to freeze. To prevent this, lower the "max_time" argument to a lower value.

### Q: What platforms are supported
A: Linux, Windows, Mac, and any other mainstream platforms.

## Building
To build crackcat, you will need `go 1.16>=` installed. You can find details on how to do that here: `https://go.dev/doc/install`.

With go installed, download or clone the repository to any directory by running `git clone https://github.com/nekumelon/crackcat.git` in the directory.
Once the download has finished, run the following commands:
`
cd src
go build crackcat.go
`
This will build the crackcat code to an executable called `crackcat.exe`. 
You can run crackcat directly from the build directory by running `./crackcat.exe` or `crackcat.exe` (Depending on the console), or you can add it to the system PATH and run it from any command line. To add it to the system PATH on Windows, copy the path to the directory crackcat.exe is in, go to the Windows search and type `system environment variables`, click "Environment Variables" on the bottom right, in "User Variables", double click "Path" and then click "New". Paste the directory path and then click ok to exit out. (You will need to restart any command line instances you currently have open for them to refresh their PATH)
## Features
* Distributed multi threaded cracking
* Attack methods (Dictionary)
* Rules
* Large platform compatability 
* Attack modes (left-right, right-left)
* Dictionary/rules optimization
* Time limiting
* Quite a few Hashing algorithms
* Auto hash detecting
* Benchmarking

## TODO
See TODO for the TODO list. (lol)

## Troubleshooting

## License
License information is in LICENSE.

## Author
crackcat is made with ❤ by cadenmarinozzi.
