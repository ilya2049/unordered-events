# Unordered streaming updates of entity attributes

## Domain

``` plantuml
@startuml

object Cat
Cat : age = 3

object CatOwner
CatOwner : catAge = 3

Cat "1" o-- "0..*" CatOwner

@enduml
```

## Infrastructure

The kafka topic plays the role of an event bus.

``` plantuml
@startuml

[Event Bus] as B

[Microservice] -> [B] : published events
[B] -> [Microservice] : events to handle

note top of B
  The postgres outbox 
  + 
  the kafka topic
end note

@enduml
```

A more detailed view of the components.

``` plantuml
@startuml

[Event Publisher] as P
[Event Handler] as H
[Event Bus] as B

interface "Postgres" as PG

[P] -> [B] : single updates
[B] -> [H] : single updates
[H] -> [B] : batch updates
[PG] - [P]
[PG] - [H]

@enduml
```
