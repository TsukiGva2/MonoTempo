#include <string.h>

#define __STDC_FORMAT_MACROS 
#include <inttypes.h>

#define BUTTON_VANCE 6
#define BUTTON_START 7

bool g_clicked_vance = false;
bool g_clicked_start = false;

int
check_clicked()
{
  if (digitalRead(BUTTON_VANCE) == 0) {
    if (!g_clicked_vance) {
      g_clicked_vance = true;
      return BUTTON_VANCE;
    }
  } else
    g_clicked_vance = false;

  if (digitalRead(BUTTON_START) == 0) {
    if (!g_clicked_start) {
      g_clicked_start = true;
      return BUTTON_START;
    }
  } else
    g_clicked_start = false;

  return 0;
}

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
parse_data(const char* data)
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

const char fill_pattern[20] = "                    ";

#define virt_scr_sprintf(x, y, fmt, ...) \
  snprintf(g_virt_scr[y] + x, (VIRT_SCR_COLS - x), fmt, __VA_ARGS__);
#define virt_scr_fill_from(n, y) \
  (n < 20 && snprintf(g_virt_scr[y] + n, (VIRT_SCR_COLS - n), fill_pattern + n));

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
| Confirmation screen     | - Pressione START para confirmar...              | Waits for user confirmation before an action                          | START: Confirmation                                                                         |
*/
#define INFORM_SCREEN 0
#define NETWRK_SCREEN 1
#define NETCFG_SCREEN 2
#define READER_SCREEN 3
#define SYSTEM_SCREEN 4
#define UPLOAD_SCREEN 5
#define BACKUP_SCREEN 6
#define DELETE_SCREEN 7
#define SHTDWN_SCREEN 8
#define NAV_SCREENS_COUNT 9

#define OFFMSG_SCREEN 9
#define CONFRM_SCREEN 10
#define WAITNG_SCREEN 11
#define SCREENS_COUNT 12
unsigned int g_current_screen = 0;
unsigned int g_confirm_target = 0; // target screen for events that need confirmation

bool    g_locked;
bool    g_screen_waiting_confirmation;
int32_t g_screen_waiting_timestamp;

const char desc[SCREENS_COUNT][VIRT_SCR_COLS] = {
  "START:Reset tela   ",
  "                   ",
  "START:Reconectar   ",
  "                   ",
  "START:Atualizar    ",
  "START:Upload Regist",
  "START:Upload Backup",
  "START:Apagar tudo  ",
  "START:Desligar     ",
  "                   ",
  "START:Confirma     ",
  "                   "
};

