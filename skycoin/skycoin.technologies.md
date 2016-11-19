
## techonlogies used in Skycoin
- electron
- angular 2.0
- qt
- DHT distributed hash table
- Web API
- Web service
- RPC (Remote Procedure Call)
- JSON-RPC
- JSON


## electron
[About electron](http://electron.atom.io/docs/tutorial/about/)
Electron is an open source library developed by GitHub for building cross-platform desktop applications with HTML, CSS, and JavaScript. Electron accomplishes this by combining Chromium and Node.js into a single runtime and apps can be packaged for Mac, Windows, and Linux.

** Electron enables you to create desktop applications with pure JavaScript by providing a runtime with rich native (operating system) APIs **. You could see it as a variant of the Node.js runtime that is focused on desktop applications instead of web servers.

Generally, an Electron app is structured like this:

your-app/
├── package.json
├── main.js
└── index.html

## GULP

## node.js

### install node.js on mac 

`brew install node`

## Qt
Qt is a cross-platform application framework that is widely used for developing application software that can be run on various software and hardware platforms with little or no change in the underlying codebase, while still being a native application with native capabilities and speed. 


## Web API
[How to Leverage an API for Conferencing](http://money.howstuffworks.com/business-communications/how-to-leverage-an-api-for-conferencing.htm)

An application-programming interface (API) is a set of programming instructions and standards for accessing a Web-based software application or Web tool. A software company releases its API to the public so that other software developers can design products that are powered by its service.

For example, Amazon.com released its API so that Web site developers could more easily access Amazon's product information. Using the Amazon API, a third party Web site can post direct links to Amazon products with updated prices and an option to "buy now."

**An API is a software-to-software interface**, not a user interface. With APIs, applications talk to each other without any user knowledge or intervention. When you buy movie tickets online and enter your credit card information, the movie ticket Web site uses an API to send your credit card information to a remote application that verifies whether your information is correct. Once payment is confirmed, the remote application sends a response back to the movie ticket Web site saying it's OK to issue the tickets.

As a user, you only see one interface -- the movie ticket Web site -- but behind the scenes, many applications are working together using APIs. This type of integration is called seamless, since the user never notices when software functions are handed from one application to another [source: TConsult, Inc.]

**An API resembles Software as a Service (SaaS), since software developers don't have to start from scratch every time they write a program**. Instead of building one core application that tries to do everything -- e-mail, billing, tracking, etcetera -- the same application can contract out certain responsibilities to remote software that does it better.­

Let's use the same example of Web conferencing from before. Web conferencing is SaaS since it can be accessed on-demand using nothing but a Web site. With a conferencing API, that same on-demand service can be integrated into another Web-based software application, like an instant messaging program or a Web calendar.

The user can schedule a Web conference in his Web calendar program and then click a link within the same program to launch the conference. The calendar program doesn't host or run the conference itself. It uses a conferencing API to communicate behind the scenes with the remote Web conferencing service and seamlessly delivers that functionality to the user.

Now we'll explain some of the technology that makes conferencing APIs work.

A conferencing API -- or any API for that matter -- allows a software application to communicate with a remote application over the Internet through a series of calls [source: TConsult, Inc.] An API is, by definition, an interface, something that defines the way in which two entities communicate [source: Thom Robbins.net Weblog].

With APIs, the calls back and forth between applications are managed through something called Web services. Web services are a collection of technological standards and protocols, including XML (Extensible Markup Language), the programming language by which applications communicate over the Internet.

he API itself is a chunk of software code written as a series of XML messages. Each XML message corresponds to a different function of the remote service. For example, in a conferencing API, there are XML messages that correspond with each element required to schedule a new Web conference. Those elements include the conference time, the organizer's name and contact information, who's invited to the conference, the duration of the conference, et cetera.

Exactly how does a software developer leverage a conferencing API? He programs new or existing software to generate the right XML messages to tap the power of remote applications. Let's take conference scheduling, for example. With the right code, he could build conference-scheduling functionality into the company's existing e-mail system. Or, perhaps he could develop an entirely new instant messaging application that has one-click instant audio conferences.

Companies who release their API often do so as part of a larger software development kit (SDK) that includes the API, programming tools and other instructional documents to make the developer's job easier.

APIs and Web services are completely invisible to Web site surfers and software users. Their job is to run silently in the background, providing a way for applications to work with each other to get the user the information or functionality he needs.

Along with XML, the following technological standards, protocols and programming languages are what make Web services work:

- SOAP (Simple Object Access Protocol): SOAP is responsible for encoding XML messages so that they can be received and understood by any operating system over any type of network protocol.

- UDDI (Universal Description, Discovery and Integration): Described as a "yellow pages for the Internet," UDDI is an XML-based directory that allows businesses to list themselves, find each other and collaborate using Web services.

- WSDL (Web Services Description Language): WDSL is the SOAP of the UDDI (enough acronyms for you?). Basically, WDSL is the XML-based language that businesses use to describe their services in the UDDI.

Now let's look at some examples of how software developers and businesses can leverage a conferencing API.

## Web service
RPC: Remote procedure call (RPC) is an Inter-process communication technology that allows a computer program to cause a subroutine or procedure to execute in another address space (commonly on another computer on a shared network) without the programmer explicitly coding the details for this remote interaction.

Web Service: Web services are typically application programming interfaces (API) or web APIs that are accessed via Hypertext Transfer Protocol and executed on a remote system hosting the requested services. Web services tend to fall into one of two camps: Big Web Services[1] and RESTful Web Services.

A Web service is a service offered by an electronic device to another electronic device, communicating with each other via the World Wide Web. In a Web service, Web technology such as HTTP, originally designed for human-to-machine communication, is utilized for machine-to-machine communication, more specifically for transferring machine readable file formats such as **XML** and **JSON**. In practice, the Web service typically provides an object-oriented Web-based interface to a database server, utilized for example by another Web server, or by a mobile application, that provides a user interface to the end user. Another common application offered to the end user may be a mashup, where a Web server consumes several Web services at different machines, and compiles the content into one user interface.

The W3C defines a Web service generally as:

    A Web service is a software system designed to support interoperable machine-to-machine interaction over a network.

## JSON-RPC
JSON-RPC is a remote procedure call protocol encoded in JSON. It is a very simple protocol (and very similar to XML-RPC), defining only a handful of data types and commands. JSON-RPC allows for notifications (data sent to the server that does not require a response) and for multiple calls to be sent to the server which may be answered out of order.

JSON-RPC works by sending a request to a server implementing this protocol. The client in that case is typically software intending to call a single method of a remote system. Multiple input parameters can be passed to the remote method as an array or object, whereas the method itself can return multiple output data as well. (This depends on the implemented version.)

All transfer types are single objects, serialized using JSON.[1] A request is a call to a specific method provided by a remote system. It must contain three certain properties:

- method - A String with the name of the method to be invoked.
- params - An Object or Array of values to be passed as parameters to the defined method.
- id - A value of any type used to match the response with the request that it is replying to.

The receiver of the request must reply with a valid response to all received requests. A response must contain the properties mentioned below.

- result - The data returned by the invoked method. If an error occurred while invoking the method, this value must be null.
- error - A specified error code if there was an error invoking the method, otherwise null.
- id - The id of the request it is responding to.

Since there are situations where no response is needed or even desired, notifications were introduced. A notification is similar to a request except for the id, which is not needed because no response will be returned. In this case the id property should be omitted (Version 2.0) or be null (Version 1.0).

## JSON
In computing, JSON (sometimes JavaScript Object Notation) is an open-standard format that uses human-readable text to transmit data objects consisting of attribute–value pairs. It is the most common data format used for asynchronous browser/server communication, largely replacing XML which is used by AJAX.





