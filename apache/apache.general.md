
### check whether Apache is installed 
`dpkg --get-selections | grep apache`

It lists all installed packages that contain "apache" in their name. For example:
    apache2                                         install
    apache2-doc                                     install
    apache2-mpm-prefork                             install
    apache2-utils                                   install
    apache2.2-bin                                   install
    apache2.2-common                                install
    libapache2-mod-php5                             install
    libapache2-svn                                  install

Another approach, to find any running HTTP daemon on the default port would be:

`sudo lsof -nPi | grep ":80 (LISTEN)"`

Also, you can
`pt-cache policy apache2`

    apache2:
      Installed: (none)
      Candidate: 2.2.20-1ubuntu1.1
      Version table:
         2.2.20-1ubuntu1.1 0
            500 http://us.archive.ubuntu.com/ubuntu/ oneiric-updates/main amd64 Packages
            500 http://security.ubuntu.com/ubuntu/ oneiric-security/main amd64 Packages
         2.2.20-1ubuntu1 0
            500 http://us.archive.ubuntu.com/ubuntu/ oneiric/main amd64 Packages
    

[http://httpd.apache.org/docs/1.3/logs.html#accesslog](http://httpd.apache.org/docs/1.3/logs.html#accesslog)

## Apach error log
[http://httpd.apache.org/docs/1.3/logs.html](http://httpd.apache.org/docs/1.3/logs.html)
The error log is usually written to a file (typically *error_log* on unix systems and error.log on Windows and OS/2). On unix systems it is also possible to have the server send errors to *syslog* or *pipe them to a program*.

The format of the error log is relatively free-form and descriptive. But there is certain information that is contained in most error log entries. For example, here is a typical message.

`[Wed Oct 11 14:32:52 2000] [error] [client 127.0.0.1] client denied by server configuration: /export/home/live/ap/htdocs/test`

During testing, it is often useful to continuously monitor the error log for any problems. On unix systems, you can accomplish this using:
`tail -f error_log`

## access log
The server access log records all requests processed by the server. 

## 	

## Apache Core Features
[http://httpd.apache.org/docs/1.3/mod/core.html#errorlog](http://httpd.apache.org/docs/1.3/mod/core.html#errorlog)

### Directivies 
- AcceptFilter
- AcceptMutex
- AccessConfig
- AccessFileName
- AddDefaultCharset
- AddModule
- AllowOverride
- AuthName
- AuthDigestRealmSeed
- AuthType
- BindAddress
- BS2000Account
- CGICommandArgs
- ClearModuleList
- ContentDigest
- CoreDumpDirectory
- DefaultType
- <Directory>
- <DirectoryMatch>
- DocumentRoot
- EBCDICConvert
- EBCDICConvertByType
- EBCDICKludge
- EnableExceptionHook
- ErrorDocument
- ErrorLog
- FileETag
- <Files>
- <FilesMatch>
- Group
- HostnameLookups
- IdentityCheck
- <IfDefine>
- <IfModule>
- Include
- KeepAlive
- KeepAliveTimeout
- <Limit>
- <LimitExcept>
- LimitInternalRecursion
- LimitRequestBody
- LimitRequestFields
- LimitRequestFieldsize
- LimitRequestLine
- Listen
- ListenBacklog
- <Location>
- <LocationMatch>
- LockFile
- LogLevel
- MaxClients
- MaxKeepAliveRequests
- MaxRequestsPerChild
- MaxSpareServers
- MinSpareServers
- NameVirtualHost
- Options
- PidFile
- Port
- ProtocolReqCheck
- Require
- ResourceConfig
- RLimitCPU
- RLimitMEM
- RLimitNPROC
- Satisfy
- ScoreBoardFile
- ScriptInterpreterSource
- SendBufferSize
- ServerAdmin
- ServerAlias
- ServerName
- ServerPath
- ServerRoot
- ServerSignature
- ServerTokens
- ServerType
- ShmemUIDisUser
- StartServers
- ThreadsPerChild
- ThreadStackSize
- TimeOut
- TraceEnable
- UseCanonicalName
- User
- <VirtualHost>

