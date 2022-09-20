package main

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("could not read env file:", err)
		return
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Println("error unmarshalling configuration variables:", err)
		return

	}

	credential := options.Credential{
		Username: config.DBUsername,
		Password: config.DBPassword,
	}

	clientOpts := options.Client().ApplyURI("mongodb://localhost:27000").SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Println("error connecting:", err)
		return
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("users").Collection("coll")

	tpl := template.Must(template.ParseFiles("./index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if err := tpl.Execute(w, nil); err != nil {
				http.Error(w, "something went wrong", http.StatusInternalServerError)
			}
		case http.MethodPost:
			fname, lname := r.FormValue("fname"), r.FormValue("lname")

			collection.InsertOne(context.TODO(), bson.M{
				"firstname": fname,
				"lastname":  lname,
			})
			tpl.Execute(w, "successfully added record")
		}
	})
	
	http.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		var values []bson.M

		cursor, err := collection.Find(context.TODO(), bson.M{})
		if err != nil {
			log.Println("error finding collection:",err)
			http.Error(w, "something went wrong", 500)
			return
		}
		defer cursor.Close(context.TODO())

		if err := cursor.All(context.TODO(), &values); err != nil {
			log.Println("error writing data:",err)
			http.Error(w, "something went wrong", 500)
			return
		}
		
		json.NewEncoder(w).Encode(values)
	})

	http.ListenAndServe(":4563", nil)
}

// bs, _ := bcrypt.GenerateFromPassword([]byte("securedandrewwilder05"), bcrypt.MinCost)
// fmt.Println(string(bs))
