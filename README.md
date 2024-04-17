This is a small and portable version of ssh-keyscan intended to be used on embedded platforms where ssh-keyscan is not installed.
Note that not all regular options are available.

# Usage

    ssh-keyscan [-p port] <host>

# Install

    git clone https://github.com/jetibest/ssh-keyscan
    cd ssh-keyscan
    go build
    # to cross-compile, use environment variables such as: GOOS=linux GOARCH=mips go build
    cp ssh-keyscan /usr/local/bin/ssh-keyscan

# Uninstall

    rm /usr/local/bin/ssh-keyscan

