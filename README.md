# golangprojektnovember22

project description:
this is a small golang project i wrote a cowsay (with a pinguin instead of a cow). the program prints a function that "draws" a pinguin. one can scan input from the terminal via a scan function. the size of the speaking bubble of the pinguin is dependend on the lenght of the input from the scanner.

project:
we were asked to create the code for a cowsay alternative written in golang ( cowsay = a "drawing" of a cow or any other any other animal integrated in the golangcode with ascii code) that can scan input from the terminal and print the given input in a speechbubble that varies in size dependend on the size of the given input from the scanner.

goals of the project:
this code was written as part of an assignment of a cloud and devops course (techstarter course). 
the goal of this assignment was to learn golang basics, the use of golang documentation, projectdocumentation and gaining more competence in using github.

table of contents:
this repository contains the main.go file with the sourcecode (name of the file is pinguinsay1.go), my projectdocumentation and a readme.md

structure of the code:
-package main
-import needed packages
-main function:
main function contains:
os.Args as a commmand line argument
-function with the ascii code for the "drawing" 


how to run the project:
i wrote and run the project in vscode.
for running the project download the code from this repository, open the file or the copied code in vscode or any other sourcecode editor and run the program with go run pinguinsay1.go (pinguinsay1.go = name of the project).

used sources:
https://go.dev/doc/
https://freshman.tech/snippets/go/read-console-input/
https://github.com/flaviocopes/gocowsay
https://github.com/AnonymousAAArdvark/cowsay-go
https://github.com/dhruvbird/go-cowsay
https://github.com/ShumilaMalik/Cowsay-in-Golang
https://gobyexample.com/command-line-arguments
https://yourbasic.org/golang/gotcha-index-out-of-range/


personal notes and impressions:
a small and not very complex project like this is - in my opinion - good for learning basic golang structure as well as gaining confidence in the usage of basic golang functions.

what i learned:
- thinking about how to create the code to fulfill the project requirements.
- the implementation of a divide and conquer solution finding strategy to break down the requirements into smaller goals
- the implementation of a divide and conquer code building strategy to find solutions for the smaller goals in order to find a solution for the big goal (fullfiling the requirements)
- strengthening my understanding of the basic golang structure and basic golang functions
- how to use scanner function in golang
- why the scanner is not the required tool for the given requirements and why the os.Args command line arguments is better to fulfill the requirements
- -how to use os.Args for implementing command line arguments
