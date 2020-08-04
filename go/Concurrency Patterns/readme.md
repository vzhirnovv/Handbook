# Concurrency patterns

## Overview

Link: [Google I/O 2012 - Go Concurrency Patterns](http://www.youtube.com/watch?v=f6kdp27TYZs)

This directory contains common concurrency patterns for Go and information on common concurrency terms and patterns using goroutines and channels.

## Common multithreading definitions

- **Process** - the UNIX environment set up to run a program.

- **Thread** - a sequence of instructions executed within the context of a process.

- **Single-threaded** - only a single thread can be accessed, limiting execution to sequential processing on one thread only.

- **Multithreaded** - multiple threads can be accessed, execution occurs via parallel or concurrent processing.

- **Mutual exclusion lock (mutex)** - an object used to lock and unlock access to shared data.

- **Condition variables** - objects used to block threads until a change of state.

- **Read-write locks** - objects used to allow multiple read-only access to data, but maintain _exclusive_ access to modification of data.

- **Parallelism** - a condition that arises when at least two threads are _executing_ simultaneously.

- **Concurrency** - a condition that arises when at least two threads are _making progress_. A more general form of parallelism.

- **Goroutine (Go)** - a lightweight thread managed by the Go runtime. It has it's own automatically resized stack and operates concurrently with the main Go runtime and other goroutines. Many goroutines can be run within a runtime efficiently.

- **Channels (Go)** - typed conduits through which data can be sent. By default, sends and recieves block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

- **Lockstep** - a condition which occurs when two channels are sending data one after the other. In this case, the second channel will not send data until data from the first channel is recieved, then the first channel will wait until the second channel sent data, and so on.

## Generators

Channels are _first-class values_, like integers or strings, which allows us to create functions that return channels. This, in turn, allows us to separate goroutines and channels from the main functions, making our code more concise.

## Multiplexing

We can create a generator that wraps multiple generators and aggregates their channels. This allows us to avoid lockstep, allowing channels to send data immediately upon being requested.

## Restoring sequencing

Pass a channel to a channel, making a gorouting wait its turn. Recieve all messages, then enable them again by sending on a private channel.

## Select statement (Go)

A **select statement** is a control structure, somewhat like a switch statement that lets you control the behaviour of your program based on what communications can proceed at a given moment.

Flow:

- All channels are evaluated
- Selection blocks until one communication can proceed, which it then does.
- If multiple can proceed, select chooses pseudo-randomly
- A default clause executes immediately if no channel is ready (won't proceed if no default clause is present).
