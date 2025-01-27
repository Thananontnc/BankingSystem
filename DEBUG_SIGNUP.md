# Debugging the Sign-up Process

If data is not reaching the database during sign-up, follow these steps:

1. **Verify MySQL Server**
   - Ensure MySQL server is running on localhost:3306
   - Test credentials: `root:Plai1412`
   - Database name: `bankingsystem`
   - Run: `mysql -u root -p` to test connection

2. **Check Application Logs**
   - Start the application and watch for these log messages:
     ```
     Initializing application...
     Attempting to connect to database...
     Database connection successful
     Starting server on :8080
     ```
   - If you don't see these messages, the server isn't starting properly

3. **Test Sign-up Process**
   - When submitting the form, you should see these log messages:
     ```
     Attempting to register user with email: [your-email]
     Successfully registered user with email: [your-email]
     ```
   - If you see "Failed to register user" or no logs, check:
     - Form data in browser's developer tools (Network tab)
     - MySQL error logs

4. **Common Issues**
   - MySQL server not running
   - Wrong database credentials
   - Database/tables not created
   - Form submission not working

5. **Verify Database Tables**
   ```sql
   USE bankingsystem;
   SHOW TABLES;
   DESCRIBE users;
   SELECT * FROM users;
   ```

If you still don't see data in the database after verifying all these steps, check the browser's developer tools Network tab when submitting the form to ensure the request is being sent to `/register` with proper POST method.