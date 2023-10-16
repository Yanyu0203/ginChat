# ginChat
* Implemented a instant messaging platform using Go and WebSocket, facilitating various functionalities including user registration, login, and message sending and reception, etc.  
* Enabled both private and group chatting, along with handling diverse message formats such as text, voice, images, and emoji.  
* Utilized the Gin and go-swagger for web API implementation and testing.  
* Employed channels and goroutines to enhance system concurrency and real-time responsiveness, ensuring a seamless user experience.  
* Leveraged MySQL for the storage of fundamental user and group information, with GORM utilized for efficient database manipulation.  
* Applied Redis for caching historical data, optimizing system performance and data retrieval efficiency.  

### Building Environment
go 1.20  
MySQL  
Redis  
If you want to build and run this application in your local environment. Please make sure you have environments mentioned above. And then you should update **app.yml** according your environment  
![image](https://github.com/Yanyu0203/ginChat/assets/132418583/973e267a-42b4-4979-8de7-9095ff0d9653)
And then run **build.bat** (for windows, **build.sh** for Linux) to build the application.
