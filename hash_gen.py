import hashlib, os, sys

with open('crackcat.conf') as file:
    version = file.read().split("\n")[0];

hash = '';
lines = 0;

for path, folders, files in os.walk('./' + version):
    for fileName in files:
        filePath = os.path.join(path, fileName);

        with open(filePath, 'rb') as file: 
            fileContents = file.read();

            if (fileName.find('.go') != -1):
                lines += len(fileContents.splitlines());

            hash += hashlib.sha256(bytes(fileContents)).hexdigest();

hash = hashlib.sha256(bytes(hash, encoding = 'utf-8')).hexdigest();
print(hash);
print(lines);

if (len(sys.argv) > 1):
    with open('HASHES', 'r') as file:
        hashesContent = file.read();

    with open('HASHES', 'w') as file:
        file.write(hashesContent + f'[{version}] {hash}\n');