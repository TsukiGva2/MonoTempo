static float calculatePinVoltage(int analog_in_pin,
				 float adc_voltage,
				 float R1, float R2,
				 float ref_voltage)
{
	int adc_value = analogRead(analog_in_pin);

	adc_voltage = (adc_value * ref_voltage) / 1024.0;
	return adc_voltage * (R1 + R2) / R2;
}
