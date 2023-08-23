# poker-hands

## Poker hand evaluator program
Pass in a string of poker hands and the program will sort them in descending order by strongest hand

## Getting Started
### Input
- Format for each hand is as follows: "XY,XY,XY,XY,XY" where X represents the rank, and Y represents the suit.
- Valid X characters are [2, 3, 4, 5, 6, 7, 8, 9, T, J, Q, K, A].
- Valid Y characters are [S, H, D, C] for Spades, Hearts, Diamonds, Clubs.

### Running the program
- Run the program with `go run main.go`.
- The program will prompt you to enter your poker hands.
- Input a poker hand with the format above, then click enter.
- To compare multiple hands, enter additional poker hands and click enter in between each.
- To finish and start comparing, type `done` and then click enter.

### Results
- The program will output all poker hands listed in order from strongest to weakest.
- Each hand will also be provided with a value, which is a unique representation of its strength.

### Structures
- Each card is represented by a struct `Card -> {Rank: int, Suit: string}`.
  - The rank is represented by an integer in the range [2, 14] mapping from 2 to Ace.
- Each hand is represented by a struct `Hand -> {Cards: []Card, Type: int, Score: int}`.
  - The Cards are represented by the Card struct above.
  - The type is an integer value ranging from [0, 9], mapping the different type of hands (i.e. Royal Flush = 9, High Card = 0).
- These structs allow easier access to specific values, and provides cleaner code.

### Tests
- To run tests, run `go test -v ./...` to run the full suite of tests.

### Algorithm
- To compare different hands, this program uses a bit representation of each hand to determine a unique value.
- The card types can be coded in 4 bits, and each card afterwards can also be coded in 4 bits each.
- This brings the total to 4 * 6 = 24 bits.
- In sorted order, we can organize our bits as [type card card card card card] with the strongest card in front.
- If we convert this entire binary representation into an integer, we can obtain a unique value to compare against.
- For hands of the same type but different suits (i.e. royal flush with spades vs royal flush with diamonds),
  this algorithm will give the same value as it won't compare against suits. This is due to the fact that in game,
  it's impossible to achieve these two hands, so it's been left out here.

### Time Complexity
- For input size k, and hand of size n cards.
- For each input, we loop through `O(k)` times and validate each card `O(n)` -> `O(k*n)`.
- When evaluating, we loop through
  - `O(k)` hands
    - `O(n)` cards
    - `O(nlogn)` card sort
    - `O(9 * n) -> O(n)` evaluations to determine hand type
  - We get `O(k * (2n + nlogn))`.
  - Then we sort the results again in `O(k)` time.
- Our total time is `O(k * (2n + nlogn))`.
- Improvements can be made when determining hand types. Perhaps we can combine some of the calculations to cut down on 9 each time.
