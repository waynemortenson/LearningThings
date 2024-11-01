# 

### Latency
Network
GC
Disk
Algo

### Structure

Set values to zero with var to be clear

```
var thing int
```

Set other values with :=

thing := createThing();

### Escape Analysis

The Go compiler can do static code analysis and does some escape analysis
What happens when we do struct construction through a pointer, we do so on the heap.
```

func createUser() *user{
    u = user {
        Name: "Name
    }

    return &u
    }

```

With escape analysis Go will know that this return back up the heap requires that the construction of user happens on the heap, requiring the GC
While being a user and reading the code however, if we were to move the & to u = &user instead of at the return, we lose readability

To see all the escape analysis, you can do 

```
go build -gcflags -m=2
```

You do not have to guess about anything in this language




### Array Pointer Semantic Loops/ranges

The below code is not effecient, but goes to show that this a pointer semantic range that will change the friends array, as the array was passed in by reference:
```
friends := [3]string{"alice","betty","frank"}

for i := range friends{
    friends[1] = "suzy"

    if(i == 0){
        fmt.Println(friends[1])
    }
}

```

Alternatively, using value semantics will prodivde its own copy of the array

```
friends = [3]string{"alice","betty","frank"}

for i, v := range friends{
     friends[1] = "suzy"

    if(i == 0){
        fmt.Println(friends[1])
    }
}

```

v is our copy of the array.






## Reference Types

### Slice

### Maps

### Channels

### Interfaces

### Functions


### 