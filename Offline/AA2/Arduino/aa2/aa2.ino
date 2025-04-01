#include <LiquidCrystal_I2C.h>
#include <string.h>

#define BUTTON_VANCE 6
#define BUTTON_START 7

#define VIRT_SCR_COLS 20
#define VIRT_SCR_ROWS 4

char g_virt_scr[VIRT_SCR_ROWS][VIRT_SCR_COLS + 1];

#define virt_scr_sprintf(x, y, fmt, ...) \
  snprintf(g_virt_scr[y] + x, (VIRT_SCR_COLS - x), fmt, __VA_ARGS__);

LiquidCrystal_I2C lcd(0x27, VIRT_SCR_COLS, VIRT_SCR_ROWS);

/*
| **Data**             | **Type** | **Description**                                            | **Target screen** | **Format**                                                                                                                       |
|----------------------|----------|------------------------------------------------------------|-------------------|----------------------------------------------------------------------------------------------------------------------------------|
| Tags                 | Int32    | Number of tags read by the RFID reader.                    | 1                 | %dK (K only included after 10K tags) |
| Unique-Tags          | Int32    | Number of Unique tags read by the RFID reader.             | 1                 | %d                                                                                                                               |
| Communication status | Bool     | Status of the connection between the PC and mytempo.esp.br | 2                 | SIM(True) / NÃO(False)                                                                                                           |
| WI-FI status         | Bool     | Status of the general internet connection of the device.   | 3                 | OK(True) / X(false)                                                                                                              |
| 4G status            | Bool     | Status of the LTE/4G connection of the device.             | 3                 | OK(True) / X(false)                                                                                                              |
| Reader status        | Bool     | Status of the RFID reader.                                 | 4                 | OK(True) / X(false)                                                                                                              |
| System version       | Int32    | Version number of the system.                              | 5                 | %d                                                                                                                               |
| Backup count         | Int32    | Number of backups currently stored.                        | 6                 | %d                                                                                                                               |
| Envio count          | Int32    | Number of envios currently stored.                         | 7                 | %d                                                                                                                               |
*/
typedef struct __attribute__((packed)) PCData
{
  int32_t tags       ;
  int32_t unique_tags;
  bool    comm_status;
  bool    wifi_status;
  bool    lte4_status;
  bool    rfid_status;
  int32_t sys_version;
  int32_t backups    ;
  int32_t envios     ;
} PCData;

constexpr size_t pc_data_size = sizeof(PCData);

/*
| **Screen**              | **Content**                                      | **Description**                                                       | **Action (optional)**                                                                       |
|-------------------------|--------------------------------------------------|-----------------------------------------------------------------------|---------------------------------------------------------------------------------------------|
| Info                    | - Regist.: (Tags)<br>- Atletas: (Unique-Tags)    | Displays the<br>tag count, and<br>unique tag count.                   | START: Reset visual tag data<br>information, does not<br>actually touch anything<br>stored. |
| Network                 | - Comunicando: (Communication status)            | Displays the communication<br>with mytempo.esp.br                     |                                                                                             |
| Network Mgmt            | - Wi-Fi: (WI-FI status)<br>- LTE/4G: (4G status) | Displays basic PC network<br>connectivity info.                       | START: Issue a reconnection<br>of both wifi and 4g networks.                                |
| Reader info             | - Leitor: (Reader status)                        | Displays RFID reader<br>connectivity info.                            |                                                                                             |
| System                  | - Version: (System version)                      | Displays the system version,<br>i.e. the current update.              | START: Fetch and Install the<br>latest version from github.                                 |
| Upload                  | - Regist.: (Tags)<br>- Pendentes: (Envio count)  | Displays the current tag<br>count + the number of pending<br>uploads. | START: Upload all tag data<br>currently stored + pending<br>tag data.                       |
| Upload (Backup)         | - Backups: (Backup count)                        | Displays the number of backups.                                       | START: Upload all backups.                                                                  |
| #15 (Erase data)        |                         -                        |                                   -                                   |                                              -                                              |
| #15 (Shutdown)          |                         -                        |                                   -                                   |                                              -                                              |
| #15 (Shutdown [Helper]) |                         -                        |                                   -                                   |                                              -                                              |
*/
// Screens

void
handleButtons()
{
//  (digitalRead(BUTTON_VANCE) == 0) && screen_advance();
//  (digitalRead(BUTTON_START) == 0) && event_send_start();
}

#define START_DELIMITER 0x3C
#define END_DELIMITER 0x3E
void
handleSerial()
{  
  char buf[pc_data_size];

  if (Serial.available() < pc_data_size + 2) return;
  if (Serial.read() != START_DELIMITER) goto consume_buf;

consume_buf:
  Serial.readBytes(pc_data_size + 1);
}

void
draw()
{
  for (int i = 0; i < VIRT_SCR_ROWS; i++){

    lcd.setCursor(0, i);

    for (char* c = g_virt_scr[i], i = 0; *c != '\0' && i < VIRT_SCR_COLS; c++, i++)
      lcd.write(*c);
  }
}

void
setup()
{
  lcd.init();      // Initialize the LCD
  lcd.backlight(); // Turn on the backlight
  
  memset(g_virt_scr, '\0', sizeof(g_virt_scr));

  Serial.begin(9600);
  while(!Serial);

  pinMode(BUTTON_START, INPUT_PULLUP);
  pinMode(BUTTON_VANCE, INPUT_PULLUP);
}

void
loop()
{
  handleButtons();
  handleSerial();
  //draw();
}
