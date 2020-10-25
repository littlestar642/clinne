# Clinne

Clinne is an open-source cli application that helps one learn clean code and play all the while doing so. One can install it on their operating systems and then play the game using the terminal. Clinne got its name from the abbreviations CLI and NNE. CLI stands for Command Line Interface and NNE for Non-Negotiable Etiquettes.

The NNEs we are expected to follow while writing Clean Code can be summarized in eleven points. These are:-

1. Indentation and spacing between code constructs (classes/methods/specs) must be consistent
2. Use only spaces (no tabs) for indentation
3. Newlines at end of file
4. Follow accepted naming conventions for your language/framework
5. Follow accepted naming file and Directory structure for your language/framework
6. Use namespaces
7. No comments/Unused Code must ever be checked in
8. Runtime environment should be consistent with IDE environment - i.e there should be no difference in running a build or a spec from your IDE and from the command line
9. Use .gitignore
10. Ensure there is a README.md that includes: Problem Description, Dev Environment Setup, How to run test, Build Instructions, Run instructions.
11. Test Driven Development (this should show clear pattern in the commit log: red, green, commit; refactor commit;)

## Prerequisites

- Golang - 1.14

## Build Instructions

```shell script
git clone https://github.com/littlestar642/clinne.git
cd clinne
go build cmd/cli/clinne.go
```

## Run Instructions

```shell script
clinne <command>
```

## Usage Instructions

Once you have built the executable you can run the application using your command line. 

For a quick walkthrough see this video:

[![IMAGE ALT TEXT HERE](https://img.youtube.com/vi/ct2-Fyj1B44/0.jpg)](https://www.youtube.com/watch?v=ct2-Fyj1B44)

## Known Problems

Presently you would not be able to install the application and use it directly from cli. You would need to cd into the folder to run it. The reason for that being the `docs/` folder in the project root which wont be present in your `$GOPATH`. 

## Test Instructions
TBD