package main

import (
  "testing"
)

func (t testing.T) TestSomething(){
  if(1==1){
    t.Fail()
  }
}
