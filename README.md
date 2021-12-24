# crackcat
# License
crackcat is licensed under the MIT license. This permits anyone to use, copy, merge, distribute and sell this software with proper credit to the original developers. Check the "Crediting" section for more information.

# Building
First install go from https://go.dev/doc/install. Specific installation instructions are located on the page.

To build crackcat from the source files, download the source files by clicking the green "Code" button and then "Download ZIP" or clone it using git by running the command: `git clone https://github.com/hashcat/hashcat.git` in the building directory. Then, build it by running a command line from the build directory (The same directory the source files are located in) and then running the following command: `go build crackcat.go`. After the build process is done, an executable named "crackcat.exe" will be created in the build directory. Run it directly from the exe or add it to the systems %PATH% and run it from a command line.

# Features
* Multi threading and automatic thread detection
* Dictionary attacks
* Rule based attacks
* Multiple attack modes
* Dictionary optimization
* Time limiting
* Lots of hashing algorithms
* Username password pairs
* Benchmarking
* Hash auto detecting

# Usage
Run crackcat either from the executable directly or from the system %PATH%.. If running from the system %PATH%, use any command line and run the command "crackcat".

Add arguments by putting a "-" or "--" infront of the argument name and then a space connected to the argument value. For example: `crackcat -crack_mode font-back` or `crackcat --benchmark`
  
# Todo
* GPU assisted attacks

# Crediting
To redistribute this code, credit the author in any of the main source files or in the program itself.

# Author
crackcat was created with ❤ by nekumelon.