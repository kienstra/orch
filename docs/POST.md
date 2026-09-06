## Testing Without Mocks

Good tests come from good architecture.

Advanced testing techniques are a sign that the architecture isn't good.

Mocks give a false sense of confidence. They say that "when you call X with Y, it returns Z".

But how do you know it'll return Z?

If you haven't seen that, it's just wishful thinking.

Eventually, tests with mocks become testing the tests.

But there's no guarantee that X will happen.

Mocks have their place.

But sometimes when you make a small change, you have to update 10 mocks in test cases.

Here's a way to avoid that.

1. Make every part isolated.
2. Make as much as possible a pure function that you can test by itself.
3. If something can't be a pure function, like DB access, test that by itself.

If it's so simple that it doesn't need tests, just manually test it once, and don't add automated tests.

DB access is actually a small part of the system.

Ideal
Architecture that doesn't require many mocks. Just simple stubs.

Mocks
1. Make you change all of them when you change production code
2. Give a false sense of confidence. "With a query of X, the DB will return y." But have you seen it return Y?
The tests are based on what you think, not what has happened.
Better to not have the tests, as they're deceptive.

Most API logic doesn't need integration tests.
For example, if there's a malformed HTTP query, there's no need to connect to the DB.

The query will fail before that.
Also, if there's an error in the domain logic, there's no need to test it with the DB.
If the DB (repository) is so complex that it requires tests, that's fine.
It can test only that layer.


