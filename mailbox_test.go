package jmap

import (
	"encoding/json"
	"testing"

	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
)

var mailboxBlob = `{
    "id": "MB23cfa8094c0f41e6",
    "name": "Inbox",
    "parentId": null,
    "role": "inbox",
    "sortOrder": 10,
    "totalEmails": 16307,
    "unreadEmails": 13905,
    "totalThreads": 5833,
    "unreadThreads": 5128,
    "myRights": {
      "mayAddItems": true,
      "mayRename": false,
      "maySubmit": true,
      "mayDelete": false,
      "maySetKeywords": true,
      "mayRemoveItems": true,
      "mayCreateChild": true,
      "maySetSeen": true,
      "mayReadItems": true
    },
    "isSubscribed": true
}`

func TestMailboxUnmarshal(t *testing.T) {
	m := Mailbox{}
	assert.NilError(t, json.Unmarshal([]byte(mailboxBlob), &m), "json.Unmarshal")

	assert.Check(t, cmp.Equal(ID("MB23cfa8094c0f41e6"), m.ID))
	assert.Check(t, cmp.Nil(m.ParentID))
	assert.Check(t, cmp.Equal(m.MyRights.MayAddItems, true))
}

func TestMailboxMarshal(t *testing.T) {
	m := Mailbox{}
	assert.NilError(t, json.Unmarshal([]byte(mailboxBlob), &m), "json.Unmarshal")

	blob, err := json.MarshalIndent(m, "", "  ")
	assert.NilError(t, err, "json.Marshal")

	var original, remarshaled map[string]interface{}
	assert.NilError(t, json.Unmarshal([]byte(mailboxBlob), &original))
	assert.NilError(t, json.Unmarshal(blob, &remarshaled))

	assert.Check(t, cmp.DeepEqual(original, remarshaled))
}
