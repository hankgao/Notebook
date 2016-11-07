## install golang

### Ubuntu 
version: 

1. get golang binary 
  + you can download from here: 
  []()
  + or if you have a local copy, you can use scp to copy it to the target machine
  `scp golangversion hank@www.mzworld.cn:/home/hank`
2. unzip the binary 
> `tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz` 

3. Add go to your path `export PATH=$PATH:/usr/local/go/bin`
4. Add GOPATH: `export GOPATH=~/Goworkspace`
5. Add bin to your path: `export PATH=$PATH:$GOPATH/bin`

- use gvm
[http://stackoverflow.com/questions/17480044/how-to-install-the-current-version-of-go-in-ubuntu-precise](http://stackoverflow.com/questions/17480044/how-to-install-the-current-version-of-go-in-ubuntu-precise)

> I like to use GVM for managing my Go versions in my Ubuntu box. Pretty simple to use, and if you're familiar with RVM, it's a nobrainer. It allows you to have multiple versions of Go installed in your system and switch between whichever version you want at any point in time.

Install GVM with:

> `sudo apt-get install bison mercurial
> `bash < <(curl -LSs 'https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer') . "$HOME/.gvm/scripts/gvm"`

and then it's as easy as doing this:

    gvm install go1.1.1
    gvm use go1.1.1 --default

The default flag at the end of the second command will set go1.1.1 to be your default Go version whenever you start a new terminal session.