void
screen_build()
{
  unsigned int l1 = 0, l2 = 0;

  if (g_current_screen < 5) {
    virt_scr_sprintf(0, 0, "PORTAL my50x", 0);
  } else
    virt_scr_fill_from(0, 0);

  switch (g_current_screen) {
    case INFORM_SCREEN:
      l1 = virt_scr_sprintf(0, 1, "Regist.: %" PRId32, g_system_data.tags);
      l2 = virt_scr_sprintf(0, 2, "Atletas: %" PRId32, g_system_data.unique_tags);
      break;
    case NETWRK_SCREEN:
      l2 = virt_scr_sprintf(0, 2, "Comunicando: %3s", g_system_data.comm_status ? "SIM" : "NAO");
      break;
    case NETCFG_SCREEN:
      l1 = virt_scr_sprintf(0, 1, "Wi-Fi: %2s", g_system_data.wifi_status ? "OK" : "X");
      l2 = virt_scr_sprintf(0, 2, "LTE/4G: %2s", g_system_data.lte4_status ? "OK" : "X");
      break;
    case READER_SCREEN:
      l1 = virt_scr_sprintf(0, 1, "Leitor: %2s", g_system_data.rfid_status ? "OK" : "X");
      break;
    case SYSTEM_SCREEN:
      l1 = virt_scr_sprintf(0, 1, "Versao: %" PRId32, g_system_data.sys_version);
      break;
    case UPLOAD_SCREEN:
      l1 = virt_scr_sprintf(0, 1, "Atletas: %" PRId32, g_system_data.unique_tags);
      //l2 = virt_scr_sprintf(0, 2, "Pendentes: %" PRId32, g_system_data.envios);
      break;
    case BACKUP_SCREEN:
      l1 = virt_scr_sprintf(0, 1, "Backups: %" PRId32, g_system_data.backups);
      break;
    case DELETE_SCREEN:
      l1 = virt_scr_sprintf(0, 1, "Apagar dados", NULL);
      l2 = virt_scr_sprintf(0, 2, "do equipamento", NULL);
      break;
    case SHTDWN_SCREEN:
      l1 = virt_scr_sprintf(0, 1, "Desligar o", NULL);
      l2 = virt_scr_sprintf(0, 2, "equipamento", NULL);
      break;
    
    /* end of NAV_SCREENS */
    /* Extra screens */

    case OFFMSG_SCREEN:
      l1 = virt_scr_sprintf(0, 1, "Aguarde 30 segundos", NULL);
      l2 = virt_scr_sprintf(0, 2, "E pressione POWER", NULL);
      break;
    case CONFRM_SCREEN:
      l1 = virt_scr_sprintf(0, 1, "Pressione START", NULL);
      l2 = virt_scr_sprintf(0, 2, "para confirmar", NULL);
      break;
    case WAITNG_SCREEN:
      l2 = virt_scr_sprintf(0, 2, "Aguarde...", NULL);
      break;
  }

  virt_scr_fill_from(l1, 1);
  virt_scr_fill_from(l2, 2);

  virt_scr_sprintf(0, 3, "%s", desc[g_current_screen]);
}

void
screen_unlock()
{
  if (!g_locked) return;

  g_current_screen = 0;
  g_locked = false;
}

void
screen_lock()
{
  g_current_screen = WAITNG_SCREEN;
  g_locked = true;
}

void
screen_next()
{
  g_current_screen = (g_current_screen + 1) % NAV_SCREENS_COUNT;
}

void
screen_confirm()
{
  g_screen_waiting_timestamp = millis();
  g_screen_waiting_confirmation = true;
  g_confirm_target = g_current_screen;
  g_current_screen = CONFRM_SCREEN;
}

void
screen_wait_confirm()
{
  if (check_clicked()) {
    g_screen_waiting_confirmation = false;
    event_send();
  }

  if (millis() - g_screen_waiting_timestamp > 2000) {
    g_screen_waiting_confirmation = false;
    g_current_screen = g_confirm_target;
  }
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

#define START_DELIMITER 0x3C
#define END_DELIMITER 0x3E
void
event_send()
{
  // screens that need confirmation
  if (g_current_screen == DELETE_SCREEN || g_current_screen == SHTDWN_SCREEN) {
    screen_confirm();
    return;
  }

  if (g_current_screen == CONFRM_SCREEN)
    g_current_screen = g_confirm_target;

  Serial.write(START_DELIMITER);
  Serial.write(g_current_screen);
  Serial.write(END_DELIMITER);

  screen_lock();
}

#define CAP 256
void
handle_serial()
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

  parse_data(buf);

  // a successful data receival toggles an UNLOCK
  screen_unlock();
}

void
handle_buttons()
{
  // LOCKED
  if (g_locked) return;

  switch (check_clicked()) {
    case BUTTON_VANCE:
      screen_next();
      break;
    case BUTTON_START:
      event_send();
      break;
  }
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
  handle_serial();
  if (g_locked)                      goto skip_screen_tasks;
  if (g_screen_waiting_confirmation) goto wait_confirm;

  handle_buttons();
  screen_draw();
  delay(50);
  return;

skip_screen_tasks:
  delay(200);
  return;

wait_confirm:
  screen_wait_confirm();
  screen_draw();
  delay(50);
}
