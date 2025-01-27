## Login System Update

The login system has been fixed by implementing the following changes:

1. Replaced the insecure SHA256 hashing with bcrypt, which:
   - Automatically handles salt generation and storage
   - Is specifically designed for password hashing
   - Provides protection against rainbow table attacks

2. Updated the password comparison logic to use bcrypt's secure comparison function

### Important Notes:
1. Existing users will need to reset their passwords as the hashing algorithm has changed
2. Make sure to run `go get golang.org/x/crypto/bcrypt` to install the required dependency

These changes will resolve the login issues while also improving the security of the system.