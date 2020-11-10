#define BUTTON_PIN 2
#define RELAY_PIN 10
#define DELAY 20  // Delay per loop in ms

String inputString = "";
boolean inputComplete = false;
String command = "";


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
  
  inputEvent();
  if (inputComplete) {
    processInput(inputString);
    inputString = "";
    inputComplete = false;
  }

  if (command == "ON" || button_pressed) {
    digitalWrite(RELAY_PIN, HIGH);
  } else {
    digitalWrite(RELAY_PIN, LOW);
  }

  delay(DELAY);
}

void inputEvent() {
  while (Serial.available()) {
    // get the new byte:
    char inChar = (char)Serial.read();
    // add it to the inputString:
    inputString += inChar;
    // if the incoming character is a newline, set a flag
    // so the main loop can do something about it:
    if (inChar == '\n') {
      inputComplete = true;
    }
  }
}

void processInput(String input) {
//  FROM: https://github.com/0bscur3/Arduino-LedStripe-IR
  if(input.startsWith("cmd:",0)){
    int start = input.indexOf(":")+1;
    int end = input.length()-1;

    String cmd = input.substring(start, end);
   
    if(cmd == "POWER_ON"){
      command = "ON";
    }else if(cmd == "POWER_OFF"){
      command = "OFF";
    }
  }
}