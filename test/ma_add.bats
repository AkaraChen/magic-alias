#!/usr/bin/env bats

load test_helper

@test "ma add creates a new alias with valid arguments" {
  # Add a simple alias for a common command
  run ma add ls ls
  [ "$status" -eq 0 ]
  [[ "$output" == *"Add alias successfully"* ]]
  
  # Check that the alias file exists
  assert_file_exists "$MAGIC_ALIAS_DIR/ls"
  
  # Check that the alias file is executable
  [ -x "$MAGIC_ALIAS_DIR/ls" ]
  
  # Check the content of the alias file
  assert_file_contains "$MAGIC_ALIAS_DIR/ls" "#!/usr/bin/env bash"
  assert_file_contains "$MAGIC_ALIAS_DIR/ls" "exec ls \"\$@\""
}

@test "ma add fails with invalid alias name" {
  # Try to add an alias with invalid characters
  run ma add "invalid@alias" ls
  [ "$status" -ne 0 ]
  [[ "$output" == *"alias must contain only alphanumeric characters, underscores, and hyphens"* ]]
  
  # Check that the alias file doesn't exist
  assert_file_not_exists "$MAGIC_ALIAS_DIR/invalid@alias"
}

@test "ma add fails with empty alias name" {
  # Try to add an alias with empty name
  run ma add "" ls
  [ "$status" -ne 0 ]
  [[ "$output" == *"alias cannot be empty"* ]]
}

@test "ma add fails with nonexistent command" {
  # Try to add an alias for a command that doesn't exist
  run ma add test nonexistentcommand
  [ "$status" -ne 0 ]
  [[ "$output" == *"command not found"* ]]
  
  # Check that the alias file doesn't exist
  assert_file_not_exists "$MAGIC_ALIAS_DIR/test"
}

@test "ma add fails with only one argument" {
  # Try to add an alias with only one argument
  run ma add onlyalias
  [ "$status" -ne 0 ]
  [[ "$output" == *"requires either no arguments or at least 2 arguments"* ]]
}
