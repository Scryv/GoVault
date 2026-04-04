# GoVault
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://github.com/Scryv/GoVault/blob/main/LICENSE)

**GoVault** will be a Self Hosted application for storing passwords. It will be made in Golang and Sqlite and the liberaries Gorm and Gin. 
> Right now i am focusing on the Cobra CLI version Gin frontend comes later.
<img src="https://www.sdgscryv.xyz/images/images/GoVault.jpg" width="600px">

---
## Structure
<pre>
          ┌─────────────────────────┐
          │        main.go          │
          └───────────┬─────────────┘
                     │
                     ▼
           ┌─────────────────────────┐
           │        cmd/root.go      │
           └─────────┬───────────────┘
                     │
      ┌──────────────┴────────────────┐
      │                               │
      ▼                               ▼
┌──────────────┐                ┌───────────────┐
│ create.go    │                │ govault.go    │
│ (Add new     │                │ (Login &      │
│ user)        │                │view passwords)│
└─────┬────────┘                └──────┬────────┘
      │                                 │
      │                                 ▼
      │                       ┌────────────────┐
      │                       │ addPaswd.go    │
      │                       │ (Add accounts/ │
      │                       │ passwords)     │
      │                       └──────┬─────────┘
      │                              │
      ▼                              ▼
┌───────────────────────────┐  ┌─────────────────────────────┐
│ database.go               │  │ database.go                 │
│ - VaultDB: stores users   │  │ - UserDB: per-user DB       │
│ - UserData/Data structs   │  │ - Data/UserData structs     │
│ - Functions: initDB,      │  │ - Functions: AddData,       │
│   initUserDB, createPost, │  │   encrypt/decrypt, hash     │
│   getUser, doPasswdMatch  │  │   password                  │
│ - Hashing & Salting       │  │ - AES Encryption/Decryption │
└───────────────────────────┘  └─────────────────────────────┘
</pre>
## Features
### Create `go run main.go create`
will ask for an **username** and a **password** and will take the password generate 
a **random** salt and hash the password together with that salt then save it to a local **SQLite** database

### addPswd `go run main.go addPswd`
Will ask for **Username** and your accounts **password** after auth it will ask what option so 
1. Username-Password
2. Email-Password
3. Email-Password-Username<br>

you will fill those in and it will encrypt it and store it in your user his db

### govault `go run main.go govault`
Will ask for a **login** and **password** then will search the **SQLite** database for your **username** then take its hash and hash your **password** with it and if its a match it will let you log in
and display the other database with stored **accounts** and **passwords** that will be unlocked by your **masterPasswd**


## Working on adding

#### Logs
#### Maybe TUI with bubble tea
