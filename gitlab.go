package main

import (
	"fmt"
	"log"

	"github.com/xanzy/go-gitlab"
)

//User を取得する関数です
func getUserList( client *gitlab.Client) []*gitlab.User {
	fmt.Println("User を取得します。")

	users,_,err := client.Users.ListUsers(nil)
	if err != nil {
		fmt.Println("User の取得に失敗しました。")
		fmt.Println(err)
	}

	return users
}

//Issue を取得する関数です
func getIssueList( client *gitlab.Client ) []*gitlab.Issue {
	fmt.Println("Issue を取得します。")
	want,_,err := client.Issues.ListIssues(nil)

	if err != nil {
		fmt.Println("Issue の取得に失敗しました。")
		fmt.Println(err)
	}

	return want
}

//Issue を新規作成する関数です
func createIssue ( client *gitlab.Client, issuename string ) {
	fmt.Println("新規 Issue を作成します。")

	newTask := &gitlab.CreateIssueOptions{
		Title: gitlab.String(issuename),
	}
	_, _, err := client.Issues.CreateIssue(2,newTask)
	if err != nil {
		log.Fatal(err)
	}
}

//Issue の Assignee を変更
func changeIssueAssignee ( client *gitlab.Client ) {
	fmt.Println("1 番目の Issue の Assignee を test1234 に変更します。")

	opt := &gitlab.ListUsersOptions{
		Username: gitlab.String("tetsuya"),
	}
	users,_,err := client.Users.ListUsers(opt)
	if err != nil {
		fmt.Println("User の取得に失敗しました。")
		fmt.Println(err)
	}

	ids := []int{users[0].ID}

	assignt := &gitlab.UpdateIssueOptions{
		AssigneeIDs: ids,
	}
	_, _, err = client.Issues.UpdateIssue(2, 1, assignt)

	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	git, err := gitlab.NewClient("TdSSKEHqCsFdBVEiXdTH",gitlab.WithBaseURL("http://192.168.3.22:9010/api/v4"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// User の取得
	list := getUserList(git)
	for index,element := range list {
		fmt.Println(index,":",element.Name)
	}

	// Issue の取得
	want := getIssueList(git)
	for index,element := range want {
		fmt.Println(index,":",element.Title)
	}

	// Issue の作成
	//createIssue(git,"gitlab.go")

	// Issue を変更 (Assignee を test1234 に変更)
	changeIssueAssignee(git)
}