#!/bin/bash

read -p "Enter your database username: " DB_USERNAME
read -s -p "Enter your database password: " DB_PASSWORD
echo
read -p "Enter your database host: " DB_HOST
read -p "Enter your database name: " DB_NAME
read -p "Enter your JWT secret: " JWT_SECRETKEY
read -p "Enter username of your first admin: " ADMIN_USERNAME
read -s -p "Enter password of your first admin: " ADMIN_PASSWORD
read -p "Enter email address of first admin: " USER_EMAIL

EMAIL_PATTERN="^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,4}$"

if [[ $USER_EMAIL =~ $EMAIL_PATTERN ]]; then
    echo "Thank you for providing a valid email address: $USER_EMAIL"
else
    echo "Invalid email address format."
fi
echo

HASHED_PASSWORD=$(python3 -c "import bcrypt; print(bcrypt.hashpw('$ADMIN_PASSWORD'.encode('utf-8'), bcrypt.gensalt()).decode('utf-8'))")

cat <<EOF > config.yaml
DB_USERNAME: $DB_USERNAME
DB_PASSWORD: '$DB_PASSWORD'
DB_HOST: $DB_HOST
DB_NAME: $DB_NAME
JWT_SECRET: "$JWT_SECRET"
EOF

echo "config.yaml created successfully with the provided values."

migrate -path database/migration/ -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp(localhost:3306)/$DB_NAME" -verbose up

sudo apt install python3-pip
pip install mysql-connector-python

python3 <<EOF
import mysql.connector
import bcrypt

# Database connection parameters
db_username = "$DB_USERNAME"
db_password = "$DB_PASSWORD"
db_host = "$DB_HOST"
db_name = "$DB_NAME"

# User information
admin_username = "$ADMIN_USERNAME"
hashed_password = bcrypt.hashpw('$ADMIN_PASSWORD'.encode('utf-8'), bcrypt.gensalt()).decode('utf-8')
admin_email = "$USER_EMAIL"

try:
    connection = mysql.connector.connect(user=db_username, password=db_password, host=db_host, database=db_name)
    cursor = connection.cursor()

    insert_query = "INSERT INTO user (name, email, hash, admin_id) VALUES (%s, %s, %s, %s)"
    user_data = (admin_username, admin_email, hashed_password, 1)

    cursor.execute(insert_query, user_data)
    connection.commit()

    print("Admin inserted successfully!")

except mysql.connector.Error as error:
    print("Error:", error)

finally:
    if connection.is_connected():
        cursor.close()
        connection.close()
EOF