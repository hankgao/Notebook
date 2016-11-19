[Way to avoid ssh connection timeout & freezing of GNOME Terminal](http://superuser.com/questions/98562/way-to-avoid-ssh-connection-timeout-freezing-of-gnome-terminal)

- Press Enter, ~, . one after the other to disconnect from a frozen session.
- change ~/.ssh/config to add: 
    Host *
      ServerAliveInterval 240