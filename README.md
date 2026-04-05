# GoVault
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://github.com/Scryv/GoVault/blob/main/LICENSE)

**GoVault** will be a Self Hosted application for storing passwords. It will be made in Golang and Sqlite and the liberaries Gorm and later on Gin with a web interface and uses SHA-512+salting and aes encryption (will change out sha-512 for argon2 later on). 
> Right now i am focusing on the Cobra CLI version Gin with frontend comes later.
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
      │                       │ add.go    │
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
Database folder gets stored in /home/user/

## How to install
1. `git clone https://github.com/Scryv/GoVault.git`
2. `cd GoVault`
3. `chmod +x install.sh`
4. `sudo ./install.sh`
5. `GoVault`<br>

And enjoy :)

## Features

### Info about tool `GoVault`
will just display info about tool and all the commands there are available with
a tiny explanation about each command

### Create `GoVault create`
will ask for an **username** and a **password** and will take the password generate 
a **random** salt and hash the password together with that salt then save it to a local **SQLite** database

### add `GoVault add`
Will ask for **Username** and your accounts **password** after auth it will ask what option so 
1. Username-Password
2. Email-Password
3. Email-Password-Username<br>

you will fill those in and it will encrypt it and store it in your user his db

### govault `GoVault vault`
Will ask for a **login** and **password** then will search the **SQLite** database for your **username** then take its hash and hash your **password** with it and if its a match it will let you log in
and display the other database with stored **accounts** and **passwords** that will be unlocked by your **masterPasswd**


## Working on adding

#### Logs
#### Maybe TUI with bubble tea
