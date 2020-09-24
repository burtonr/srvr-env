#include<OneWire.h>
#include<DallasTemperature.h>

#define ONE_WIRE_BUS 2

OneWire oneWire(ONE_WIRE_BUS);
DallasTemperature sensors(&oneWire);

void setup(void) {
  Serial.begin(9600);
  Serial.println("Dallas Temperature IC Control Library Demo");
  sensors.begin();
}

void loop(void) {
  Serial.print("Requesting temperatures...");
  sensors.requestTemperatures();
  Serial.print("Temperature is: ");
  float temp = sensors.getTempFByIndex(0);
  Serial.println(temp);
  delay(1000);
}
