// by setting package as main, Go will compile this as an executable file.
// Any other package turns this into a library
package main

// These are your imports / libraries / frameworks
import (
	// this is Go's built-in sql library
	"database/sql"
	"log"
	"net/http"
	"os"
	"strings"
	"strconv"

	// this allows us to run our web server
	"github.com/gin-gonic/gin"

	// this lets us connect to Postgres DB's
	_ "github.com/lib/pq"
)

var (
	// this is the pointer to the database we will be working with
	// this is a "global" variable (sorta kinda, but you can use it as such)
	db *sql.DB
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	var errd error
	// here we want to open a connection to the database using an environemnt variable.
	// This isn't the best technique, but it is the simplest one for heroku
	db, errd = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errd != nil {
		log.Fatalf("Error opening database: %q", errd)
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("html/*.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/register.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	router.GET("/appointment.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "appointment.html", nil)
	})


	router.GET("/ping", func(c *gin.Context) {
		ping := db.Ping()
		if ping != nil {
			// our site can't handle http status codes, but I'll still put them in cause why not
			c.JSON(http.StatusOK, gin.H{"error": "true", "message": "db was not created. Check your DATABASE_URL"})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": "false", "message": "db created"})
		}
	})

	router.POST("/patientQuery", func(c *gin.Context) {

		var name NameCommand

		c.BindJSON(&name)
		log.Println(name.FirstName+" "+name.LastName)
		// var FirstInput = c.PostForm("firstname");
		// var LastInput = c.PostForm("lastname");
		
		// var firstInput = "Peter"
		// var lastInput = "Seo"

		table := "<table class='table'><thead><tr>"

		// put your query here
		rows, err := db.Query("SELECT Visit.VisitDate AS date, Medication.name AS medication, Treatment.type AS treatment, Dentist.FirstName AS dFirstname, Dentist.LastName AS dLastname FROM ChartNote JOIN Medication ON Medication.med_ID = ChartNote.med_ID JOIN Patient ON Patient.patient_ID = ChartNote.patient_ID JOIN Treatment ON Treatment.Treat_ID = ChartNote.Treat_ID JOIN Visit ON Visit.patient_ID = Patient.patient_ID JOIN Dentist ON Dentist.dentist_ID = Visit.dentist_ID WHERE Patient.firstname = $1 AND Patient.lastname = $2;",name.FirstName,name.LastName) // <--- EDIT THIS LINE
		if err != nil {
			// careful about returning errors to the user!
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		// foreach loop over rows.Columns, using value
		cols, _ := rows.Columns()
		if len(cols) == 0 {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		for _, value := range cols {
			table += "<th class='text-center'>" + value + "</th>"
		}
		// once you've added all the columns in, close the header
		table += "</thead><tbody>"
		// declare all your RETURNED columns here
		var date string
		var dateOnly = strings.SplitAfter(date,"T")[0]
		var dFirstname string
		var dLastname string
		var medication string 
		var treatment string
		log.Println(date)
		log.Println(dateOnly)
		for rows.Next() {
			// assign each of them, in order, to the parameters of rows.Scan.
			// preface each variable with &
			rows.Scan(&dateOnly,&medication,&treatment,&dFirstname,&dLastname) // <--- EDIT THIS LINE
			// can't combine ints and strings in Go. Use strconv.Itoa(int) instead
			table += "<tr><td>"+dateOnly+"</td><td>"+medication+"</td><td>"+treatment+"</td><td>"+dFirstname+"</td><td>"+dLastname+"</td></tr>" // <--- EDIT THIS LINE
		}
		// finally, close out the body and table
		table += "</tbody></table>"
		c.Data(http.StatusOK, "text/html", []byte(table))
	})

	router.POST("/addressQuery", func(c *gin.Context) {

		var name NameCommand

		c.BindJSON(&name)
		log.Println(name.FirstName+" "+name.LastName)

		table := "<table class='table'><thead><tr>"
		// put your query here
		rows, err := db.Query("SELECT DISTINCT Address.Street_1 AS street1, Address.Street_2 AS street2, Address.City AS city, Address.State AS state, Address.PostalCode AS postal FROM ChartNote JOIN Medication ON Medication.med_ID = ChartNote.med_ID JOIN Patient ON Patient.patient_ID = ChartNote.patient_ID JOIN Treatment ON Treatment.Treat_ID = ChartNote.Treat_ID JOIN Visit ON Visit.patient_ID = Patient.patient_ID JOIN Dentist ON Dentist.dentist_ID = Visit.dentist_ID JOIN Address ON Patient.address_ID = Address.address_ID WHERE Patient.firstname = $1 AND Patient.lastname = $2;",name.FirstName,name.LastName) // <--- EDIT THIS LINE
		if err != nil {
			// careful about returning errors to the user!
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		// foreach loop over rows.Columns, using value
		cols, _ := rows.Columns()
		if len(cols) == 0 {
			c.AbortWithStatus(http.StatusNoContent)
		}
		for _, value := range cols {
			table += "<th class='text-center'>" + value + "</th>"
		}
		// once you've added all the columns in, close the header
		table += "</thead><tbody>"
		// columns
		var street1 string
		var street2 string
		var city string
		var state string
		var postal string

		for rows.Next() {
			rows.Scan(&street1,&street2,&city,&state,&postal) // put columns here prefaced with &
			table += "<tr><td>"+street1+"</td><td>"+street2+"</td><td>"+city+"</td><td>"+state+"</td><td>"+postal+"</td></tr>" // <--- EDIT THIS LINE
		}
		// finally, close out the body and table
		table += "</tbody></table>"
		c.Data(http.StatusOK, "text/html", []byte(table))
	})

	router.POST("/dentistQuery", func(c *gin.Context) {
		var name NameCommand

		c.BindJSON(&name)
		log.Println(name.FirstName+" "+name.LastName)

		table := "<table class='table'><thead><tr>"
		// put your query here
		rows, err := db.Query("SELECT DISTINCT Dentist.Phone AS phone, Dentist.Email AS email, Dentist.HireDate AS date FROM Dentist WHERE Dentist.firstname = $1 AND Dentist.lastname = $2;",name.FirstName,name.LastName) // <--- EDIT THIS LINE
		if err != nil {
			// careful about returning errors to the user!
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		// foreach loop over rows.Columns, using value
		cols, _ := rows.Columns()
		if len(cols) == 0 {
			c.AbortWithStatus(http.StatusNoContent)
		}
		for _, value := range cols {
			table += "<th class='text-center'>" + value + "</th>"
		}
		// once you've added all the columns in, close the header
		table += "</thead><tbody>"
		// columns
		var phone string
		var email string
		var date string

		for rows.Next() {
			rows.Scan(&phone,&email,&date) // put columns here prefaced with &
			table += "<tr><td>"+phone+"</td><td>"+email+"</td><td>"+date+"</td></tr>" // <--- EDIT THIS LINE
		}
		// finally, close out the body and table
		table += "</tbody></table>"
		c.Data(http.StatusOK, "text/html", []byte(table))
	})

	router.POST("/registerQuery", func(c *gin.Context) {
		var info InfoCommand

		c.BindJSON(&info)
		
		// result, err := db.Exec("INSERT INTO Patient (Address_ID, FirstName, LastName, Email, DOB) VALUES ($1,$2,$3,$4,$5);",nil,info.FirstName,info.LastName,info.Email,info.Date);
		rows2, err := db.Query("INSERT INTO Patient (FirstName, LastName, Email, DOB) VALUES ($1,$2,$3,$4);",info.FirstName,info.LastName,info.Email,info.Date);
		// log.Println(info.FirstName)
		// log.Println(info.LastName)
		// log.Println(info.Email)
		// log.Println(info.Date)
		_=rows2
		// _=result
		_=err

	})

	router.POST("/registerShow", func(c *gin.Context) {

		table := "<table class='table'><thead><tr>"
		// put your query here
		rows, err := db.Query("SELECT patient_ID As id, firstName AS firstname, lastName AS lastname FROM Patient ORDER BY patient_ID ASC") // <--- EDIT THIS LINE
		if err != nil {
			// careful about returning errors to the user!
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		// foreach loop over rows.Columns, using value
		cols, _ := rows.Columns()
		if len(cols) == 0 {
			c.AbortWithStatus(http.StatusNoContent)
		}
		for _, value := range cols {
			table += "<th class='text-center'>" + value + "</th>"
		}
		// once you've added all the columns in, close the header
		table += "</thead><tbody>"
		// columns
		var id int
		var firstname string
		var lastname string

		for rows.Next() {
			rows.Scan(&id,&firstname,&lastname) // put columns here prefaced with &
			table += "<tr><td>"+strconv.Itoa(id)+"</td><td>"+firstname+"</td><td>"+lastname+"</td></tr>" // <--- EDIT THIS LINE
		}
		// finally, close out the body and table
		table += "</tbody></table>"
		c.Data(http.StatusOK, "text/html", []byte(table))
	})

	router.POST("/appointQuery", func(c *gin.Context) {
		var info InfoCommand

		c.BindJSON(&info)
		
		// result, err := db.Exec("INSERT INTO Patient (Address_ID, FirstName, LastName, Email, DOB) VALUES ($1,$2,$3,$4,$5);",nil,info.FirstName,info.LastName,info.Email,info.Date);
		rows2, err := db.Query("INSERT INTO Visit (Patient_ID, Dentist_ID, VisitDate, IsAppointed) VALUES ((SELECT patient_ID FROM Patient WHERE firstname = $1 AND lastname = $2),null,CAST($3 AS date),'TRUE');",info.FirstName,info.LastName,info.Date);
		// log.Println(info.FirstName)
		// log.Println(info.LastName)
		// log.Println(info.Email)
		// log.Println(info.Date)
		_=rows2
		// _=result
		_=err
	})

	router.POST("/appointShow", func(c *gin.Context) {

		var today TodayCommand
		c.BindJSON(&today)

		table := "<table class='table'><thead><tr>"
		// put your query here
		rows, err := db.Query("SELECT Visit.VisitDate AS date, Patient.patient_ID AS id, Patient.firstName AS firstname, Patient.lastName AS lastname FROM Visit JOIN Patient ON Patient.patient_ID = Visit.patient_ID WHERE Visit.VisitDate >= date($1) ORDER BY VisitDate ASC;",today.Today) // <--- EDIT THIS LINE
		if err != nil {
			// careful about returning errors to the user!
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		// foreach loop over rows.Columns, using value
		cols, _ := rows.Columns()
		if len(cols) == 0 {
			c.AbortWithStatus(http.StatusNoContent)
		}
		for _, value := range cols {
			table += "<th class='text-center'>" + value + "</th>"
		}
		// once you've added all the columns in, close the header
		table += "</thead><tbody>"
		// columns
		var id int
		var date string
		var firstname string
		var lastname string

		for rows.Next() {
			rows.Scan(&date,&id,&firstname,&lastname) // put columns here prefaced with &
			table += "<tr><td>"+date+"</td><td>"+strconv.Itoa(id)+"</td><td>"+firstname+"</td><td>"+lastname+"</td></tr>" // <--- EDIT THIS LINE
		}
		// finally, close out the body and table
		table += "</tbody></table>"
		c.Data(http.StatusOK, "text/html", []byte(table))
	})
	// NO code should go after this line. it won't ever reach that point
	router.Run(":" + port)
}

type NameCommand struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

type InfoCommand struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Date string `json:"date"`
}

type TodayCommand struct {
    Today string `json:"today"`
}