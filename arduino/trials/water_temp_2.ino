#include<DS18B20.h>
// https://github.com/matmunk/DS18B20
DS18B20 ds(2);

void setup() {
  Serial.begin(9600);
  Serial.print("Devices: ");
  Serial.println(ds.getNumberOfDevices());
  Serial.println();

}

void loop() {
    while (ds.selectNext()) {
      uint8_t address[8];
      ds.getAddress(address);

      Serial.print("Address:");
      for (uint8_t i = 0; i < 8; i++) {
        Serial.print(" ");
        Serial.print(address[i]);
      }
      Serial.print(" | ");
      Serial.print("Temperature: ");
      Serial.print(ds.getTempC());
      Serial.print(" C / ");
      Serial.print(ds.getTempF());
      Serial.println(" F");
    }
    delay(1000);
}
