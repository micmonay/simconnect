package simconnect

// EasySimConnect use for best easy use SimConnect in golang
type EasySimConnect struct {
	sc *SimConnect
}

// NewEasySimConnect create instance of EasySimConnect
func NewEasySimConnect() (*EasySimConnect, error) {
	sc, err := NewSimConnect()
	if err != nil {
		return nil, err
	}
	return &EasySimConnect{
		sc,
	}, nil
}

// Connect to sim or return error
func (esc *EasySimConnect) Connect(appName string) error {
	err := esc.sc.Open(appName)
	if err != nil {
		return err
	}
	return nil
}

func (esc *EasySimConnect) ConnectStructToSimObject(s SimVar) {

}
