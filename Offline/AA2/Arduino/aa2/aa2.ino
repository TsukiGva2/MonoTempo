#include <string.h>

#define __STDC_FORMAT_MACROS 
#include <inttypes.h>

#define BUTTON_VANCE 6
#define BUTTON_START 7

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

PCData g_system_data;

static void
parseData(const char* data)
{
  sscanf(data, "%" PRId32 ";%" PRId32 ";%d;%d;%d;%d;%" PRId32 ";%" PRId32,
              &g_system_data.tags,
              &g_system_data.unique_tags,
              &g_system_data.comm_status,
              &g_system_data.wifi_status,
              &g_system_data.lte4_status,
              &g_system_data.rfid_status,
              &g_system_data.sys_version,
              &g_system_data.backups,
              &g_system_data.envios);
}

/* SCREEN_H */
#include <LiquidCrystal_I2C.h>
#include <string.h>

#define VIRT_SCR_COLS 20
#define VIRT_SCR_ROWS 4

char g_virt_scr[VIRT_SCR_ROWS][VIRT_SCR_COLS + 1];

LiquidCrystal_I2C lcd(0x27, VIRT_SCR_COLS, VIRT_SCR_ROWS);

#define virt_scr_sprintf(x, y, fmt, ...) \
  snprintf(g_virt_scr[y] + x, (VIRT_SCR_COLS - x), fmt, __VA_ARGS__);

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
#define SCREENS_COUNT 11
unsigned int g_current_screen = 0;

const char desc[SCREENS_COUNT][VIRT_SCR_COLS] = {
  "<START: Reset tela>",
  "",
  "<START: Reconectar>",
  "",
  "<START: Atualizar>",
  "<START: Envia Regist>",
};

void
screen_build()
{
  virt_scr_sprintf(0, 0, "PORTAL my50x", 0);

  switch (g_current_screen) {
    case 0:
      break;
  }

  virt_scr_sprintf(0, 3, "%s", desc[g_current_screen]);
}

void
screen_draw()
{
  screen_build();

  for (int i = 0; i < VIRT_SCR_ROWS; i++) {
    lcd.setCursor(0, i);
    for (char* c = g_virt_scr[i], i = 0; *c != '\0' && i < VIRT_SCR_COLS; c++, i++)
      lcd.write(*c);
  }
}

void
screen_init()
{
  lcd.init();      // Initialize the LCD
  lcd.backlight(); // Turn on the backlight
  memset(g_virt_scr, '\0', sizeof(g_virt_scr));
}

void
handleButtons()
{
  if (digitalRead(BUTTON_VANCE) == 0)
    g_current_screen = (g_current_screen + 1) % SCREENS_COUNT;

  //(digitalRead(BUTTON_START) == 0) && event_send_start();
}

#define START_DELIMITER 0x3C
#define END_DELIMITER 0x3E
#define CAP 256
void
handleSerial()
{  
  char buf[CAP];
  char imm[CAP];
  int c, i;

  while (Serial.available() > 0)
    if (Serial.read() == START_DELIMITER) goto read_delimited;
  
  return;

read_delimited:
  for (i = 0; i < CAP && Serial.available() > 0 && (c = Serial.read()) != END_DELIMITER; i++) imm[i] = c;
  if (c != END_DELIMITER) return;

  memcpy(buf, imm, i);
  buf[i] = '\0';

  parseData(buf);
}

void
setup()
{
  screen_init();

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
  screen_draw();
  delay(200);
}
