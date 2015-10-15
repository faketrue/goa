package main

import "github.com/raphael/goa/examples/cellar/app"

// AccountController implements the account resource.
type AccountController struct {
	db *DB
}

// NewAccountController creates a account controller.
func NewAccountController() *AccountController {
	return &AccountController{db: NewDB()}
}

// Show retrieves the account with the given id.
func (b *AccountController) Show(c *app.ShowAccountContext) error {
	account := b.db.GetAccount(c.AccountID)
	if account == nil {
		return c.NotFound()
	}
	return c.OK(account, "default")
}

// Create records a new account.
func (b *AccountController) Create(c *app.CreateAccountContext) error {
	account := b.db.NewAccount()
	payload := c.Payload
	account.Name = payload.Name
	c.ResponseHeader().Set("Location", app.AccountHref(account.ID))
	return c.Created()
}

// Update updates a account field(s).
func (b *AccountController) Update(c *app.UpdateAccountContext) error {
	account := b.db.GetAccount(c.AccountID)
	if account == nil {
		return c.NotFound()
	}
	payload := c.Payload
	if payload.Name != "" {
		account.Name = payload.Name
	}
	b.db.SaveAccount(account)
	return c.NoContent()
}

// Delete removes a account from the database.
func (b *AccountController) Delete(c *app.DeleteAccountContext) error {
	account := b.db.GetAccount(c.AccountID)
	if account == nil {
		return c.NotFound()
	}
	b.db.DeleteAccount(account)
	return c.NoContent()
}
