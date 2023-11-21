package main

type Account struct {
	ID               int64  `json:"id"`
	Email            string `json:"email"`
	Username         string `json:"username"`
	SubscriptionPlan string `json:"subscription_plan"`
}
