package main_test

import (
  "os"
  "testing"
)

func TestMain(m *testing.M) {
    println("before all...")

    code := m.Run()

    println("after all...")

    os.Exit(code)
}

func TestHello(t *testing.T) {
  actual := "hello"
  expected := "hell"
  if actual != expected {
      t.Errorf("got: %v\nwant: %v", actual, expected)
  }
}
