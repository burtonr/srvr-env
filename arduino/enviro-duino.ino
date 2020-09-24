float temp;
int tempPin = 0;

String inputString = "";
boolean inputComplete = false;

void setup() {
  // put your setup code here, to run once:
  Serial.begin(9600);
}

void loop() {
  inputEvent();
  if (inputComplete) {
    processInput(inputString);
    inputString = "";
    inputComplete = false;
  }
  
  float temp = getFTemp(tempPin);
  writeTemp("TMP1", temp);
  delay(6000);
}

float getVoltage(int pin){
  return (analogRead(pin) * .004882814);
}

float getFTemp(int pin) {
  float v = getVoltage(pin);
  return (((v -.5) * 100L) *9.0/5.0) + 32.0;
}

void writeTemp(String sensor, float temp) {
  Serial.print("[");
  Serial.print(sensor);
  Serial.print("] ");
  Serial.println(temp);
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
  Serial.print("Input received: ");
  Serial.println(input);
//  FROM: https://github.com/0bscur3/Arduino-LedStripe-IR
//  if(cmd.startsWith("cmd:",0)){
//    int start = cmd.indexOf(":")+1;
//    int end = cmd.length()-1;
//
//    String command = cmd.substring(start, end);
//   
//    if(command == "POWER_ON"){
//      powerOn();
//    }else if(command == "POWER_OFF"){
//      powerOff();
//    }
//  }
}
