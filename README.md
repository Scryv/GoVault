# GoVault
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://github.com/Scryv/GoVault/blob/main/LICENSE)

**GoVault** will be a Self Hosted application for storing passwords. It will be made in Golang and Sqlite and the liberaries Gorm and Gin. 
> Right now i am focusing on the Cobra CLI version Gin frontend comes later.
<img src="https://www.sdgscryv.xyz/images/images/GoVault.jpg" width="600px">

---

## Features
### Create `go run main.go create`
will ask for an **username** and a **password** and will take the password generate 
a **random** salt and hash the password together with that salt then save it to a local **SQLite** database

### GoVault `go run main.go govault`
Will ask for a **login** and **password** then will search the **SQLite** database for your **username** then take its hash and hash your **password** with it and if its a match it will let you log in
and display the other database with stored **accounts** and **passwords** that will be unlocked by your **masterPasswd**

### Database structure
<pre>
GoVaultDB/<br/>
├─ Users/<br/>
│  ├─ jens.db<br/>
│  ├─ lazer.db<br/>
│  ├─ admin.db<br/>
├─ users.db<br/>
</pre>
## Working on adding

#### Adding master password

#### Password/Email account adder to DB
