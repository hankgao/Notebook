### 参考
- [vsftp虚拟用户配置](http://www.cnblogs.com/cnsanshao/p/3539047.html)
- [vsftpd Configuration Options](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/3/html/Reference_Guide/s1-ftp-vsftpd-conf.html)
- [http://www.cnblogs.com/JemBai/archive/2009/02/05/1384413.html](http://www.cnblogs.com/JemBai/archive/2009/02/05/1384413.html)

## 创建每个虚拟用户自己的配置文件
配置文件的路径是/etc/vsftpd/vsftpd.conf中的
user_config_dir=/etc/vsftpd/vuser_conf路径

在 /etc/vsftpd/vuser_conf/下面创建以用户名为名称的文件
内容如下

    local_root=/var/www(虚拟用户的根目录根据实际修改)
    write_enable=YES （可写）
    download_enable=YES
    anon_world_readable_only=NO
    anon_upload_enable=YES
    anon_mkdir_write_enable=YES
    anon_other_write_enable=YES
    local_umask=022


## 配置文件：
- /etc/vsftpd.conf          //主配置文件
- /etc/vsftpd.ftpusers      //被禁止登录FTP的用户文件
- /etc/vsftpd.user_list     //允许登录FTP的用户文件

[https://help.ubuntu.com/community/vsftpd](https://help.ubuntu.com/community/vsftpd)
## Virtual users with TLS/SSL/FTPS and a common upload directory - Complicated VSFTPD
Virtual users are users that do not exist on the system - they are not in /etc/passwd, do not have a home directory on the system, can not login but in vsftpd - or if they do exist, they can login in vsftpd with a non system password - security. 

You can set different definitions to each virtual user, granting to each of these users different permissions. If TLS/SSL/FTPS and virtual users are enabled, the level of security of your vsftpd server is increased: encrypted passwords, with passwords that are not used on the system, and users that can't access directly to their home directory (if you want). 

The following example is based and adapted on the example for virtual users in vsftpd site, on documentation and the very good examples in this forum that can be found here and here. Currently there is a restriction that with guest_enable enabled, local users also get mapped to guest_username. This is a polite way to say that if the default vsftpd PAM file is used, the system users will be guests too. To avoid confusions change the PAM file used by vsftpd to authenticate only virtual users, make all vsftpd users as virtual users and set their passwords, home and permissions based on this example. 

#### the workshop
- create the vitual user database
To create a "db4" format file to store usernames (another option here would be an apache htpasswd style file, not discussed), first create a plain text files with the usernames and password on alternating lines. For e.g. create user called "vivek" with password called "vivekpass" and sayali with password "sayalipass": 

`# mkdir /etc/vsftpd # if necessary`
`# cd /etc/vsftpd`
`# sudo vi vusers.txt`

Next, create the actual database file like this (may require the db_util package to be installed first): 

# db_load -T -t hash -f vusers.txt vsftpd-virtual-user.db 
# chmod 600 vsftpd-virtual-user.db # make it not global readable
# rm vusers.txt

- Configure VSFTPD for virtual user
Edit the vsftpd configuration file (/etc/vsftpd.conf). Add or correct the following configuration options, depending on if they're already listed somewhere in the file or not (or just add these all to the bottom): 

    anonymous_enable=NO
    local_enable=YES
    # Virtual users will use the same privileges as local users.
    # It will grant write access to virtual users. Virtual users will use the
    # same privileges as anonymous users, which tends to be more restrictive
    # (especially in terms of write access).
    virtual_use_local_privs=YES
    write_enable=YES    

    # Set the name of the PAM service vsftpd will use
    pam_service_name=vsftpd.virtual    

    # Activates virtual users
    guest_enable=YES    

    # Automatically generate a home directory for each virtual user, based on a template.
    # For example, if the home directory of the real user specified via guest_username is
    # /home/virtual/$USER, and user_sub_token is set to $USER, then when virtual user vivek
    # logs in, he will end up (usually chroot()'ed) in the directory /home/virtual/vivek.
    # This option also takes affect if local_root contains user_sub_token.
    user_sub_token=$USER    

    # Usually this is mapped to Apache virtual hosting docroot, so that
    # Users can upload files
    local_root=/home/vftp/$USER    

    # Chroot user and lock down to their home dirs
    chroot_local_user=YES    

    # Hide ids from user
    hide_ids=YES

Save and close the file. 

- Create a PAM File Which Uses Your New Database
The following PAM is used to authenticate users using your new database. Create /etc/pam.d/vsftpd.virtual: # sudo vi /etc/pam.d/vsftpd.virtual

- Append (or create with) the following:
#%PAM-1.0
auth       required     pam_userdb.so db=/etc/vsftpd/vsftpd-virtual-user
account    required     pam_userdb.so db=/etc/vsftpd/vsftpd-virtual-user
session    required     pam_loginuid.so

Create The Location Of The Files 

You need to set up the location of the files / dirs for the virtual users. Type the following command: # mkdir /home/vftp
# mkdir -p /home/vftp/{vivek,sayali}
# chown -R ftp:ftp /home/vftp


