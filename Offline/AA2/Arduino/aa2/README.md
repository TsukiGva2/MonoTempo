# Build Instructions

## Prerequisites

### 1. **Arduino IDE**: Install the Arduino IDE from [Arduino's official website](https://www.arduino.cc/en/software)

### 2. **Libraries**: Ensure the following libraries are installed in the Arduino IDE

- `SafeString`
- `LiquidCrystal_I2C`

### 3. **Hardware**

- Arduino Mega 2560 board.
- 20x4 LCD screen with I2C interface.
- Two buttons connected to digital pins (`BUTTON_VANCE` and `BUTTON_START`).
- RFID reader for tag data.

## Steps to Build and Upload

1. Open the Arduino IDE.
2. Load the `aa2.ino` sketch file.
3. Select the correct board and port:

    - Board: `Arduino Mega 2560`
    - Port: `COM6` (or the port your Arduino is connected to).

4. Install the required libraries if not already installed:

    - Go to **Sketch > Include Library > Manage Libraries**.
    - Search for and install `SafeString` and `LiquidCrystal_I2C`.

5. Compile the sketch to ensure there are no errors.
6. Upload the sketch to the Arduino Mega 2560.

---

## Program Description

This program is designed to manage and display system information on an LCD screen, handle user inputs via buttons, and process data received through serial communication. It is primarily used for monitoring and controlling an RFID-based system.

### Key Features

- **LCD Display**: Displays system statuses, tag counts, network information, and other data.
- **Button Navigation**: Two buttons (`BUTTON_VANCE` and `BUTTON_START`) allow users to navigate between screens and perform actions.
- **Serial Communication**: Parses and processes data received via serial input.
- **Screen Locking**: Locks the screen for specific operations and provides confirmation prompts for critical actions.
- **System Actions**: Includes functionalities like uploading data, resetting, and shutting down.

### Screens

1. **Informational Screens**:

    - Display tag counts, network statuses, and system version.

2. **Action Screens**:

    - Allow actions like uploading data or resetting the system.

3. **Confirmation Screens**:

    - Prompt the user for confirmation before executing critical actions.

---

## Serial Communication Input Format

The program expects serial messages in the following format:

```perl
$MYTMP;<tags>;<unique_tags>;<comm_status>;<wifi_status>;<lte4_status>;<rfid_status>;<usb_status>;<sys_version>;<backups>;<permanent_unique_tags>;<timestamp>*<checksum>
```

### Field Descriptions

- **`<tags>`**: Total number of tags read (integer).
- **`<unique_tags>`**: Number of unique tags read (integer).
- **`<comm_status>`**: Communication status (`1` for true, `0` for false).
- **`<wifi_status>`**: Wi-Fi connection status (`1` for true, `0` for false).
- **`<lte4_status>`**: LTE/4G connection status (`1` for true, `0` for false).
- **`<rfid_status>`**: RFID reader status (`1` for true, `0` for false).
- **`<usb_status>`**: USB connection status (`1` for true, `0` for false).
- **`<sys_version>`**: System version number (integer).
- **`<backups>`**: Number of backups stored (integer).
- **`<permanent_unique_tags>`**: Number of permanent unique tags (integer).
- **`<timestamp>`**: Unix timestamp adjusted for the year 2000 (integer).
- **`<checksum>`**: XOR checksum of the message.

### Example

```perl
$MYTMP;12345;678;1;1;0;1;0;42;5;100;946684800*5A
```

---

## Usage Instructions

1. **Navigation**:

    - Use `BUTTON_VANCE` to navigate between screens.
    - Use `BUTTON_START` to trigger actions on the current screen.

2. **Serial Communication**:

     - Send data in the specified format to update system information.

3. **Monitor LCD**:

     - Observe the LCD for system statuses, prompts, and confirmation requests.
