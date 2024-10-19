package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"

// 	"gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"
// )

// // Model representing a simple item
// type Item struct {
// 	ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
// 	Name string        `bson:"name" json:"name"`
// }

// // MongoDB connection settings
// const (
// 	MongoDBHosts = "mongo:27017" // MongoDB host
// 	AuthDatabase = "mydb"        // Database name
// 	AuthUserName = "root"        // Username
// 	AuthPassword = "example"     // Password
// )

// // Global session
// var mongoSession *mgo.Session

// // Get MongoDB collection
// func getCollection() (*mgo.Session, *mgo.Collection) {
// 	session := mongoSession.Copy()
// 	collection := session.DB("mydb").C("items") // You can set the database and collection names here
// 	return session, collection
// }

// // Create a new item
// func createItem(w http.ResponseWriter, r *http.Request) {
// 	var item Item
// 	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	session, collection := getCollection()
// 	defer session.Close()

// 	item.ID = bson.NewObjectId()
// 	if err := collection.Insert(item); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(item)
// }

// // Get all items
// func getItems(w http.ResponseWriter, r *http.Request) {
// 	session, collection := getCollection()
// 	defer session.Close()

// 	var items []Item
// 	if err := collection.Find(nil).All(&items); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(items)
// }

// // Update an item
// func updateItem(w http.ResponseWriter, r *http.Request) {
// 	var item Item
// 	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	session, collection := getCollection()
// 	defer session.Close()

// 	if err := collection.UpdateId(item.ID, &item); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(item)
// }

// // Delete an item
// func deleteItem(w http.ResponseWriter, r *http.Request) {
// 	var item Item
// 	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	session, collection := getCollection()
// 	defer session.Close()

// 	if err := collection.RemoveId(item.ID); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }

// // func CreateSession() (*mgo.Session, error) {
// // 	fmt.Println("create session")
// // 	// DialInfo holds the details for the connection
// // 	mongoDBDialInfo := &mgo.DialInfo{
// // 		Addrs:    []string{"mongo:27017"},
// // 		Timeout:  60 * time.Second,
// // 		Database: "admin",   // Database to connect to
// // 		Username: "root",    // MongoDB root username
// // 		Password: "example", // MongoDB root password
// // 		Source:   "admin",   // Authentication database (admin for root user)
// // 	}

// // 	fmt.Println("dia info with ", mongoDBDialInfo)
// // 	// Create a session to the MongoDB database
// // 	session, err := mgo.DialWithInfo(mongoDBDialInfo)
// // 	if err != nil {
// // 		fmt.Printf("error - dial info %+v \n", err)
// // 		return nil, err
// // 	}

// // 	fmt.Println("ping...")
// // 	// Optional: Ensure the session is connected
// // 	err = session.Ping()
// // 	if err != nil {
// // 		fmt.Printf("error - ping %+v \n", err)
// // 		return nil, err
// // 	}

// // 	log.Println("Successfully connected to MongoDB!")
// // 	return session, nil
// // }

// // func connectToMongo(mongoURL string) (*mgo.Session, error) {
// // 	for i := 0; i < 3; i++ {
// // 		fmt.Println("mongo url is ", mongoURL)
// // 		mongoDBDialInfo := &mgo.DialInfo{
// // 			Addrs:    []string{"mongo:27017"},
// // 			Timeout:  60 * time.Second,
// // 			Database: AuthDatabase,
// // 			Username: AuthUserName,
// // 			Password: AuthPassword,
// // 		}

// // 		session, err := mgo.DialWithInfo(mongoDBDialInfo)
// // 		if err == nil {
// // 			return session, nil
// // 		}
// // 		log.Printf("Failed to connect to MongoDB (attempt %d/%d): %v", i+1, 3, err)
// // 		time.Sleep(5 * time.Second)
// // 	}
// // 	return nil, errors.New("exceeded max retries to connect to MongoDB")
// // }

// // Initialize MongoDB connection using mgo.ParseURL
// func connectToMongo(mongoURL string) (*mgo.Session, error) {
// 	var err error
// 	var session *mgo.Session

// 	for i := 0; i < 3; i++ {
// 		fmt.Println("mongoURL ", mongoURL)
// 		// Connect to MongoDB
// 		session, err = mgo.Dial("mongodb://root:example@mongo:27017/admin")
// 		if err == nil {
// 			fmt.Println("Successfully connected to MongoDB!")
// 			return session, nil
// 		}

// 		log.Printf("Failed to connect to MongoDB (attempt %d/%d): %v", i+1, 3, err)
// 		time.Sleep(2 * time.Second)
// 	}
// 	return nil, fmt.Errorf("could not connect to MongoDB after %d attempts: %v", 3, err)
// }

// // HTTP Handlers setup
// func main() {
// 	mongoURL := os.Getenv("MONGO_URL")
// 	fmt.Println("mongo url is ", mongoURL)
// 	// Connect to MongoDB
// 	var err error
// 	// mongoSession, err = CreateSession()
// 	mongoSession, err = connectToMongo(mongoURL)
// 	if err != nil {
// 		log.Fatalf("Error connecting to MongoDB: %v", err)
// 	}
// 	defer mongoSession.Close()

// 	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case "POST":
// 			createItem(w, r)
// 		case "GET":
// 			getItems(w, r)
// 		default:
// 			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		}
// 	})

// 	http.HandleFunc("/items/update", updateItem)
// 	http.HandleFunc("/items/delete", deleteItem)

// 	log.Println("Server started at :8080")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatalf("Error starting server: %v", err)
// 	}
// }
