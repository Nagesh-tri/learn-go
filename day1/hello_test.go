package hello;

import (
	"testing"
	);

func TestGreet(t *testing.T){
	want := "Hello, test!";
	got :=Greet("test");

	if want !=got {
		t.Errorf("wanted %s, got %s", want , got );
	}
}

func TestGreetMany(t *testing.T){
	subtests := []struct{
		items []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items: []string{"test"},
			result: "Hello, tst!",
		},
	}
	for _,st := range subtests{
		if s:=GreetMany(st.items); s!=st.result{
			t.Errorf("wanted %s [%v], and got %s",st.result,st.items,s);
		}
	}
}