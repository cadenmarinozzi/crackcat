cd src

if !(go build -o ../bin/crackcat.exe ./crackcat.go); then
    $SHELL
fi