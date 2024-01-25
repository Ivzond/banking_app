**Banking App**

This is my simple banking application with server written on Go using gorilla/mux framework and client written in html/css.


**General information**

I used PostgreSQL as database for this project.
The database consists of three tables: accounts, transactions and users. It is assumed that each user can have multiple accounts.
I've also used gorm models in my server logic to make communication with the database easier and more convenient.

This app implements basic logic of bank app. You can register/login. Application privides jwt-authentication and passswords are hashed and salted before inserting to the database.
On user page you can manage your accounts, make transactions and also get history of them.
I've tried to handle all basic errors and ensure safety while using application. Transactions use begin, rollback, commit logic.


I've tried to orginize structure of project according to best practises and Go project-layout.
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

**The future of the project**

Later, I plan to create a Docker container with app, so I can easily transport and share it and also run using Docker Compose.
I also want to expand the functionality of the application.
