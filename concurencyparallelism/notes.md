<!-- Let’s now discuss how the Go scheduler works to overview how goroutines are han-
dled. Internally, the Go scheduler uses the following terminology (see http://mng.bz/
N611):
    - G—Goroutine
    - M—OS thread (stands for machine)
    - P—CPU core (stands for processor)

        Each OS thread (M) is assigned to a CPU core (P) by the OS scheduler. Then, each goroutine (G) runs on an M. The GOMAXPROCS variable defines the limit of Ms in
        charge of executing user-level code simultaneously. But if a thread is blocked in a system call (for example, I/O), the scheduler can spin up more Ms. As of Go 1.5, GOMAX-PROCS is by default equal to the number of available CPU cores.
        A goroutine has a simpler lifecycle than an OS thread. It can be doing one of the
        following:

                - Executing—The goroutine is scheduled on an M and executing its instructions.
                - Runnable—The goroutine is waiting to be in an executing state.
                - Waiting—The goroutine is stopped and pending something completing, such
                as a system call or a synchronization operation (such as acquiring a mutex).
-->
