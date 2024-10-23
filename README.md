<h1>Rule Engine Backend</h1>

<p>This is a Golang server along with Gofiber as the framework, this server features API's to create rules, combine rules, evaluate rules, and fetch rules from the database</p>
<be>

 - The entire core logic of rule creation using AST Nodes is implemented in the Golang server served with endpoints.
 - The AST nodes are stored in postgreSQL database hosted on AWS to attain data persistance.
 - The AST nodes are stored with child-parent relation ships in the database using specific schema.
  
 <h3>Hosting</h3>
 
 - The server is hosted on the Railway cloud provider under a free tier.

<h3>Get Started</h3>
1. Clone the Repository 

```bash
git clone https://github.com/Hemanth5603/Rule-Engine-Backend.git
```

2. Install all dependencies

```bash
go get .
```

3. Run the server

```bash
go run main.go
```

<h3>Implementation</h3>

Here is the AST Node structure
```bash
type RuleNode struct {
	ID         int    `db:"id"`
	NodeType   string `db:"node_type"`
	LeftChild  *int   `db:"left_child"`
	RightChild *int   `db:"right_child"`
	Attribute  string `db:"attribute"`
	Operator   string `db:"operator"`
	Value      string `db:"value"`
}

```


