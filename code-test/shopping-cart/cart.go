package main

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type product struct {
	ProductName, SellerName string
}

var tpl *template.Template

var fs = template.FuncMap{
	"rws": ReplaceWhiteSpace,
}

// database
var cartSession = map[string][]product{} // assign cookies to productname
var productDB = map[string]product{
	"Tom Cruise": {"Mercedes Benz", "Tom Criuse"},
} // database for cart items

func ReplaceWhiteSpace(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func init() {
	tpl = template.Must(template.New("").Funcs(fs).ParseFiles("home.html", "cart.html"))
}

func Home(w http.ResponseWriter, r *http.Request) {
	var products []product
	for i, v := range productDB {
		log.Println(i, v)
		products = append(products, v)
	}
	tpl.ExecuteTemplate(w, "home.html", products)
}

func Advertise(w http.ResponseWriter, r *http.Request) {
	sellerName := r.FormValue("seller")
	productName := r.FormValue("product")

	newProduct := product{productName, sellerName}

	productDB[sellerName] = newProduct

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Cart(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("cart")
	if err != nil {
		http.Error(w, "No item(s) in cart", http.StatusNoContent)
		return
	}

	productsInCart := cartSession[cookie.Value]
	tpl.ExecuteTemplate(w, "cart.html", productsInCart)
}

func addToCart(w http.ResponseWriter, r *http.Request) {
	URLpath := r.URL.Path
	productName := URLpath[6:]
	log.Println("Product Name:", productName)

	cookie, err := r.Cookie("cart")
	if err == http.ErrNoCookie {
		uid, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "cart",
			Value: uid.String(),
		}
	}
	var cartProducts []product
	cartProducts = append(cartProducts, productDB[productName])

	cartSession[cookie.Value] = cartProducts

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	productName := path[8:]
	log.Println("Prouctname:", productName)

	cookie, err := r.Cookie("cart")
	if err != nil {
		http.Error(w, "No item(s) in cart", http.StatusNoContent)
		return
	}

	ProductsInCartBeforeRemoval := cartSession[cookie.Value]

	var ProductsInCartAfterRemoval []product

	for _, v := range ProductsInCartBeforeRemoval {
		if v.ProductName != productName {
			ProductsInCartAfterRemoval = append(ProductsInCartAfterRemoval, v)
		}
	}

	cartSession[cookie.Value] = ProductsInCartAfterRemoval

	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/advertise", Advertise)
	http.HandleFunc("/cart", Cart)
	http.HandleFunc("/cart/", addToCart)
	http.HandleFunc("/remove/", RemoveFromCart)

	log.Println("Serving on 3000....\nVisit localhost:3000")
	http.ListenAndServe(":3000", nil)
}
