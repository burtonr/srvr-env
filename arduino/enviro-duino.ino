#include<DS18B20.h>
#include <DHT.h>
#include <DHT_U.h>

// https://github.com/matmunk/DS18B20
DS18B20 ds(2);

// https://github.com/adafruit/DHT-sensor-library
#define DHT_PIN 3
#define DHT_TYPE DHT22
DHT dht(DHT_PIN, DHT_TYPE);

int timer;
String inputString = "";
boolean inputComplete = false;

void setup() {
  // put your setup code here, to run once:
  Serial.begin(9600);
  dht.begin();
}

void loop() {
  int mainDelay = 1000; // every second
  int innerDelay = 60000; // every minute
//  int innerDelay = 2000; // for testing...
  
  inputEvent();
  if (inputComplete) {
    processInput(inputString);
    inputString = "";
    inputComplete = false;
  }
  
  if (timer == innerDelay) {
    // Things to run at a longer interval
    readTemperatures();
    readHumidity();
    timer = 0;
  } else {
    timer += mainDelay;
  }

  delay(mainDelay);
}

void readTemperatures() {
  while (ds.selectNext()) {
    uint8_t address[8];
    String addr = "";
    char tempBuffer[16];
    
    ds.getAddress(address);
    for (uint8_t i = 0; i < 8; i++){
      addr += address[i];
    }

    float temp = ds.getTempF();
    String output = formatOutput(addr, temp);
    
    Serial.println(output);
  }
}

void readHumidity() {
  char humBuffer[8];
  String name = "DHT22";
  
  float hum = dht.readHumidity();
  String humOutput = formatOutput(name, hum);

  float temp = dht.readTemperature(true); // true for F

  Serial.println(humOutput);
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

String formatOutput(String name, float value) {
  char valueBuff[16];
  
  dtostrf(value, 6, 2, valueBuff);

  String output = "[" + name + "] " + valueBuff;
  return output;
}
