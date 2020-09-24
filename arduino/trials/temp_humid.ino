#include <DHT.h>
#include <DHT_U.h>

#define DHT_PIN 7
#define DHT_TYPE DHT22
DHT dht(DHT_PIN, DHT_TYPE);

int chk;
float hum;
float temp;

void setup() {
  Serial.begin(9600);
  dht.begin();
}

void loop() {
  delay(1000);
  hum = dht.readHumidity();
  // true to indicate Fahrenheit
  temp = dht.readTemperature(true);

  Serial.print("Humidity: ");
  Serial.print(hum);
  Serial.print("% | ");
  Serial.print("Temp: ");
  Serial.print(temp);
  Serial.println("F");
}
