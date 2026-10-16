package entities

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/storedcredentialinitiator"
	"github.com/globalpayments/go-sdk/api/entities/enums/storedcredentialreason"
	"github.com/globalpayments/go-sdk/api/entities/enums/storedcredentialsequence"
	"github.com/globalpayments/go-sdk/api/entities/enums/storedcredentialtype"
)

type StoredCredential struct {
	Type      storedcredentialtype.StoredCredentialType
	Initiator storedcredentialinitiator.StoredCredentialInitiator
	Sequence  storedcredentialsequence.StoredCredentialSequence
	Reason    storedcredentialreason.StoredCredentialReason
	SchemeId  string
}

func (s *StoredCredential) GetType() storedcredentialtype.StoredCredentialType {
	return s.Type
}

func (s *StoredCredential) SetType(value storedcredentialtype.StoredCredentialType) *StoredCredential {
	s.Type = value
	return s
}

func (s *StoredCredential) GetInitiator() storedcredentialinitiator.StoredCredentialInitiator {
	return s.Initiator
}

func (s *StoredCredential) SetInitiator(value storedcredentialinitiator.StoredCredentialInitiator) *StoredCredential {
	s.Initiator = value
	return s
}

func (s *StoredCredential) GetSequence() storedcredentialsequence.StoredCredentialSequence {
	return s.Sequence
}

func (s *StoredCredential) SetSequence(value storedcredentialsequence.StoredCredentialSequence) *StoredCredential {
	s.Sequence = value
	return s
}

func (s *StoredCredential) GetReason() storedcredentialreason.StoredCredentialReason {
	return s.Reason
}

func (s *StoredCredential) SetReason(value storedcredentialreason.StoredCredentialReason) *StoredCredential {
	s.Reason = value
	return s
}

func (s *StoredCredential) GetSchemeId() string {
	return s.SchemeId
}

func (s *StoredCredential) SetSchemeId(value string) *StoredCredential {
	s.SchemeId = value
	return s
}
