int led1 = 9; //define pin for first LED
int led2 = 10; //define pin for second LED
int button1 = 2; //define pin for first button
int button2 = 3; //define pin for second button
int button3 = 4; //define pin for third button
int potentiometer = A0; //define pin for potentiometer

void setup() {
  pinMode(led1, OUTPUT);
  pinMode(led2, OUTPUT);
  pinMode(button1, INPUT_PULLUP);
  pinMode(button2, INPUT_PULLUP);
  pinMode(button3, INPUT_PULLUP);
}

void loop() {
  analogWrite(led1, 128); //turn on first LED at 50% duty cycle
  analogWrite(led2, 255); //turn on second LED at 100% duty cycle
  
  int button1Value = digitalRead(button1); //read value of first button
  int button2Value = digitalRead(button2); //read value of second button
  int button3Value = digitalRead(button3); //read value of third button
  int potentiometerValue = analogRead(potentiometer); //read value of potentiometer

  //do something with the button and potentiometer values
  delay(500); //wait for half a second
}