package usb

type Device struct {
	Name      string
	FS        FileSystem
	IsMounted bool
}

func (d *Device) Check() (check bool, err error) {

	d.Name, check, err = CheckUSBStorageDevice(d.FS)

	return
}
