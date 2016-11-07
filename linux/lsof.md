
## read lsof output
[http://www.ibm.com/developerworks/aix/library/au-lsof.html](http://www.ibm.com/developerworks/aix/library/au-lsof.html)

    COMMAND    PID   USER   FD   TYPE        DEVICE SIZE/OFF      NODE NAME
    sched        0   root  cwd   VDIR         136,8     1024         2 /
    init         1   root  cwd   VDIR         136,8     1024         2 /
    init         1   root  txt   VREG         136,8    49016      1655 /sbin/init
    init         1   root  txt   VREG         136,8    51084      3185 /lib/libuutil.so.1
    vi        2013   root    3u  VREG         136,8        0      8501 /var/tmp/ExXDaO7d
    ...

The Command, PID, and User columns represent the name of a process, process identifier (PID), and owner's name, respectively. The Device, SIZE/OFF, Node, and Name columns refer to the file itself, specifying the name of the disk, size of the file, inode (the file's identification on the disk), and actual name of the file. Depending on the flavor of UNIX, the size of the file might also be reported as the current position the application is reading in the file (offset). 

The FD and Type columns are the most cryptic and provide more information about how the file is being used. The FD column represents the file descriptor, which is how the application sees the file. The Type column gives more description about what form the file takes. Looking specifically at the file descriptor column, there are three different values represented in Listing 1. The cwd value refers to the application's current working directory, which is the directory that the application was started from, unless it has changed the directory itself. The txt files are program code, such as the application binary itself or a shared library, as in the case of the init program in the listing. Finally, a number refers to the application's file descriptor, which is an integer returned upon opening the file. In the final line of the output from Listing 1, you can see that the user is editing /var/tmp/ExXDaO7d with vi, with file descriptor 3. The u means the file has been opened in read/write mode, rather than read-only (r) or write-only (w). As a bit of helpful trivia, each application is initially opened with three file descriptors, 0 through 2, for the standard input, output, and error streams, respectively. As such, most opened files from the application start at FD 3.
Compared to the FD column, the Type column is more straightforward. Depending on the operating system, you find files and directories called REG and DIR (in the case of Solaris, VREG and VDIR). Other possible values are CHR and BLK for character and block devices, or UNIX, FIFO, and IPv4 for UNIX domain sockets, first-in-first-out (FIFO) queues, and Internet Protocol (IP) sockets.


[https://danielmiessler.com/study/lsof/](https://danielmiessler.com/study/lsof/)
> lsof is the sysadmin/security über-tool. I use it most for getting network connection related information from a system, but that’s just the beginning for this powerful and too-little-known application. The tool is aptly called lsof because it “lists open files“. And remember, in UNIX just about everything (including a network socket) is a file.

## key options
It’s important to understand a few key things about how lsof works. Most importantly, when you’re passing options to it, the default behavior is to OR the results. So if you are pulling a list of ports with -i and also a process list with -p you’re by default going to get both results.

Here are a few others like that to keep in mind:

default : without options, lsof lists all open files for active processes
grouping : it’s possible to group options, e.g. -abC, but you have to watch for which options take parameters
-a : AND the results (instead of OR)
-l : show the userID instead of the username in the output
-h : get help
-t : get process IDs only
-U : get the UNIX socket address
-F : the output is ready for another command, which can be formatted in various ways, e.g. -F pcfn (for process id, command name, file descriptor, and file name, with a null terminator)

## Getting Information About the Network

- Show all connections with -i 
`lsof -i`

- Show only TCP connections (works the same for UDP) 
`lsof -iTCP`

- Show networking related to a given port using 
`lsof -i :22`

- Show connections to a specific host
`lsof -i@172.16.12.5`

- Show connections based on the host and the port 
`lsof -i@172.16.12.5:22`

- Find listening ports
`lsof -i -sTCP:LISTEN`
`lsof -i | grep -i LISTEN`

- Find established connections
`lsof -i -sTCP:ESTABLISHED`

## User Information
You can also get information on various users and what they’re doing on the system, including their activity on the network, their interactions with files, etc.

- Show what a given user has open
`lsof -u daniel`

- Show what all users are doing except a certain user using 
`lsof -u ^daniel`

- Kill everything a given user is doing
`kill -9 `lsof -t -u daniel``

## Commands and Processes
- See what files and network connections a named command is using
`lsof -c syslog-ng`

- See what a given process ID has open
`lsof -p 10075`

- The -t option returns just a PID
`lsof -t -c Mail`

## Files and Directories
- Show everything interacting with a given directory
`lsof /var/log/messages/`

- Show everything interacting with a given file
`lsof /home/daniel/firewall_whitelist.txt`

## Advanced Usage
- Show me everything daniel is doing connected to 1.1.1.1
`lsof -u daniel -i @1.1.1.1`

- Using the -t and -c options together to HUP processes
`kill -HUP `lsof -t -c sshd``

- Show open connections with a port range
`lsof -i @fw.google.com:2150=2180`



