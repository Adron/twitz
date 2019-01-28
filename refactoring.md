# Refactoring

1. Create a new package folder/library called "helpers".
2. Rename "check" to "Check" so that it will be visible when moved into the helpers library.
3. Move the "Check" function into the "helpers" library and let it create the file in the directory, naming it "helpers" (leave off the go, since Goland will add the .go).
4. Rename "buildTwitterList" for visibility.
5. Rename "getBearerToken" for visibility.
6. Rename "validateRequiredConfig" for visibility.
7. Then move "BuildTwitterList", "GetBearerToken", "ValidateRequiredConfig", and "Contains" to the "helpers.go" file with the move refactor command.

# Adding "twitTwitz"

1. Add a folder called "twitTwitz".
2. Refactor by moving the "printUserToConsole" and "printUsersToConsole" functions into this new library. (They're in the findem.go file) Both will need renamed via refactor first for visibility.

# Deletions

1. Go ahead and just delete the "webscanem.go" file.

# Tests

1. Add a "main_test.go" file to the root of the project. Add the following code to get the remote build going again.

func TestTimeConsuming(t *testing.T) {
	if true {
		t.Skip("skipping test to get build started again.")
	}
}

