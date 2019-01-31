package watchtower

import (
	//"google.golang.org/api/iterator"
	"cloud.google.com/go/firestore"

	//"github.com/julienschmidt/httprouter"
	"context"
	"log"
	"time"

	// firebase "firebase.google.com/go"
	//"firebase.google.com/go/auth"
	// "google.golang.org/api/option"

)

type database struct {
	store *firestore.Client
	ctx context.Context
}

func NewDatabase(store *firestore.Client, ctx context.Context) *database {
	return &database{store, ctx}
}

func (db *database) Update(w Watch) error {

	// currentTime := time.Now().UTC().Format("2006-01-02 15:04:05")
	// fmt.Println(currentTime)
	// // "lastUpdate": currentTime,
	// _, _, err := db.store.Collection("watches").Add(db.ctx, map[string]interface{}{
    //     "id": w.Id,
    //     "recordId":  w.RecordId,
	// 	"active":  w.Active,
	// 	"network": w.Network,
		
	// })
	// if err != nil {
	// 		log.Fatalf("Failed adding watch: %v", err)
	// 		return err
	// }

	// return nil

	currentTime := time.Now().UTC().Format("2006-01-02 15:04:05")
	// "lastUpdate": currentTime,
	_, err := db.store.Collection("watches").Doc(w.Id).Set(db.ctx, map[string]interface{}{
        "id": w.Id,
        "recordId":  w.RecordId,
		"active":  w.Active,
		"network": w.Network,
		"lastUpdate": currentTime,
		
	})
	if err != nil {
			log.Fatalf("Failed adding watch: %v", err)
			return err
	}

	return nil

	// currentTime := time.Now().UTC().Format("2006-01-02 15:04:05")
	// // "lastUpdate": currentTime,
	// _, err := firestore.Collection("watches").Doc("490").Set(ctx, map[string]interface{}{
    //     "id": "123",
    //     "recordId":  "233",
	// 	"active":  true,
	// 	"network": "sunny",
	// 	"lastUpdate": currentTime,
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("HERE")


}

// iter := client.Collection("users").Documents(ctx)
// for {
// 	doc, err := iter.Next()
// 	if err == iterator.Done {
// 		break
// 	}
// 	if err != nil {
// 		log.Fatalf("Failed to iterate: %v", err)
// 	}
// 	fmt.Println(doc.Data())
// }
// [END fs_get_all_users]