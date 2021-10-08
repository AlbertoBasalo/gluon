# 2 - Control

## 2.1 For loop


```go
// Classic for loop with index
for i := 0; i < len(transactions); i++ {
	var transaction = transactions[i]
	fmt.Printf("%+v\n", transaction)
}
// For range loop can interate over slices
for index, transaction := range transactions {
	fmt.Printf("%+v\n", transaction)
}
```


## 2.2 Conditionals

```go
// simple conditionals
if transaction.transactionType == Deposit {
	balance += transaction.amount
} else {
	balance -= transaction.amount
}
// nested conditionals
if balance < 0 {
	fmt.Println("😖 You have no money")
} else if balance < 100 {
	fmt.Println("😃 You have money")
} else {
	fmt.Println("🤑 You have a lot of money")
}
```

## 2.3 Switch


### 2.3.1 With an initial expression

```go
switch transaction.transactionType {
case Deposit:
	balance += transaction.amount
case Withdrawal:
	fallthrough
case Transfer:
	balance -= transaction.amount
}
```

### 2.3.2 Without initialization


```go
switch {
case balance < 0:
	fmt.Println("😖 You have no money")
case balance < 100:
	fmt.Println("😃 You have money")
default:
	fmt.Println("🤑 You have a lot of money")
}
```

### [Back to index](https://github.com/AtomicBuilders/gluon/blob/main/docs/index.md)