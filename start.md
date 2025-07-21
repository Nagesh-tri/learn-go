# first thing: install go

# for cmd line args, go has no tkaing args in main fxn, instead it uses "os" lib.
- good with hello program, 
- now create a little local package for hello / greeting fxn.
- time for test package: 
    - need module "testing"
    - made some good test cases for greet and greetMany.
- now work with, stdinp and output, with numbers and stuff
  - how to give files as input '$ go run . < nums.txt' and `$ cat nums.txt | go run .`