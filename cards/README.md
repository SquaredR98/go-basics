### 1. Declare and define variables
- In order to decalre and define variable we must use the following syntaxes:

  1. **Decalre and Define Separately**: Here we decide the type of the variable.
      ```
      func main () {
        var card string
        card = "Spade of Diamonds"  
      }
      ```

  2. **Declare and Define at the same time**: Here compiler decides the type of the variable.
      ```
      ---
        card := "Spade of diamonds"
      ---
      ```
  
  *Note*: Never try to reassign variables by `card := "Spade of diamonds"` because it will throw error that can't redeclare or something.

### 2. Defining Functions

```
func newCard () {
  return "Ace of Diamonds"
}
```

### 3. Data Structure to hold list

- Two data structures available:

    1. `Array`: Static Size. Primitive DS for holding list. Must br declared with proper size.
    2. `Slices`: Dynamic Size. Can be increased and decreased based on the requirement dynamically. Every element must be of same type. To create a new `slice` we must do it in the following way:

        ```
        cards := []string { "stringOne", "stringTwo" }
        ```


### 4. Slices

1. Add item to slices: `append(cards, "New Suite Card")` 
2. Looping a slice:
    ```
      for i, card := range cards {
        fmt.Println(i, card)
      }
    ```