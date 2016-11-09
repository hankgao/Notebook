
### delete a user
[Linux: Delete / Remove User Account](http://www.cyberciti.biz/faq/linux-remove-user-command/)
`userdel username`
`userdel [options] username`
`userdel -r username`

### disable a user
`passwd -l vivek`

### umask
In computing, umask is a command that determines the settings of a mask that controls how file permissions are set for newly created files. It also may refer to a function that sets the mask, or it may refer to the mask itself, which is formally known as the file mode creation mask. The mask is a grouping of bits, each of which restricts how its corresponding permission is set for newly created files. The bits in the mask may be changed by invoking the umask command.

In UNIX, each file has a set of attributes which control who can read, write or execute it. When a program creates a file, UNIX requires that the file permissions be set to an initial setting. The mask restricts permission settings. If the mask has a bit set to "1", it means the corresponding initial file permission will be disabled. A bit set to "0" in the mask means that the corresponding permission will be determined by the program and the system. In other words, the mask acts as a last-stage filter that strips away permissions as a file is created; each bit that is set to a "1" strips away its corresponding permission. Permissions may be changed later by users and programs using chmod.


### PAM - Linux Pluggable Authentication Module 
Linux Pluggable Authentication Modules (PAM) provide dynamic authentication support for applications and services in a Linux. 

Linux-PAM separates the tasks of authentication into four independent management groups:

- account modules check that the specified account is a valid authentication target under current conditions. Thisdelete may include conditions like account expiration, time of day, and that the user has access to the requested service.

- authentication modules verify the user's identity, for example by requesting and checking a password or other secret. They may also pass authentication information on to other systems like a keyring.

- password modules are responsible for updating passwords, and are generally coupled to modules 
employed in the authentication step. They may also be used to enforce strong passwords.

- session modules define actions that are performed at the beginning and end of sessions. A session 
starts after the user has successfully authenticated.

#### PAM configuration files 
[PAM Configuration Files](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/6/html/Managing_Smart_Cards/PAM_Configuration_Files.html)

Each PAM-aware application or service has a file in the /etc/pam.d/ directory. Each file in this directory has the same name as the service to which it controls access.

** PAM Configuration File Format**
`module_interface     control_flag     module_name module_arguments`

** PAM module interfaces **
- auth — This module interface authenticates use. For example, it requests and verifies the validity of a password. Modules with this interface can also set credentials, such as group memberships or Kerberos tickets.

- account — This module interface verifies that access is allowed. For example, it checks if a user account has expired or if a user is allowed to log in at a particular time of day.

- password — This module interface is used for changing user passwords.

- session — This module interface configures and manages user sessions. Modules with this interface can also perform additional tasks that are needed to allow access, like mounting a user's home directory and making the user's mailbox available.

* Note: An individual module can provide any or all module interfaces. For instance, pam_unix.so provides all four module interfaces. *

** AM Control Flags
All PAM modules generate a success or failure result when called. Control flags tell PAM what do with the result. Modules can be stacked in a particular order, and the control flags determine how important the success or failure of a particular module is to the overall goal of authenticating the user to the service.

- required — The module result must be successful for authentication to continue. If the test fails at this point, the user is not notified until the results of all module tests that reference that interface are complete.

- requisite — The module result must be successful for authentication to continue. However, if a test fails at this point, the user is notified immediately with a message reflecting the first failed required or requisite module test.

- sufficient — The module result is ignored if it fails. However, if the result of a module flagged sufficient is successful and no previous modules flagged required have failed, then no other results are required and the user is authenticated to the service.

- optional — The module result is ignored. A module flagged as optional only becomes necessary for 
successful authentication when no other modules reference the interface.

- include — Unlike the other controls, this does not relate to how the module result is handled. This flag pulls in all lines in the configuration file which match the given parameter and appends them as an argument to the module.

Note： The order in which required modules are called is not critical. Only the sufficient and requisite control flags cause order to become important.



#### further readings
- [How PAM Works: The Basics](http://www.informit.com/articles/article.aspx?p=20968&seqNum=3)
- [The Linux-PAM System Administrators' Guide](http://linux-pam.org/Linux-PAM-html/Linux-PAM_SAG.html)
- [How PAM works](http://www.tuxradar.com/content/how-pam-works)

## 

- Show All Members of a Group
> The `/etc/group` file is a text file that defines the groups on the Linux and Unix based systems. You can simply query this file to find and list all members of a group.
 
 + list all members in group-name: `members group-name`
 + list all members in group-name: `lid -g group-name`
 + list all groups a member belong to: `lid member`

 Note that you need to install members and lid by running the followng commands: 
 > `apt-get install members`
 > `apt-get install libuser`

 - See Which Groups Your Linux User Belongs To
   + `groups user-name`
   + `lid user-name`


- add user to sudoers
  + option 1: edit /etc/sudoers.tmp
    1. Open a Root Terminal and type visudo (to access and edit the list)
    2. Using the up/down arrows, navigate to the bottom of the sudoers file that is now displayed in the terminal
    3. Just under the line that looks like the following:
    `root ALL=(ALL) ALL`
    Add the following (replacing user with your actual username):
    `user ALL=(ALL) ALL`
    4. Now press Ctrl+X and press Y when promted to save
   + option 2: Add a existing user to existing group
   `# usermod -a -G sudoer tony`

- change ownership of a file/directory
`chown -R user:group directory/`

- delete a folder 
rmdir foldername
rmdir -p dir1/dir2/dir3
rm -rf /path/to/dir

- cd into a folder 
if you get the following message when cd-ing into a directoy 
> permission denied 

Try the following: 
`sudo su`
`cd <folder> `

### chmod change file permisions
chmod commands take the form: 
`chmod options permissions filename`

Note: only the owner or privilaged user can use chmod to change file access permissions. 

there are three groups of permssions, one for the onwer, one for the owner's group, one for others, they are represented by u, g, o, respetively. 

there are three opcodes: + meaning add permissions, - deleting permissins, = allocate permissions

permissions: read (r), write(w), execute(x)

examples:
`chmod u+x file`
`chmod u=rwx,g=rx,o=x file `
 `chmod 444 file`

 Note that permission can use 3 octol digits to represent, the first one for owner, the 2dn for group, the last one for o, where 4 means read, 2 means write, 1 means execution.

options:
- c only print files who have been changed
- f --silent, --quiet 
--help
-R --recursive 
--reference=filename 
-v --verbose
--version

## init.d and init
[http://askubuntu.com/questions/5039/what-is-the-difference-between-etc-init-and-etc-init-d](http://askubuntu.com/questions/5039/what-is-the-difference-between-etc-init-and-etc-init-d)

`/etc/init.d` contains scripts used by the System V init tools (SysVinit). This is the traditional service management package for Linux, containing the init program (the first process that is run when the kernel has finished initializing¹) as well as some infrastructure to start and stop services and configure them. Specifically, files in /etc/init.d are shell scripts that respond to start, stop, restart, and (when supported) reload commands to manage a particular service. These scripts can be invoked directly or (most commonly) via some other trigger (typically the presence of a symbolic link in /etc/rc?.d/).

/etc/init contains configuration files used by Upstart. Upstart is a young service management package championed by Ubuntu. Files in /etc/init are configuration files telling Upstart how and when to start, stop, reload the configuration, or query the status of a service. As of lucid, Ubuntu is transitioning from SysVinit to Upstart, which explains why many services come with SysVinit scripts even though Upstart configuration files are preferred. In fact, the SysVinit scripts are processed by a compatibility layer in Upstart.

.d in directory names typically indicates a directory containing many configuration files or scripts for a particular situation (e.g. /etc/apt/sources.list.d contains files that are concatenated to make a virtual sources.list; /etc/network/if-up.d contains scripts that are executed when a network interface is activated). This structure is usually used when each entry in the directory is provided by a different source, so that each package can deposit its own plug-in without having to parse a single configuration file to reference itself. In this case, it just happens that “init” is a logical name for the directory, SysVinit came first and used init.d, and Upstart used plain init for a directory with a similar purpose (it would have been more “mainstream”, and perhaps less arrogant, if they'd used /etc/upstart.d instead).






