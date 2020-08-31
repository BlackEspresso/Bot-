# Bot++
A language making it simple to write chat bots.

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
Define the state with "trainingdata" and possible "responses". The example creates a bot with 2 states (feeling and greetings)
each state has some learning data and some possible replys. It looks limited now but with the tag-feature from [Lookout](#Lookout) complex conversations should't be problem.

## Get started
see main.b++ in editor and 
> go run main.go

Example output for previous example
```
Entering REPL, type exit to quit
how are you
>>  everything ok
hello
>>  nice to meet you
```

Example output for main.b++
```
Entering REPL, type exit to quit
do you know the time
>>  its $time (variables yet to come)
book a flight
>>  whats your destination?
hamburg
>>  ok i'll book it for tomorrow
thanks
>>  a pleasure
```

(the flight state in this example is very limited and needs a lot of the [Lookout](#Lookout) features to work nicely )

 ## Lookout

- language specs

- string interpolation (powershell like)
  ```
  set greetings.responses [
    "hello from $name"
    "hello, its a beautiful $timeNow.weekday"
  ]
  ```

- improve string selection logic
  - use fuzzy logic
  - use machine learning ?

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

- switch between edit and talk mode in bot++ console application. Live tune your bot while your're talking to it

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