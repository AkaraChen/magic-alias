#!/usr/bin/env bats

load test_helper

# Setup for list tests - create test aliases
setup_list() {
  setup
  # Create test aliases
  ma add test1 ls
  ma add test2 pwd
  ma add test3 echo
}

@test "ma list shows all aliases" {
  setup_list
  
  # Run the list command
  run ma list
  [ "$status" -eq 0 ]
  
  # Check that all aliases are listed
  [[ "$output" == *"test1"* ]]
  [[ "$output" == *"test2"* ]]
  [[ "$output" == *"test3"* ]]
  
  teardown
}

@test "ma list shows empty message when no aliases exist" {
  setup
  
  # Run the list command when no aliases exist
  run ma list
  [ "$status" -eq 0 ]
  [[ "$output" == *"No aliases found"* ]]
  [[ "$output" == *"Use 'ma add' to create one"* ]]
  
  teardown
}

@test "ma list fails with arguments" {
  setup
  
  # Run the list command with arguments
  run ma list extraarg
  [ "$status" -ne 0 ]
  [[ "$output" == *"accepts no arg"* ]]
  
  teardown
}
