# ginChat
* Implemented a instant messaging platform using Go and WebSocket, facilitating various functionalities including user registration, login, adding contacts, message sending and reception, etc.  
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
And then run **build.bat** (for windows, **build.sh** for Linux) to build the application. You can access to the system in **localhost:8081**

### Functions

#### User Registration
![image](https://github.com/Yanyu0203/ginChat/assets/132418583/376eaf35-76e6-4ae4-81a3-a5472c2a3e27)

#### User Login
![image](https://github.com/Yanyu0203/ginChat/assets/132418583/35bbf4b2-b32b-45be-b2bd-21916146c717)

#### Contacts List
![image](https://github.com/Yanyu0203/ginChat/assets/132418583/ccf2b45f-64a6-4de9-a19b-1572c0c7cae3)

#### Group Chat List
![image](https://github.com/Yanyu0203/ginChat/assets/132418583/27ef42fa-2169-484a-bfd5-4980f5dc4cee)

#### Chat Page
![image](https://github.com/Yanyu0203/ginChat/assets/132418583/72d18486-487e-4516-bba4-77748c7d000a)

### Test
Tools: JMeter  
Test machine: personal laptop(AMD R7 4800H)  
I first tested in normal case with 300 threads for 10 loops, achieving 5000+ TPS. The setting and results are as followed  
![set2](https://github.com/Yanyu0203/ginChat/assets/132418583/85e66ea6-12c6-44a4-9b16-f0579774781e)  
![test1](https://github.com/Yanyu0203/ginChat/assets/132418583/abd13d19-538d-4503-8fa1-b58043e263aa)  
Then I tested 300 user visited at the same time. The setting and results are as followed  
![set1](https://github.com/Yanyu0203/ginChat/assets/132418583/3e1c8586-4d59-4ea5-8c2e-49984595b47e)  
![test2](https://github.com/Yanyu0203/ginChat/assets/132418583/0a9f2b8f-cfc7-4282-8ca9-8a48ccb0efa9)


[Reference](https://www.bilibili.com/video/BV1rK4y1w7JB/?spm_id_from=333.999.0.0)
