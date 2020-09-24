int DOOR_PIN = 7;
int PREV_STATE = -1;
int CUR_STATE = -1;
int CLOSED = 1;
int OPEN = 0;

void setup() {
  Serial.begin(9600);
  pinMode(7, INPUT_PULLUP); //Door Sensor is connected to D7
}

void loop() {
  delay(1000);
  int state = !digitalRead(DOOR_PIN);
  String status;

  if (state != PREV_STATE) {
    Serial.print("State Change. New State: ");
    Serial.println(state);
  }

  if (state != CUR_STATE) {
    if (state == OPEN) {
      Serial.println("Door has been opened!");
    }

    if (state == CLOSED) {
      Serial.println("Door is now closed");
    }
//    Serial.print("Current State: ");
//    Serial.print(CUR_STATE);
//    Serial.print(" State: ");
//    Serial.println(state);
    CUR_STATE = state;
  }

  PREV_STATE = state;
}
