package api

import (
    "time"
)

type ResourceError struct {
    code int32 `json:"code"`
    message string `json:"message"`
}

type ResourceDeletedResponse struct {
    _ interface{} `json:"_"`
}

type Opportunity struct {
    id string `json:"id"`
    source string `json:"source"`
    title string `json:"title"`
    description string `json:"description"`
    fundingDetails FundingDetails `json:"fundingDetails"`
    applicationTimeline *[]Event `json:"applicationTimeline,omitempty"`
    customFields *map[string]CustomField `json:"customFields,omitempty"`
}

type FundingDetails struct {
    totalAmountAvailable *Currency `json:"totalAmountAvailable,omitempty"`
    minAwardAmount *Currency `json:"minAwardAmount,omitempty"`
    maxAwardAmount *Currency `json:"maxAwardAmount,omitempty"`
    minAwardCount *int64 `json:"minAwardCount,omitempty"`
    maxAwardCount *int64 `json:"maxAwardCount,omitempty"`
    estimatedAwardCount *int64 `json:"estimatedAwardCount,omitempty"`
}

type Currency struct {
    amount float64 `json:"amount"`
    display string `json:"display"`
    currencyCode string `json:"currencyCode"`
}

type Event struct {
    name string `json:"name"`
    date time.Time `json:"date"`
    time *time.Time `json:"time,omitempty"`
    description *string `json:"description,omitempty"`
}

type CustomField struct {
    name string `json:"name"`
    type string `json:"type"`
    schema *string `json:"schema,omitempty"`
    value interface{} `json:"value"`
    description *string `json:"description,omitempty"`
    context *string `json:"context,omitempty"`
}
