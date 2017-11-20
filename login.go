package main

/**
 * Login routines.
 */
func MatrixLogin() {
	// Prompt for user name/handle/alias
	handle := Prompt(64, "Enter your handle/ID#:")
	// Make sure it contains valid characters.
	if (ValidUserName(handle)) {
		// Prompt for password and compare.
		if (GetPassword("users.password")) {

		}
	}
}

/**
 * New user application.
 */
func MatrixApply() {

}
