### check tomcat version

`http://localhost:8080/host-manager/html`

## Logging in Tomcat

### Introduction
The internal logging for Apache Tomcat uses JULI, a packaged renamed fork of Apache Commons Logging that, by default, is hard-coded to use the java.util.logging framework. 

A web application running on Apache Tomcat can:

- Use any logging framework of its choice.
- Use system logging API, java.util.logging.
- Use the logging API provided by the Java Servlets specification, javax.servlet.ServletContext.log(...)

#### Console
When running Tomcat on unixes, the console output is usually redirected to the file named catalina.out(/var/log/tomcat8/catalina.out). The name is configurable using an environment variable. (See the startup scripts). Whatever is written to System.err/out will be caught into that file. That may include:

- Uncaught exceptions printed by java.lang.ThreadGroup.uncaughtException(..)
- Thread dumps, if you requested them via a system signal

[https://wiki.apache.org/tomcat/FAQ/Logging#Q1](https://wiki.apache.org/tomcat/FAQ/Logging#Q1)
### How should I log in my own webapps?

While you can use System.out and System.err to log, we strongly recommend using a toolkit like Log4J or JDK 1.4's java.util.logging package. With these toolkits, you have significantly more functionality. For example, sending emails, logging to a database, controlling at runtime the logging level of different classes, inspecting the logs with a graphical viewer, etc. 

We also recommend that you separate your logging from Tomcat's internal logging. That means you should bundle your logging toolkit with your webapp. If you are using Log4J, for example, place the Log4J jar in the WEB-INF/lib directory of your webapp and the Log4J configuration file in the WEB-INF/classes directory of your webapp. This way different web applications can have different logging configurations and you don't need to worry about them interfering with each other. 

### Where does System.out go?

System.out and System.err are both redirected to CATALINA_BASE/logs/catalina.out when using Tomcat's startup scripts (bin/startup.sh/.bat or bin/catalina.sh/.bat). Any code that writes to System.out or System.err will end up writing to that file. 

If your webapp uses System.out and/or System.err a lot, you can suppress this via the 'swallowOutput' attribute in your <Context> configuration element and send to different log files (configured elsewhere: see the documentation for configuring logging). 