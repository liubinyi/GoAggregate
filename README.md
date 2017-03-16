# GoAggregate
A go aggregate api boilerplate that connects all your core apis


## What's inside  
* swagger  
* dependency injection  
* tests  
* docker  
* CI  

### tests
* run unit tests  
```go test```  

* generate test coverage  
```go test -v -cover```  





## How does it work
                                                      _____________
                                                      |            |
                                                      |entity A api|
                                                      |____________|
                                                     /
                                                    /
                                   ______________  /
          client          DTO     |              |/          ______________
                   O     <--------|aggregate api |---------- |entity B api|
                  -|-             |______________|\          |____________|
                  / \                              \
                                                    \
                                                     \
                                                     _____________
                                                     |            |
                                                     |entity C api|
                                                     |____________|
