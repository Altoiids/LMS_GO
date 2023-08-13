import argparse
import mysql.connector
import bcrypt

def main(args):
    db_username = args.db_username
    db_password = args.db_password
    db_host = args.db_host
    db_name = args.db_name
    admin_username = args.admin_username
    admin_password = args.admin_password
    user_email = args.user_email
    
    if not db_username:
        db_username = input("Enter your database username: ")

    if not db_password:
        db_password = input("Enter your database password: ")

    if not db_host:
        db_host = input("Enter your database host: ")

    if not db_name:
        db_name = input("Enter your database name: ")

    if not admin_username:
        admin_username = input("Enter username of your first admin: ")

    if not admin_password:
        admin_password = input("Enter password of your first admin: ")

    if not user_email:
        user_email = input("Enter email address of first admin: ")
    
    connection = None 

    try:
        connection = mysql.connector.connect(user=db_username, password=db_password, host=db_host, database=db_name)
        cursor = connection.cursor()

        insert_query = "INSERT INTO user (name, email, hash, adminId) VALUES (%s, %s, %s, %s)"
        hashed_password = bcrypt.hashpw(admin_password.encode('utf-8'), bcrypt.gensalt()).decode('utf-8')
        user_data = (admin_username, user_email, hashed_password, 1)

        cursor.execute(insert_query, user_data)
        connection.commit()

        print("Admin inserted successfully!")

    except mysql.connector.Error as error:
        print("Error:", error)

    finally:
        if connection is not None and connection.is_connected():
            cursor.close()
            connection.close()

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Setup script for database and admin user creation.")
    parser.add_argument("--db-username", default="", help="Database username")
    parser.add_argument("--db-password", default="", help="Database password")
    parser.add_argument("--db-host", default="", help="Database host")
    parser.add_argument("--db-name", default="", help="Database name")
    parser.add_argument("--admin-username", default="", help="Admin username")
    parser.add_argument("--admin-password", default="", help="Admin password")
    parser.add_argument("--user-email", default="", help="Admin email address")
    
    args = parser.parse_args()
    main(args)