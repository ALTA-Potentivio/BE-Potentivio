# Potentivio apps

![Logo](https://github.com/ALTA-Potentivio/BE-Potentivio/blob/main/logo.png)

Readme dalam bahasa indonesia [klik disini](https://github.com/ALTA-Potentivio/BE-Potentivio/blob/main/README.md).

Potentivio-apps is a platform used for cafe owners to find artists to appear on the cafe stage. From the artist's point of view, they can accept or reject work from the cafe. Artists can also submit requests to appear on the selected cafe stage.


## Feature as cafe

- Register
- Login User
- Edit User
- Cafe users can search for artists by category, genre, location and sort according to the highest or lowest price
- Cafe users can hire artists according to the time specified
- Cafe users can cancel artists who have been booked with valid reasons
- Cafe users can make payments through the integrated Xendit
- Cafe users can give ratings and comments when the event is done

## Featur as artist

- Register
- Login User
- Edit User
- User artists will get information if a cafe invites them
- User artist can reject or accept offers
- User artist can cancel the job that has been received

## Open APIs

For Open API, see more [click here](https://app.swaggerhub.com/apis-docs/satriacening/project_group3/1.0.0#/)


## Run Local

Cloning project

```bash
  $ git clone git@github.com:ALTA-BE7-SATRIA-PUTRA/group_project3.git
```

Go to the project directory

```bash
  $ cd ~/project
```
Create new `database`

Create a file with the name in the project root folder `.env` with the following format. Adjust the configuration on the local computer

```bash
  export APP_PORT=""
  export JWT_SECRET="S3CR3T"
  export DB_PORT="3306"
  export DB_DRIVER="mysql"
  export DB_NAME=""
  export DB_ADDRESS="127.0.0.1"
  export DB_USERNAME=""
  export DB_PASSWORD=""
```

Run the app

```bash
  $ go run main.go
```


## Authors

- [@satriacening](https://github.com/satriacening)
- [@usamaha07](https://github.com/usamaha07)
- [@abbiyudha](https://github.com/abbiyudha)

 

