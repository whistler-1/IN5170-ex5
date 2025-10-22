# Exercise 1 

Implement an actor for a calculator that has two states (`DUAL` and `SNGL`). 

In state `DUAL`, it accepts messages of the following form. 
- Message (`from`, `ADD`, `n`, `m`) should send back `n`+`m` to the sending actor `from`
- Message (`from`, `SNGL`) should switch to state `SNGL`

In state `SNGL`, it accepts messages of the following form. 
- Message (`from`, `STORE`, `n`) should store value `n` in the memory of the actor 
- Message (`from`, `INC`, `n`) should send back the sum of `n` and the stored value 
    to the sending actor `from`
- Message (`from`, `DUAL`) should switch to state `DUAL`

### Actors
Info on actors here.


```shell



```