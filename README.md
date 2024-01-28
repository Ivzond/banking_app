**Banking App**

This is my simple banking application with server written on Go using gorilla/mux framework and client written in html/css.


**General information**

I used PostgreSQL as database for this project.
The database consists of three tables: accounts, transactions and users. It is assumed that each user can have multiple accounts.
I've also used gorm models in my server logic to make communication with the database easier and more convenient.

This app implements basic logic of bank app. You can register/login. Application provides jwt-authentication
and passwords are hashed and salted before inserting to the database.
On user page you can manage your accounts, make transactions and also get history of them.
I've tried to handle all basic errors and ensure safety while using application. 
Transactions use begin, rollback, commit logic and also have FOR UPDATE option when querying the database,
so they will lock the selected rows for update until the transaction is committed or rolled back.


I've tried to organize structure of project according to best practises and Go project-layout.
You can find main logic of application in fintech_app/internal/ folder.

**Technologies were used**

* Go
* HTML
* CSS
* JavaScript
* gorilla/mux
* gorm
* golang.org/x/crypto
* github.com/dgrijalva/jwt-go
* PostgreSQL
* Dockerfile
* Docker-compose


I'm happy to inform that I've added Dockerfile, docker-compose.yml and init.sql files so my application
can be easily run on any system that has a docker. I added saving data to a local database, 
the path to which can be changed in docker-compose.yml, so that data is not lost when the server is turned off.

Also, I've uploaded docker image on docker Hub.
Here's the link: https://hub.docker.com/repository/docker/ivzond/banking_app/general.
But you can just use "docker pull ivzond/banking_app:latest" command to get docker image.


**Instruction to run application with Docker**
1) If you don't want to rebuild docker image, make sure that the data from Postgresql in docker-compose matches yours
2) In the volumes column of docker-compose.yml (line 23 in my file), before the colon, change the path to the folder where your local database will be stored
3) Make sure that you have Docker, docker-compose and my docker image installed.
4) Next, from the folder with docker-compose.yml and the init folder where init.sql is located, enter the command "docker-compose up"
5) Server must work correctly, You can also use my small client interface to make use and testing more enjoyable. It's located in frontend folder

P.S. If you want to use your own settings on Postgres, you must change a bit database.go, docker-compose.yml and rebuild docker image manually.



**The future of the project**

I've made this project for experience in projecting RESTAPI service, trying to connect PostgreSQL and Go,
understanding of transactions and their features and much more. In addition, I've tried myself in frontend, 
even if it was quite simple. So now I will be developing this project rarely, in my free time, expanding the functionality
