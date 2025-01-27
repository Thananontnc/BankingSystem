## Important Update - Password System Migration

The password system has been updated to use bcrypt consistently throughout the application. This change affects both new registrations and existing logins. To ensure proper functionality:

1. The database needs to be cleared of existing users OR
2. Existing users need to reset their passwords

### Technical Changes Made:
1. Switched from SHA256 to bcrypt for password hashing
2. Implemented proper password comparison using bcrypt's secure comparison
3. Standardized password handling through utility functions
4. Fixed inconsistent hashing methods between registration and login

### Required Actions:
1. Run: `go get golang.org/x/crypto/bcrypt`
2. Clear existing user accounts or implement a password reset mechanism
3. Restart the application

The login system should now work correctly for all new registrations. Existing users will need to re-register or reset their passwords due to the change in hashing algorithm.