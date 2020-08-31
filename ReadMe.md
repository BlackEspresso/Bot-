# Bot++
A language for writing chat bots.

add a state, some training data and responses
```
add state feeling
set feeling.trainingdata [
   "how are you?"
   "how do you do?"
]
set feeling.responses [
    "i'm fine thanks"
    "everything ok"
]

add state greetings
set greetings.trainingdata [
   "hello"
   "hi"
   "hallo"
   "good morning"
]
set greetings.responses [
   "nice to meet you"
   "hello"
   "hallo"
]
```

run main.go with

> go run