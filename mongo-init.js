db = db.getSiblingDB('admin');

// Create 'guest' user with necessary permissions on 'goinggo' database
db.createUser({
  user: "guest",
  pwd: "welcome",
  roles: [
    {
      role: "readWrite",
      db: "goinggo"
    }
  ]
});

// Switch to the 'goinggo' database and add an example collection (optional)
db = db.getSiblingDB('goinggo');
db.testCollection.insert({ message: "Hello, World!" });
