package models

type FeeTransaction struct {
	ID                     int    `json:"id,omitempty"`
	TransactionAmount      int    `json:"transaction_amount,omitempty"`
	GlobalStartingBalance  int    `json:"global_starting_balance,omitempty"`
	TransactionDate        string `json:"transaction_date,omitempty"`
	TransactionDescription string `json:"transaction_description,omitempty"`
	GroupTransactionID     int    `json:"group_transaction_id,omitempty"`
	PaymentMethod          string `json:"payment_method,omitempty"`
	SchoolID               int    `json:"school_id,omitempty"`
	YearID                 int    `json:"year_id,omitempty"`
	StartingBalance        int    `json:"starting_balance,omitempty"`
	GlobalNetBalance       int    `json:"global_net_balance,omitempty"`
	TransactionType        string `json:"transaction_type,omitempty"`
	PaymentReferenceNumber int    `json:"payment_reference_number,omitempty"`
}

type Fee struct {
	ID                  int            `json:"id,omitempty"`
	FeeAmount           int            `json:"fee_amount,omitempty"`
	FeeBalance          int            `json:"fee_balance,omitempty"`
	FeeDcscription      string         `json:"fee_dcscription,omitempty"`
	DateCreated         string         `json:"date_created,omitempty"`
	DateModified        string         `json:"date_modified,omitempty"`
	TransactionDate     string         `json:"transaction_date,omitempty"`
	DepartmentName      string         `json:"department_name,omitempty"`
	Priority            int            `json:"priority,omitempty"`
	ProRatableIndicator int            `json:"pro_ratable_indicator,omitempty"`
	GroupTransactionID  int            `json:"group_transaction_id,omitempty"`
	SchoolID            int            `json:"school_id,omitempty"`
	SchoolFeeID         int            `json:"school_fee_id,omitempty"`
	TermID              int            `json:"term_id,omitempty"`
	YearID              int            `json:"year_id,omitempty"`
	CourseName          string         `json:"course_name,omitempty"`
	CourseNumber        int            `json:"course_number,omitempty"`
	CategoryName        string         `json:"category_name,omitempty"`
	TypeID              int            `json:"type_id,omitempty"`
	TypeName            string         `json:"type_name,omitempty"`
	FeeTransaction      FeeTransaction `json:"fee_transaction,omitempty"`
}

type Fees struct {
	Fee []Fee `json:"fee,omitempty"`
}
