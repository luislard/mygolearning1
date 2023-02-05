## Go was designed to handle multiple cores natively

First dual core processor was launch in 2005

Golang creators started this language in 2006 and 2012 was the 1st release of the language

No other languages was built with concurrency in mind, 
no python, no C# no C++, no JAVA, no other language handles this natively.


## Difference between Concurrency and Paralelism

If you Write code and you run that code on a machine that only has 1 CPU
your code is not gonna run in parallel. You will NOT have multiple threads of code
running at the same time.

1 CPU -> Its gonna run sequentially

Concurrency is a design pattern, its a way you could write your code, you can
have code which has the possibility and the potential to run in parallel.

If you have multiple cores that concurrent code can run in parallel. 
Executing at the same time 

Nice talk: https://www.youtube.com/watch?v=oV9rvDllKEg
Slides https://go.dev/talks/2012/waza.slide#10


Rob Pikes defines:

Concurreny:
Its a composition of independently running processes


Parallelism is the simultaneously execution possible related, possibly related.


Concurrency is about dealing with lots of things at once.
Parallelism is about doing lots of things at once

Concurrency is about instruction
Parallelism is about execution

Comunicating sequential processing -> https://www.cs.cmu.edu/~crary/819-f09/Hoare78.pdf
