package jmap

// A Mailbox represents a named set of Emails. This is the primary mechanism for
// organising Emails within an account. It is analogous to a folder or a label
// in other systems.
//
// See https://tools.ietf.org/html/rfc8621#section-2 for details.
type Mailbox struct {
	// The (immutable, server-set) id of the Mailbox.
	ID ID `json:"id"`

	// User-visible name for the Mailbox, e.g., "Inbox".
	Name string `json:"name"`

	// The Mailbox id for the parent of this Mailbox, or nil if this Mailbox is
	// at the top level.
	ParentID *ID `json:"parentId"`

	// Identifies Mailboxes that have a particular common purpose
	// (e.g., the "inbox"), regardless of the name property (which may be
	// localised).
	Role *string `json:"role"`

	// Defines the sort order of Mailboxes when presented in the client’s UI,
	// so it is consistent between devices.
	SortOrder UnsignedInt `json:"sortOrder"`

	// The number of Emails in this Mailbox.
	TotalEmails UnsignedInt `json:"totalEmails"`

	// The number of Emails in this Mailbox that have neither the $seen keyword
	// nor the $draft keyword.
	UnreadEmails UnsignedInt `json:"unreadEmails"`

	// The number of Threads where at least one Email in the Thread is in this
	// Mailbox.
	TotalThreads UnsignedInt `json:"totalThreads"`

	// An indication of the number of “unread” Threads in the Mailbox.
	UnreadThreads UnsignedInt `json:"unreadThreads"`

	// The set of rights (Access Control Lists (ACLs)) the user has in relation
	// to this Mailbox.
	MyRights *struct {
		// If true, the user may use this Mailbox as part of a filter in an
		// Email/query call, and the Mailbox may be included in the mailboxIds
		// property of Email objects.
		MayReadItems bool `json:"mayReadItems"`

		// The user may add mail to this Mailbox (by either creating a new
		// Email or moving an existing one).
		MayAddItems bool `json:"mayAddItems"`

		// The user may remove mail from this Mailbox (by either changing the
		// Mailboxes of an Email or destroying the Email).
		MayRemoveItems bool `json:"mayRemoveItems"`

		// he user may add or remove the $seen keyword to/from an Email.
		MaySetSeen bool `json:"maySetSeen"`

		// The user may add or remove any keyword other than $seen to/from an
		// Email.
		MaySetKeywords bool `json:"maySetKeywords"`

		// The user may create a Mailbox with this Mailbox as its parent.
		MayCreateChild bool `json:"mayCreateChild"`

		// The user may rename the Mailbox or make it a child of another Mailbox.
		MayRename bool `json:"mayRename"`

		// The user may delete the Mailbox itself.
		MayDelete bool `json:"mayDelete"`

		// Messages may be submitted directly to this Mailbox.
		MaySubmit bool `json:"maySubmit"`

	} `json:"myRights"`

	// Has the user indicated they wish to see this Mailbox in their client?
	IsSubscribed bool `json:"isSubscribed"`
}
