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
	// "strconv"

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

	// NO code should go after this line. it won't ever reach that point
	router.Run(":" + port)
}

type NameCommand struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}