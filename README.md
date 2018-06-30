# Go Rest API Starter 

## Structure Folder


		├── application
		│   ├── controller
		│   │   ├── ApiController.go
		│   │   └── UserController.go
		│   ├── lib
		│   │   ├── BaseModel.go
		│   │   ├── Middleware.go
		│   │   ├── Request.go
		│   │   └── Response.go
		│   ├── middlewares
		│   │   └── Cors.go
		│   ├── router.go
		│   └── server.go
		├── config
		│   └── database.go
		├── main.go
		└── models
			└── User.go

		6 directories, 12 files

## Auto Migration Models
	
	//in main.go
	func MigrateDatabase() {
		fmt.Println(":: Migration Databases .....")
		db := config.GetDatabaseConnection()
		db.AutoMigrate(&models.User{}) // model users
		b.AutoMigrate(&models.Profile{}) // model profile
		..
		..
		fmt.Println(":: Migration Databases Done")
	}


## .env
Create .env file 
example file .env

	ENVIRONMENT="development"
	
	API_TITLE="TEST"
	API_VERSION="0.0.0"
	API_PORT=3000

	DB_DIALECT="mysql"
	DB_CONNECTION="username:password@tcp(localhost:3306)/databasename?charset=utf8&parseTime=True"


## run

	[16:21:15-golang-rest-API-rwx-]$ go run main.go 
	:: Checking databse connection ..... 
	:: Database Connected
	:: Migration Databases .....
	:: Migration Databases Done
	:: APP running on port 3000

	
## Simple CRUD endpoint

### 1. Create

[POST] localhost:3000/v1/user/

	
	{
		"FirstName" : "rahmat wahyu",
		"LastName" : "Hadi",
		"Email" : "rhmt@gmail.com"
	}

### 2. GET ALL

[GET] localhost:3000/v1/user/


### 3. GET BY ID

[GET] localhost:3000/v1/user/1


### 4. UPDATE BY ID

[PUT] localhost:3000/v1/user/1
	
### 5. DELETE BY ID

[DELETE] localhost:3000/v1/user/1


## Contact me

Rahmat Wahyu Hadi
Telegram : @rwahyu
	
