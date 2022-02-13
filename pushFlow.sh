python header.py
git add *

if !(git commit -m $1); then
    $SHELL
fi

if !(git push -u origin $2); then
    $SHELL
fi