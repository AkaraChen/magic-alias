#!/usr/bin/env bats

load test_helper

# Setup for remove tests - create test aliases
setup_remove() {
  setup
  # Create test aliases
  ma add test1 ls
  ma add test2 pwd
}

@test "ma remove deletes an existing alias" {
  setup_remove
  
  # Verify the alias exists
  assert_file_exists "$MAGIC_ALIAS_DIR/test1"
  
  # Remove the alias
  run ma remove test1
  [ "$status" -eq 0 ]
  [[ "$output" == *"Remove alias successfully"* ]]
  
  # Verify the alias is gone
  assert_file_not_exists "$MAGIC_ALIAS_DIR/test1"
  
  teardown
}

@test "ma remove fails with nonexistent alias" {
  setup_remove
  
  # Try to remove a nonexistent alias
  run ma remove nonexistentalias
  [[ "$output" == *"not found"* ]]
  
  teardown
}

@test "ma rm works as an alias for remove" {
  setup_remove
  
  # Verify the alias exists
  assert_file_exists "$MAGIC_ALIAS_DIR/test2"
  
  # Remove the alias using the 'rm' alias
  run ma rm test2
  [ "$status" -eq 0 ]
  [[ "$output" == *"Remove alias successfully"* ]]
  
  # Verify the alias is gone
  assert_file_not_exists "$MAGIC_ALIAS_DIR/test2"
  
  teardown
}

@test "ma remove with no arguments fails when no aliases exist" {
  setup
  
  # Try to run interactive remove when no aliases exist
  run ma remove
  [ "$status" -ne 1 ]
  [[ "$output" == *"No aliases found to remove"* ]]
  
  teardown
}
