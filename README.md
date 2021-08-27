# fibonacci

This is basically an implementation of the fibonacci sequence which allows arbitrarily large integers. Also, the functionality is exposed in an HTTP endpoint, so we can make a simple GET request, specifying the n parameter, and we'll get the n fibonacci term.

## Technical details

- As suggested, the [big](https://pkg.go.dev/math/big) package is used for managing any integer.
- The fibonacci logic was extracted in a separated module. A new type `FibonacciCalculator` was created, with only one method (`GetTerm`). This type contains an in-memory cache which is used to improve the performance of the calculation.
- As we're generating a new FibonacciCalculator for every incoming request, our entire system is thread-safe. One downside of this is that we're using separated caches and throwing them away when each request is finished. To improve that aspect of the implementation we could use a common cache, however if we do so we need to use (or implement) something that can actually deal with concurrency. There are many ways to solve this: adding synchronization logic, using [sync.Map](https://pkg.go.dev/sync#Map), using a centralized and more complex solution (redis for example). Every solution has its up and downs.

## Compiling and testing

In order to compile we can just type `make` and `make run`. Then the server will be accepting GET requests (on port 8080) to the / endpoint with the corresponding n parameter.
To run the tests we just need to type `make test`.

## Deployment

This code was deployed to heroku as well and the URL of the app is: https://fib-challenge-rollee.herokuapp.com/ Note that this is using the heroku free tier, so the app sleeps if it is not used for some time and it can be kind of slow to start up again. Once it's started it should work normally.
