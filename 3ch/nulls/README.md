# use pointer in incoding and decoding, use SQL NullType

- When incode or decode object In go lang, if there is not explicit setting type, the type is set to default

- example : string -> "" int -> 0

- But if the default value have some means, that occur some problems

- also when we use structure tag like json omitempty, '0' is ignore althogh 0 is vaild value
