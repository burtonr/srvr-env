// Relay Reference: https://www.circuitbasics.com/setting-up-a-5v-relay-on-the-arduino/
// Button Reference: https://www.instructables.com/Arduino-Button-Tutorial/

#define BUTTON_PIN 2
#define RELAY_PIN 10
#define DELAY 20  // Delay per loop in ms

void setup() {
  pinMode(BUTTON_PIN, INPUT);
  pinMode(RELAY_PIN, OUTPUT);
  digitalWrite(BUTTON_PIN, HIGH);
  Serial.begin(9600);
}

boolean handle_button()
{
  int button_pressed = !digitalRead(BUTTON_PIN); // pin low -> pressed
  return button_pressed;
}

void loop() {
  // handle button
  boolean button_pressed = handle_button();

  // do other things
  Serial.print(button_pressed ? "^" : ".");

  if (button_pressed) {
    digitalWrite(RELAY_PIN, HIGH);
  } else {
    digitalWrite(RELAY_PIN, LOW);
  }
  
  static int counter = 0;
  if ((++counter & 0x3f) == 0)
    Serial.println();

  delay(DELAY);
}