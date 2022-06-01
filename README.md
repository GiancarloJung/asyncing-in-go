# Process Oriented Programming

Design your program as a collection of independent processes
Design these process to eventually run in parallel
Design your code so that the outcome is always the same

Independent tasks
No race conditions and no deadlocks (same outcome)
Horizontally scalable

Process-oriented programming is a paradigm based on Communicating Sequential Processes, originally from a paper by Tony Hoare in 1977, this is also popularly called the actor model of concurrency. (formal language)

# Actor model

- Each process is build to run sequentially
- No shared state (data is communicated using channels)
- Scale by adding more

# OTP Supervisor

Each actor is a process.
Each process performs a specific task.
To tell a process to do something, you need to send it a message. The process can reply by sending back another message.
The kinds of messages the process can act on are specific to the process itself. In other words, messages are pattern-matched.
Other than that, processes don’t share any information with other processes.
