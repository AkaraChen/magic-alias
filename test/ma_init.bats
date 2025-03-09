#!/usr/bin/env bats

load test_helper

@test "ma init creates magic-alias directory" {
  # Remove the directory created in setup
  rm -rf "$MAGIC_ALIAS_DIR"
  
  run ma init
  [ "$status" -eq 0 ]
  [ -d "$MAGIC_ALIAS_DIR" ]
}

@test "ma init adds configuration to shell rc file" {
  run ma init
  [ "$status" -eq 0 ]
  
  # Check that the rc file contains the magic alias configuration
  assert_file_contains "$TEST_RC_FILE" "# Magic Alias (ma)"
  assert_file_contains "$TEST_RC_FILE" "export PATH=\"\$PATH:\$HOME/.magic-alias\""
  assert_file_contains "$TEST_RC_FILE" "eval \"\$(ma completion bash)\""
  assert_file_contains "$TEST_RC_FILE" "# End of Magic Alias"
}

@test "ma init is idempotent" {
  # First init
  run ma init
  [ "$status" -eq 0 ]
  
  # Get the file content after first init
  local first_content=$(cat "$TEST_RC_FILE")
  
  # Second init should fail because it's already initialized
  run ma init
  [ "$status" -ne 0 ]
  [[ "$output" == *"magic alias line already exists in rc file"* ]]
  
  # File content should be unchanged
  local second_content=$(cat "$TEST_RC_FILE")
  [ "$first_content" = "$second_content" ]
}
