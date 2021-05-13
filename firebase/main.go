package main

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	config, err := NewConfig()
	if err != nil {
		panic(err)
	}

	/*init firestore*/
	ctx := context.Background()
	configFirestore := &firebase.Config{
		ProjectID: config.ProjectIdFirestore,
	}
	app, err := firebase.NewApp(ctx, configFirestore, option.WithCredentialsJSON([]byte(config.FirebaseServiceAccountKeyJson)))

	if err != nil {
		panic(err)
	}

	clientFireStore, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}

	addedMembers := MemberFirestore{
		Uid:         "200000345",
		DisplayName: "Mới Tinh tinhf tinhf",
		ProfileUrl:  "",
	}
	_, err = clientFireStore.Collection("chat-group").Doc("group-channel-6402dasd0").Update(ctx, []firestore.Update{
		{
			Path:  "members",
			Value: addedMembers,
		},
	})
	if err != nil {
		panic(err)
	}
	//var groupInfo GroupFirestore
	//doc, err := clientFireStore.Collection("chat-group").Doc("group-channel-63684").Get(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = doc.DataTo(&groupInfo)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(groupInfo)
}

type MemberFirestore struct {
	Uid         string `firestore:"uid"`
	DisplayName string `firestore:"display_name"`
	ProfileUrl  string `firestore:"profile_url"`
	SeenAt      int    `firestore:"seen_at"`
	Age         int    `firestore:"age"`
}

type GroupFirestore struct {
	Name       string            `firestore:"name"`
	GroupId    string            `firestore:"group_id"`
	Type       string            `firestore:"type"`
	Members    []MemberFirestore `firestore:"members"`
	Removed    bool              `firestore:"removed"`
	CoverUrl   string            `firestore:"cover_url,omitempty"`
	CreatedAt  int64             `firestore:"created_at"`
	CreatedBy  MemberFirestore   `firestore:"created_by"`
	ModifiedAt int64             `firestore:"modified_at"`
}
