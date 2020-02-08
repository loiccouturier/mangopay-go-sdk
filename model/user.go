package model

type User struct {
	// ID is the user's unique identification.
	ID string `json:"id"`
	// PersonType is the type of user.
	PersonType PersonType `json:"PersonType"`
	// Email is the person's email address (not more than 12 consecutive numbers) - must be a valid email.
	Email string `json:"Email"`
	// KYCLevel More info here.
	KYCLevel KYCLevelType `json:"KYCLevel"`
	// Tag is a custom data that you can add to this item.
	Tag *string `json:"Tag"`
	// CreationDate is when the item was created.
	CreationDate int64 `json:"CreationDate"`
}

type PersonType string

const (
	PersonNatural PersonType = "NATURAL"
	PersonLegal   PersonType = "LEGAL"
)

// KYCLevelType there are two levels of user verification, also called API levels: Light (default) and Regular.
// Light (default) verification requires basic information provided during the user creation process.
// Regular verification requires identity proof which allows users to handle unlimited funds freely.
type KYCLevelType string

const (
	KYCLevelLight   KYCLevelType = "LIGHT"
	KYCLevelRegular KYCLevelType = "REGULAR"
)

type NaturalUser struct {
	User
	// FirstName is the name of the user.
	FirstName string `json:"FirstName"`
	// LastName is the last name of the user.
	LastName string `json:"LastName"`
	// Address is the address.
	Address Address `json:"Address"`
	// Birthday is the date of birth of the user
	// be careful to set the right timezone (should be UTC)
	// to avoid 00h becoming 23h (and hence interpreted as the day before).
	Birthday int64 `json:"Birthday"`
	// Nationality is the user’s nationality. ISO 3166-1 alpha-2 format is expected.
	Nationality string `json:"Nationality"`
	// CountryOfResidence The user’s country of residence. ISO 3166-1 alpha-2 format is expected.
	CountryOfResidence string `json:"CountryOfResidence"`
	// Occupation is the User’s occupation, ie. Work.
	Occupation *string `json:"Occupation"`
	// Could be only one of these values:
	//     1 - for incomes <18K€),
	//     2 - for incomes between 18 and 30K€,
	//     3 - for incomes between 30 and 50K€,
	//     4 - for incomes between 50 and 80K€,
	//     5 - for incomes between 80 and 120K€,
	//     6 - for incomes >120K€.
	IncomeRange     *int    `json:"IncomeRange"`
	ProofOfAddress  *string `json:"ProofOfAddress"`
	ProofOfIdentity *string `json:"ProofOfIdentity"`
}

type NaturalUserCreate struct {
	// FirstName is the name of the user.
	FirstName *string `json:"FirstName,omitempty"`
	// LastName is the last name of the user.
	LastName *string `json:"LastName,omitempty"`
	// Address is the address.
	Address *Address `json:"Address"`
	// Birthday is the date of birth of the user
	// be careful to set the right timezone (should be UTC)
	// to avoid 00h becoming 23h (and hence interpreted as the day before).
	Birthday *int64 `json:"Birthday,omitempty"`
	// Nationality is the user’s nationality. ISO 3166-1 alpha-2 format is expected.
	Nationality *string `json:"Nationality,omitempty"`
	// CountryOfResidence The user’s country of residence. ISO 3166-1 alpha-2 format is expected.
	CountryOfResidence *string `json:"CountryOfResidence,omitempty"`
	// Occupation is the User’s occupation, ie. Work.
	Occupation *string `json:"Occupation,omitempty"`
	// Could be only one of these values:
	//     1 - for incomes <18K€),
	//     2 - for incomes between 18 and 30K€,
	//     3 - for incomes between 30 and 50K€,
	//     4 - for incomes between 50 and 80K€,
	//     5 - for incomes between 80 and 120K€,
	//     6 - for incomes >120K€.
	IncomeRange *int `json:"IncomeRange,omitempty"`
	// Email The person's email address (not more than 12 consecutive numbers) - must be a valid email
	Email *string `json:"Email,omitempty"` // REQUIRED
}

type NaturalUserUpdate struct {
	NaturalUserCreate
}

type LegalPersonType string

const (
	LegalPersonBusiness     LegalPersonType = "BUSINESS"
	LegalPersonOrganization LegalPersonType = "ORGANIZATION"
	LegalPersonSoletrader   LegalPersonType = "SOLETRADER"
)

func (lpt LegalPersonType) String() string {
	return string(lpt)
}

type LegalUserCreate struct {
	// HeadquartersAddress is the address of the company’s headquarters. This field is mandatory to accept payout (More info here).
	HeadquartersAddress *Address `json:"HeadquartersAddress,omitempty"`
	// LegalPersonType is the type of legal user.
	LegalPersonType LegalPersonType `json:"LegalPersonType,omitempty"`
	// Name is the name of the legal user.
	Name *string `json:"Name,omitempty"`
	// LegalRepresentativeAddress is the address of the company’s Legal representative person.
	LegalRepresentativeAddress *Address `json:"LegalRepresentativeAddress,omitempty"`
	// LegalRepresentativeBirthday is the date of birth of the company’s Legal representative person.
	//     - be careful to set the right timezone (should be UTC) to avoid 00h becoming 23h (and hence interpreted as the day before).
	LegalRepresentativeBirthday *int64 `json:"LegalRepresentativeBirthday,omitempty"`
	// LegalRepresentativeCountryOfResidence is the country of residence of the company’s Legal representative person
	LegalRepresentativeCountryOfResidence *string `json:"LegalRepresentativeCountryOfResidence,omitempty"`
	// LegalRepresentativeNationality is the nationality of the company’s Legal representative person
	LegalRepresentativeNationality *string `json:"LegalRepresentativeNationality,omitempty"`
	// LegalRepresentativeEmail is the email of the company’s Legal representative person - must be a valid
	LegalRepresentativeEmail *string `json:"LegalRepresentativeEmail,omitempty"`
	// LegalRepresentativeFirstName is the firstname of the company’s Legal representative person
	LegalRepresentativeFirstName *string `json:"LegalRepresentativeFirstName,omitempty"`
	// LegalRepresentativeLastName is the lastname of the company’s Legal representative person
	LegalRepresentativeLastName *string `json:"LegalRepresentativeLastName,omitempty"`
	// Email is the person's email address (not more than 12 consecutive numbers) - must be a valid email.
	Email *string `json:"Email,omitempty"`
	// CompanyNumber is the official registered number of the business.
	// This field is mandatory to accept payout (More info here).
	// You can find the patterns and names in your local country here.
	CompanyNumber *string `json:"CompanyNumber,omitempty"`
}

type LegalUserUpdate struct {
	LegalUserCreate
}
