# Go Lang's Errpr handling

- Actually, In go lang, we need to imply interface for error. and handle error in all layers

- go's error is not work like exception, unhandled error can occur big problem our program

- nomally, it is normal that log the error when error is occured.

- also, the log should return stacked function's information

go lang's error interface is so simple;

```
type Error interface {
    Error() string
}
```

## What will we do?

- Error handling and interface

- use pkg/errors and wrap error

- use log package and understanding for logging timing

- structured logging using apex and logrus

- logging using context package

- use global variables

- resolve panic for long time continue process
