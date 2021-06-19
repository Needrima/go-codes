package main

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

//helper variables
var collection *mgo.Collection //mongodb collection
var tpl *template.Template     //templates

//product object
type Product struct {
	ID                                               int
	ProductName, SellerName, Email, Phone, ImageName string
}

//helper functions
func init() { //parses html templates
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func Check(msg string, err error) { //check errors
	if err != nil {
		log.Println(msg, ":", err)
		return
	}
}

func Found(items []string, item string) bool { //check if an slice contain an item
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}

//func main
func main() {
	//database connection
	dialInfo, err := mgo.ParseURL("mongodb+srv://e-shop-test:eshop15@e-shop-test.csoa6.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	if err != nil {
		log.Fatalln("Dialinfo:", err)
		return
	}
	session, err := mgo.DialWithInfo(dialInfo)
	Check("Session", err)
	defer session.Close()

	collection = session.DB("golang").C("go_practice")

	//routes
	http.HandleFunc("/", Visit)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/advertise", Advertise)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	//listener; listening on port ":8080"
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("ListenAndServe", err)
	} else {
		fmt.Println("Serving on port 8080...")
	}
}

//default route
func Visit(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

//handle operages on homepage
func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet { //if method is get
		var products []Product         //products variable to store all the products from database
		var selectedproducts []Product // selectedproducts variable to store all the products with even IDs from database

		err := collection.Find(bson.M{}).All(&products) //get all products from database
		Check("Error getting product for homepage:", err)

		for _, v := range products {
			if v.ID%2 == 0 {
				selectedproducts = append(selectedproducts, v) //get products with even IDs
			}
		}

		tpl.ExecuteTemplate(w, "index.html", selectedproducts) //send data to template
	} else if r.Method == http.MethodPost { // if method is post
		productName := r.FormValue("product")        //get product to buy
		productName = strings.Trim(productName, " ") //trim out beginning and trailing whitespaces

		var products []Product         //products variable to store all the products from database
		var selectedproducts []Product // selectedproducts variable to store all the products with given product name from database

		err := collection.Find(bson.M{}).All(&products) //get products from db
		Check("Find:", err)

		for _, v := range products {
			if strings.Contains(v.ProductName, productName) {
				selectedproducts = append(selectedproducts, v) //get selected products
			}
		}

		tpl.ExecuteTemplate(w, "index.html", selectedproducts) // send data to template
	}
}

func Advertise(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet { //if method is get
		tpl.ExecuteTemplate(w, "advertise.html", "Advertise new product")
	} else if r.Method == http.MethodPost { //if method is post
		pn := r.FormValue("product") //product name
		sn := r.FormValue("seller")  //seller's name
		em := r.FormValue("email")   //seller's email
		ph := r.FormValue("phone")   //seller's phone number
		rand.Seed(time.Now().UnixNano())
		id := rand.Intn(100) //product id
		var img string

		file, header, err := r.FormFile("image") //get image form file
		Check("Could not get image file", err)

		fileparts := strings.Split(header.Filename, ".")

		extension := fileparts[len(fileparts)-1] //get image extension

		acceptedExtensions := []string{"png", "jpeg", "jpg", "gif"}

		if !Found(acceptedExtensions, extension) { //check if file is an image
			http.Error(w, "File not an image file", 400)
			return
		}

		bs, err := ioutil.ReadAll(file) //read image file
		Check("Error reading file", err)

		tempfile, err := ioutil.TempFile("images", "*."+extension) //create tempfile to store images
		Check("Error creating tempfile:", err)

		tempfile.Write(bs) //store img files in tempfile

		fileinfos, err := ioutil.ReadDir("images") //read directory to get image files
		Check("Error reading images directory:", err)

		img = fileinfos[len(fileinfos)-1].Name() //assign name of last image to var img

		err = collection.Insert(&Product{id, pn, sn, em, ph, img}) //store product data in database
		Check("Error storing data in db", err)                     //check err
		tpl.ExecuteTemplate(w, "advertise.html", "Product added")  //send confirmation info to template
	}
}
