package main

/**
 * Login routines.
 */
func (c *Client) MatrixLogin() {
	// Prompt for user name/handle/alias
	handle := c.Prompt(64, "Enter your handle/ID#:")
	// Make sure it contains valid characters.
	if (ValidUserName(handle)) {
		// Prompt for password and compare.
		if (c.GetPassword("users.password")) {

		}
	}
}

/**
 * New user application.
 */
func MatrixApply() {

}
