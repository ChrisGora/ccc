# :bookmark_tabs: Concurrency Crash Course

A quick introduction to concurrency in Go, for people familiar with Imperative or Object-Oriented programming.

Each lab is designed to take 2-3 hours depending on your previous experience. The first 3 labs have associated preparation lectures. It is recommended that you spend an hour before each lab, going through the lecture and playing with the provided code.

## Lab :one: - Introduction to imperative programming in Go

You will become familiar with Go's syntax and features. 

> This lab will assume that you already know at least one other Imperative or Object-oriented language. A basic understanding of pointers is also required.

## Lab :two: - Basic concurrency

This lab will teach you basics of concurrent computing. You will be introduced to channels and goroutines, which are based on the CSP message-passing model for concurrency. In the final question you will be given working serial code and asked to make it faster by parallelising it. 

> No previous experience of using threads or channels is required.

## Lab :three: - Advanced concurrency

You will get to experience more complicated patterns in concurrent programming. You will get to modify your quiz from Lab 1 to include a 5s timeout and you will investigate costs and benefits to performing merge sort using hundreds of goroutines.

## Lab :four: - Beyond channels

This final lab will focus on concurrency without using channels. The classic memory-sharing model is introduced, including POSIX-style semaphores, mutexes and condition variables. The final question of this lab asks you to write a Bank Simulator using both memory-sharing and channels, asking you to apply all the skills you have gathered throughout the course.

> Note that this lab does not yet have an associated lecture. You will need at least a basic understanding of mutexes and semaphores to succeed.
