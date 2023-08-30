// Copyright 2019 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package inventory

// NewSchedulableLabels returns a new zero value instance of SchedulableLabels.
func NewSchedulableLabels() *SchedulableLabels {
	return &SchedulableLabels{
		Arc:     new(bool),
		Board:   new(string),
		Brand:   new(string),
		Callbox: new(bool),
		Capabilities: &HardwareCapabilities{
			Atrus:               new(bool),
			Bluetooth:           new(bool),
			Cbx:                 new(HardwareCapabilities_CbxState),
			CbxBranding:         new(HardwareCapabilities_CbxBranding),
			Detachablebase:      new(bool),
			Carrier:             new(HardwareCapabilities_Carrier),
			StarfishSlotMapping: new(string),
			Fingerprint:         new(bool),
			Flashrom:            new(bool),
			GpuFamily:           new(string),
			Graphics:            new(string),
			Hotwording:          new(bool),
			InternalDisplay:     new(bool),
			Lucidsleep:          new(bool),
			Modem:               new(string),
			Power:               new(string),
			Storage:             new(string),
			Telephony:           new(string),
			Touchpad:            new(bool),
			Touchscreen:         new(bool),
			Webcam:              new(bool),
		},
		Cr50Phase:     new(SchedulableLabels_CR50_Phase),
		Cr50RoKeyid:   new(string),
		Cr50RoVersion: new(string),
		Cr50RwKeyid:   new(string),
		Cr50RwVersion: new(string),
		EcType:        new(SchedulableLabels_ECType),
		HwidSku:       new(string),
		Model:         new(string),
		Sku:           new(string),
		Stability:     new(bool),
		OsType:        new(SchedulableLabels_OSType),
		Peripherals: &Peripherals{
			AudioBoard:               new(bool),
			AudioBox:                 new(bool),
			AudioCable:               new(bool),
			AudioLoopbackDongle:      new(bool),
			AudioboxJackpluggerState: new(Peripherals_AudioBoxJackPlugger),
			Camerabox:                new(bool),
			CameraboxFacing:          new(Peripherals_CameraboxFacing),
			CameraboxLight:           new(Peripherals_CameraboxLight),
			Chameleon:                new(bool),
			ChameleonState:           new(PeripheralState),
			Conductive:               new(bool),
			Huddly:                   new(bool),
			HmrState:                 new(PeripheralState),
			Mimo:                     new(bool),
			Servo:                    new(bool),
			ServoState:               new(PeripheralState),
			ServoType:                new(string),
			SmartUsbhub:              new(bool),
			Stylus:                   new(bool),
			TrrsType:                 new(Peripherals_TRRSType),
			Wificell:                 new(bool),
			Router_802_11Ax:          new(bool),
			WorkingBluetoothBtpeer:   new(int32),
		},
		Platform:        new(string),
		Phase:           new(SchedulableLabels_Phase),
		ReferenceDesign: new(string),
		WifiChip:        new(string),
		TestCoverageHints: &TestCoverageHints{
			ChaosDut:        new(bool),
			ChaosNightly:    new(bool),
			Chromesign:      new(bool),
			HangoutApp:      new(bool),
			MeetApp:         new(bool),
			RecoveryTest:    new(bool),
			TestAudiojack:   new(bool),
			TestHdmiaudio:   new(bool),
			TestUsbaudio:    new(bool),
			TestUsbprinting: new(bool),
			UsbDetect:       new(bool),
			UseLid:          new(bool),
		},
	}
}
