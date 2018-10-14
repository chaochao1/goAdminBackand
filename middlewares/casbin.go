package middlewares

import (
	"github.com/casbin/casbin"
	"github.com/casbin/xorm-adapter"
)

var (
	Enforce *casbin.Enforcer
)

func init() {
	a := xormadapter.NewAdapter("mysql", "root:@tcp(127.0.0.1:3306)/") // Your driver and data source.

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	e := casbin.NewEnforcer("casbin/rbac_model.conf", a)

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	e.Enforce("alice", "data1", "read")

	// Modify the policy.
	e.AddPolicy("alice", "data1", "read")
	e.RemovePolicy("alice", "data1", "read")

	// Save the policy back to DB.
	e.SavePolicy()
	Enforce = e
}
