#!/bin/bash

echo "Waiting for MongoDB to start..."

# Wait for MongoDB to be ready
until mongosh --host mongo --port 27017 --eval "print(\"waited for connection\")"; do
    echo "MongoDB is not ready yet, waiting..."
    sleep 2
done

echo "Creating application user and db..."

# Create the database and user
mongosh --host mongo --port 27017 \
    -u "${MONGO_INITDB_ROOT_USERNAME}" \
    -p "${MONGO_INITDB_ROOT_PASSWORD}" \
    --authenticationDatabase admin \
    --eval "db.getSiblingDB('${DB_NAME}').createUser({user: '${DB_USER}', pwd: '${DB_PASSWORD}', roles:[{role:'dbOwner', db: '${DB_NAME}'}]});"
