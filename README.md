# Analyzing Concurrency Implementation in Prog Two

**By: Kibby Hyacinth Pangilinan**


My code for Prog 2 is showing race conditions. When writing Prog 2, I tried to keep in mind the concepts mentioned in the Go Concurrency slides presentation; specifically, I implemented a mutex lock. When implementing the mutex lock, I applied the similar logic for how mutex locks work with Java threads. I used a lock and unlock to ensure that only one goroutine can access `res` at a time. However, I do not think that this was enough to completely get rid of race conditions in my code. I think that making `main()` sleep may be causing race conditions to still occur. I set `main()` to sleep for three seconds, but some goroutines could take longer to execute, so there is the possibility that `main()` could print an incorrect result. Additionally, I do not have something in the program that tracks the various goroutines that have statrted. I think to help solve the race conditions in this program, I could find a way in Go to ensure that `main()` waits until all goroutines are completed before continuing on.


Additionally, when testing to see how well the program works, I noticed that when increasing the value of `p`, once the number was large enough, the program would print the result to be zero. Although, I do not think that this is due to race conditions. Because `int` has a size limit, when `p` increases, `res` likely becomes too large and causes an overflow. Thus, the program works properly with smaller `p` values, but as the value of `p` becomes larger, the result becomes inaccurate.