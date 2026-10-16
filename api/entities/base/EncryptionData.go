package base

type EncryptionData struct {
	version     string
	trackNumber string
	ksn         string
	ktb         string
}

func (e *EncryptionData) GetVersion() string {
	return e.version
}

func (e *EncryptionData) SetVersion(version string) {
	e.version = version
}

func (e *EncryptionData) GetTrackNumber() string {
	return e.trackNumber
}

func (e *EncryptionData) SetTrackNumber(trackNumber string) {
	e.trackNumber = trackNumber
}

func (e *EncryptionData) GetKsn() string {
	return e.ksn
}

func (e *EncryptionData) SetKsn(ksn string) {
	e.ksn = ksn
}

func (e *EncryptionData) GetKtb() string {
	return e.ktb
}

func (e *EncryptionData) SetKtb(ktb string) {
	e.ktb = ktb
}

func EncryptionDataVersion1() *EncryptionData {
	rvalue := EncryptionData{}
	rvalue.SetVersion("01")
	return &rvalue
}

func EncryptionDataVersion2(ktb string, trackNumber ...string) *EncryptionData {
	rvalue := &EncryptionData{}
	rvalue.SetVersion("02")
	rvalue.SetKtb(ktb)
	if len(trackNumber) > 0 {
		rvalue.SetTrackNumber(trackNumber[0])
	}
	return rvalue
}

func EncryptionDataAdd(ksn string, trackNumber string) *EncryptionData {
	rvalue := &EncryptionData{}
	rvalue.SetTrackNumber(trackNumber)
	rvalue.SetKsn(ksn)
	return rvalue
}

func EncryptionDataSetKtbAndKsn(ktb string, ksn string) *EncryptionData {
	rvalue := &EncryptionData{}
	rvalue.SetKtb(ktb)
	rvalue.SetKsn(ksn)
	return rvalue
}

func EncryptionDataSetKSNAndEncryptedData(ktb string, ksn string) *EncryptionData {
	rvalue := &EncryptionData{}
	rvalue.SetKtb(ktb)
	rvalue.SetKsn(ksn)
	return rvalue
}
