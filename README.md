# Ent slow autocomplete on mutations
Entgo mutation autocompletion in vscode are slow it takes around 300ms - 500ms to autocomplete


## How to replciate issue

- Open `main.go`
- Try to type `client.Item.Create().Mutation()` and then `.Set` it should show all the fields related to `Set`` so full statment will be this 
`client.Item.Create().Mutation().Set`