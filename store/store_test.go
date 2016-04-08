package store_test

import (
   "testing"
   "github.com/ethanaubuchon/lightdm/store"
)

func TestGet(t *testing.T) {
   s := store.Store{}
   actual, err := s.Get("A")
   expected := ""

   if err == nil {
      t.Error("Expected an error")
   }

   if actual != expected {
      t.Error("Expected " + expected + " Actual " + actual)
   }
}

func TestSet(t *testing.T) {
   s := store.Store{}
   expected := "value"
   key := "Key"
   s.Set(key, expected)
   actual, err := s.Get(key)

   if err != nil {
      t.Error("Unexpected Error")
   }

   if actual != expected {
      t.Error("Expected " + expected + " Actual " + actual)
   }
}
