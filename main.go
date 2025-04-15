package main

func main() {
	todos := Todos{}
	todos.add("Купить молоко")
	todos.add("Купить хлеб")
	todos.toggle(0)
	todos.print()
}
