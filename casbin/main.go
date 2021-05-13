package main

import (
	"fmt"
	"github.com/casbin/casbin/v3"
)

func main() {
	/*// gorm adapter
	dsn := "root:root@tcp(127.0.0.1:3306)/casbin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(err)
	}


	filter := gormadapter.Filter{
		PType: []string{"p", "g"},
		V1: []string{"domain1"},
	}
	*/

	e, err := casbin.NewSyncedEnforcer("config/rbac_with_domains_model.conf", "policy/rbac_with_domains_policy.csv", true)
	if err != nil {
		panic(err)
	}

	//e.EnableAutoSave(true)

	//e.StartAutoLoadPolicy(1 * time.Second)

	// ---------------Watcher--------------
	//w, _ := rediswatcher.NewWatcher("127.0.0.1:6379")
	//_ = e.SetWatcher(w)
	//
	//err = w.SetUpdateCallback(func(s string) {
	//	fmt.Println("update policy from watcher")
	//	_ = e.LoadPolicy()
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//go func() {
	//	e2, err := casbin.NewSyncedEnforcer("config/rbac_with_domains_model.conf", "policy/rbac_with_domains_policy.csv", true)
	//	if err != nil {
	//		panic(err)
	//	}
	//	w, _ := rediswatcher.NewWatcher("127.0.0.1:6379")
	//	_ = e2.SetWatcher(w)
	//
	//	err = w.SetUpdateCallback(func(s string) {
	//		fmt.Println("update policy from watcher for enforcer 2")
	//		fmt.Println("string receive: "+ s)
	//		_ = e2.LoadPolicy()
	//	})
	//	if err != nil {
	//		panic(err)
	//	}
	//}()

	//time.Sleep(2*time.Second)

	/*//custom function
	e.AddFunction("build_role_tuple", func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) != 2 {
			return nil, errors.New("invalid number of arguments")
		}
		arg0 := arguments[0].(string)
		arg1 := arguments[1].(string)
		return fmt.Sprintf("%s::%s", arg0, arg1), nil
	})
	*/

	//
	//ok, err := e.Enforce("alice","domain1", "data1", "read")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(ok)
	fmt.Println(e.GetRolesForUserInDomain("bob", "domain2"))
	fmt.Println(e.GetRolesForUserInDomain("alice", "domain2"))

	ok, err := e.RemoveGroupingPolicies([][]string{{"bob", "admin", "domain2"},{"alice", "admin", "domain2"}})
	fmt.Println(ok)
	e.AddPolicies()

	fmt.Println(e.GetRolesForUserInDomain("bob", "domain2"))
	fmt.Println(e.GetRolesForUserInDomain("alice", "domain2"))
	//e.GetModel()

}
