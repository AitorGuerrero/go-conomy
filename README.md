# goConomy
A Go package to manage an economy system.

Initiallly developed for use in BadassCity project. Splitted from them to be reusable and mantain decoupling.

It exposes 4 structs:

- Merchant: A struct that can receive and give goods (e.g. a player trades with money)
- MoneyGenerator: A struct that can create goods (e.g. a bank creates money)
- MoneyDestroyer: A struct that can destroy goods (e.g. a storage consume money)
- Money: The basic good to trade.

In future develpment it would be possible to rade with different kind of goods (e.g. different currencies, or material goods) with different value relations. 
