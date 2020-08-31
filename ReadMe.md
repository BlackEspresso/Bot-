# Bot++
A language for writing chat bots.

simple define the state with trainingdata and responses
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

> go run main.go
```
Entering REPL, type exit to quit
who are you
>>  everything ok
hello
>>  nice to meet you
```

 ## Lookout
- string interpolation (powershell like)
  ```
  set greetings.responses [
    "hello from $name"
    "hello, its a beautiful $timeNow.weekday"
  ]
  ```

- input/output tags
  tags conversations, if the user triggered the greeting state
  he gets the tag "polite_user"
  ```
  set greetings.outputTag polite_user
  ```

  only if the user was polite enough to greet the charming context is considered as valid state (filtered by inputTag)
  ```
  set charming.inputTag polite-user
  ```

  This makes it possible to build complex conversations.
  e.g. ask for flights -> book flights -> list flights

- a simple web interface to manage your bot(s):
    - Easy Slack & MS Teams integration
    - see bot logs (messages per second, failes pers seconds)
    - log exceptions

- data types for trainingdata
  add entities for easy text regcognition
  ```
  add entity @weekday
  set @weekday.monday ["mondays","monday","mon"]
  set @weekday.tuesday ["tuesdays","tuesday","tues"]
  
  set greetings.trainingdata [
      "set alarm on %{mondays:@weekday:$weekday}"
  ]
  ```
  declares a variable $weekday of the type @weekday with the values:
  monday, tuesday

- string parsing for trainingdata prepares trainingdata to lookout for a username and sets the variable $username to bob if the user entered a similiar sentence
  ```
  set greetings.trainingdata [
      "hello my name is %{bob:@text:$username}"
  ]
  ```

- actions on state
  e.g. a http action:
  ```
  set feeling.action{
      http["https://myhost?flight=$bookedFlightNumber"]
  }
  ```
  sends a http request after on state enter. the http request return values can be used in the responses section