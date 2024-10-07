<!--

    # Load balancer design requirements
        - assess traffic patterns for your use case(bursty, seasonal, constant)
        - what is the nature of your traffic(HTTPS, UDP, TCP)?
        - what is the acceptable latency for your service?
    
    # Types of load balancers
        - Layer 4(transport layer): balances traffic based on IP and ports(TCP/UDP traffic, good for low-latency applications).
        - Layer 7(application layer): operates on much more granular level(https headers, cookies). useful for app based routing.
    
    # Load balancing algorithms
        - Round robin(sequential distribution of traffic across servers).
        - IP Hashing(hashing IP addresses and linking a client to a server via the hash).
        - Least connections(routing traffic to server with least amount of traffic).
        - Weighted round robin/least connections(routing traffic to servers with much higher capacites).
-->