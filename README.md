# Buggy Math SaaS - Learning Git Bisect

## About This Repository

This repository is developed as an experimental project to **learn and practice using `git bisect`**. The goal is to intentionally introduce a bug and use `git bisect` to track it down.

## Project Structure

- **math/** - Contains math functions (Add, Subtract, Multiply, etc.)
- **http/** - Implements an HTTP handler to expose math functions as a service
- **tests/** - Unit tests for math operations
- **docs/** - Documentation of lessons learned about debugging

## Development Roadmap

1. **Initial commit:** Add an incomplete math package and a failing unit test.
2. **Fix Add function:** Update implementation so tests pass.
3. **Add HTTP handler:** Create an API endpoint for math operations.
4. **Introduce a bug:** Add a faulty Subtract function to practice debugging.
5. **New feature:** Implement a Multiply function and expand tests.

## Learning Goals

- Master `git bisect` by intentionally debugging a faulty commit.
- Document findings and reflections on debugging techniques.

## How to Use `git bisect`

1. Start bisect mode:
   ```sh
   git bisect start
   ```
2. Mark a known good commit:
   ```sh
   git bisect good <commit-hash>
   ```
3. Mark a bad commit (where the bug exists):
   ```sh
   git bisect bad <commit-hash>
   ```
4. Git will checkout a commit in between. Run tests and mark as good or bad:
   ```sh
   git bisect good  # If the commit is working
   git bisect bad   # If the commit is broken
   ```
5. Repeat until the faulty commit is found.
6. Reset bisect mode:
   ```sh
   git bisect reset
   ```

## Lessons Learned
- **Lesson 1:** Always start with a bad commit marking the current HEAD as bad.
- **Lesson 2:** Find a commit that is known to be good and mark it as good.
- **Lesson 3:** Git will automatically checkout a commit in between.
- **Lesson 4:** Run tests to determine if the commit is good or bad. If good, mark as good. If bad, mark as bad.
- **Lesson 5:** Repeat the process until the faulty commit is found.
- **Lesson 6:** Reset bisect mode after finding the faulty commit.
- **Lesson 7:** Get the commit hash of the faulty commit and analyze the changes using `git show`.

## How to use this repository
- Clone the repository:
  ```sh
  git clone <repository-url>
  ```
- Navigate to the project directory:
  ```sh
   cd buggy-math-go
   ```
- Start bisect mode:
  ```sh
  git bisect start
  ```
- Enjoy debugging!

## References
- [Git Bisect Documentation](https://git-scm.com/docs/git-bisect)
- [Git Bisect — Commit นี้มีปัญหาแน่นะวิ!!!](https://medium.com/@ppraserts/git-bisect-commit-นี้มีปัญหาแน่นะวิ-1ee09c63b3d9)
