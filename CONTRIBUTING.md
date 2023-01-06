# How to contribute
## Process
1. Fork this repo
2. Clone your repo to local
3. Create a branch based on master
4. Add your code
5. You can use the `make self_test` to test your code. It will build this package, then generate code for the table in `spec/schema`.
6. Use `make fmt` format your code, and use `make lint` to check the code styles.
7. Follow the base rule to create a pull request, happy coding!
## Base Rule
1. PR only can have one commit, and it needs to rebase the master branch
2. It must have UTs, if it is possible.
3. Commit Message must fit the format. 
```
EF-{Number}(label): title
label:
- feat, for feature
- fix, for bug fix
- ut, for ut
- doc, for document
- refactor, for refactor code
```
Happy Coding, Happy Sharing!
