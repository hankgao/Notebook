## git configuration
- there are three configuraiton files 
  + one under each repository: .git/config
  > `git config -file user.name hankgao`
  + another one is located in your $HOME directory: gitconfig
  > `git config -global user.name hankgal`
  + The next config file is your system config file, gitconfig. This contains your system configuration settings used in running the git application.  You can add to your etc/gitconfig file with the --system flag like this: 
  `git config --system core.autocrlf true`

 Note that 
 > Git uses these in order when looking for parameters.  First the local repository .git/config, then the global repository .gitconfig, and then the system repository etc/gitconfig.

