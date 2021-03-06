package main

// These structs are for supplying page data to .tmpl files

// Draft describes a draft for the purposes of the index page.
type Draft struct {
	Name       string
	ID         int64
	Seats      int64
	Joined     bool
	Joinable   bool
	Replayable bool
}

// IndexPageData is the input to index.tmpl.
type IndexPageData struct {
	Drafts  []Draft
	ViewURL string
	UserID  int64
}

// ReplayPageData is the input to replay.tmpl.
type VuePageData struct {
	UserJSON string
}

// These structs are for replaying the draft on the client.

// Perspective tells the client from which user's perspective the replay data is from.
type Perspective struct {
	User  int64     `json:"user"`
	Draft DraftJSON `json:"draft"`
}

// DraftJSON describes the draft to the replay viewer.
type DraftJSON struct {
	DraftID   int64        `json:"draftId"`
	DraftName string       `json:"draftName"`
	Seats     [8]Seat      `json:"seats"`
	Events    []DraftEvent `json:"events"`
}

// Seat is part of DraftJSON.
type Seat struct {
	Packs       [3][15]interface{} `json:"packs"`
	PlayerName  string             `json:"playerName"`
	PlayerID    int64              `json:"playerId"`
	PlayerImage string             `json:"playerImage"`
}

// DraftEvent is part of DraftJSON.
type DraftEvent struct {
	Position       int64    `json:"position"`
	Announcements  []string `json:"announcements"`
	Cards          []int64  `json:"cards"`
	PlayerModified int64    `json:"playerModified"`
	DraftModified  int64    `json:"draftModified"`
	Round          int64    `json:"round"`
	Librarian      bool     `json:"librarian"`
	Type           string   `json:"type"`
}

// These structs are for sending other data to the client.

// JSONError helps to pass an error to the client when something breaks.
type JSONError struct {
	Error string `json:"error"`
}

// DraftList is turned into JSON and used for the REST API.
type DraftList struct {
	Drafts []DraftListEntry `json:"drafts"`
}

// DraftListEntry is turned into JSON and used for the REST API.
type DraftListEntry struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	AvailableSeats int64  `json:"availableSeats"`
	Status         string `json:"status"`
}

// UserInfo is JSON passed to the client.
type UserInfo struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
	ID      int64  `json:"userId"`
}

// These structs are for receiving data from the client.

// PostedPick is JSON accepted from the client when a user makes a pick.
type PostedPick struct {
	CardIds []int64 `json:"cards"`
}

// PostedJoin is JSON accepted from the client when a user joins a draft.
type PostedJoin struct {
	ID int64 `json:"id"`
}

// These structs are for exporting in bulk to .dek files.

// BulkMTGOExport is used to bulk export .dek files for the admin.
type BulkMTGOExport struct {
	PlayerID int64
	Username string
	Deck     string
}

// NameAndQuantity is used in MTGO .dek exports.
type NameAndQuantity struct {
	Name     string
	MTGO     string
	Quantity int64
}

// R38CardData is the JSON passed to the client for card data.
// Note that this does not describe everything that is in the data, just what we need
type R38CardData struct {
	MTGO     int64            `json:"mtgo_id"`
	Scryfall ScryfallCardData `json:"scryfall"`
}

// ScryfallCardData is more JSON passed to the client for card data.
// Note that this does not describe everything that is in the data, just what we need
type ScryfallCardData struct {
	Name string `json:"name"`
}
