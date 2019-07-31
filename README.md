# threadedSleepSort
A Go based multi-threaded sorting algorithm
This works by creating n threads where n is the size of the input array
Each thread sleeps by the multiple of the element at the index
The shared variable is protected by mutex locks
A test code is written in the main function
A function also compares the sorting result with the built in sorting function
and returns true if the result is correct
