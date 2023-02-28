package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bogosj/tesla"
	"golang.org/x/oauth2"
)

func teslaCli(token string, args []string) error {
	tok := new(oauth2.Token)
	tok.AccessToken = token
	c, err := tesla.NewClient(context.Background(), tesla.WithToken(tok))
	if err != nil {
		return err
	}
	v, err := c.Vehicles()
	if err != nil {
		return err
	}

	for i, v := range v {
		if i > 0 {
			fmt.Println("----")
		}
		fmt.Printf("ID: %v\n", v.ID)
		fmt.Printf("Color: %v\n", v.Color)
		fmt.Printf("DisplayName: %v\n", v.DisplayName)
		fmt.Printf("OptionCodes: %v\n", v.OptionCodes)
		fmt.Printf("VehicleID: %v\n", v.VehicleID)
		fmt.Printf("VIN: %s\n", v.Vin)
		fmt.Printf("Tokens: %v\n", v.Tokens)
		fmt.Printf("State: %v\n", v.State)
		fmt.Printf("IDS: %v\n", v.IDS)
		fmt.Printf("RemoteStartEnabled: %v\n", v.RemoteStartEnabled)
		fmt.Printf("CalendarEnabled: %v\n", v.CalendarEnabled)
		fmt.Printf("NotificationsEnabled: %v\n", v.NotificationsEnabled)
		fmt.Printf("BackseatToken: %v\n", v.BackseatToken)
		fmt.Printf("BackseatTokenUpdatedAt: %v\n", v.BackseatTokenUpdatedAt)
		fmt.Printf("AccessType: %v\n", v.AccessType)
		fmt.Printf("InService: %v\n", v.InService)
		fmt.Printf("APIVersion: %v\n", v.APIVersion)
		fmt.Printf("CommandSigning: %v\n", v.CommandSigning)
		if v.VehicleConfig != nil {
			fmt.Printf("VehicleConfig.CanAcceptNavigationRequests: %v\n", v.VehicleConfig.CanAcceptNavigationRequests)
			fmt.Printf("VehicleConfig.CanActuateTrunks: %v\n", v.VehicleConfig.CanActuateTrunks)
			fmt.Printf("VehicleConfig.CarSpecialType: %v\n", v.VehicleConfig.CarSpecialType)
			fmt.Printf("VehicleConfig.CarType: %v\n", v.VehicleConfig.CarType)
			fmt.Printf("VehicleConfig.ChargePortType: %v\n", v.VehicleConfig.ChargePortType)
			fmt.Printf("VehicleConfig.DefaultChargeToMax: %v\n", v.VehicleConfig.DefaultChargeToMax)
			fmt.Printf("VehicleConfig.DriverAssist: %v\n", v.VehicleConfig.DriverAssist)
		}
		// v.AutoparkAbort(): causes the vehicle to abort the Autopark request
		// v.AutoparkForward(): causes the vehicle to pull forward
		// v.AutoparkReverse(): causes the vehicle to go in reverse
		charge, err := v.ChargeState()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ChargeState(): %s\n", err.Error())
		} else {
			fmt.Printf("Charge.ChargingState: %v\n", charge.ChargingState)
			fmt.Printf("Charge.ChargeLimitSoc: %v\n", charge.ChargeLimitSoc)
			fmt.Printf("Charge.ChargeLimitSocStd: %v\n", charge.ChargeLimitSocStd)
			fmt.Printf("Charge.ChargeLimitSocMin: %v\n", charge.ChargeLimitSocMin)
			fmt.Printf("Charge.ChargeLimitSocMax: %v\n", charge.ChargeLimitSocMax)
			fmt.Printf("Charge.ChargeToMaxRange: %v\n", charge.ChargeToMaxRange)
			fmt.Printf("Charge.BatteryHeaterOn: %v\n", charge.BatteryHeaterOn)
			fmt.Printf("Charge.NotEnoughPowerToHeat: %v\n", charge.NotEnoughPowerToHeat)
			fmt.Printf("Charge.MaxRangeChargeCounter: %v\n", charge.MaxRangeChargeCounter)
			fmt.Printf("Charge.FastChargerPresent: %v\n", charge.FastChargerPresent)
			fmt.Printf("Charge.FastChargerType: %v\n", charge.FastChargerType)
			fmt.Printf("Charge.BatteryRange: %v\n", charge.BatteryRange)
			fmt.Printf("Charge.EstBatteryRange: %v\n", charge.EstBatteryRange)
			fmt.Printf("Charge.IdealBatteryRange: %v\n", charge.IdealBatteryRange)
			fmt.Printf("Charge.BatteryLevel: %v\n", charge.BatteryLevel)
			fmt.Printf("Charge.UsableBatteryLevel: %v\n", charge.UsableBatteryLevel)
			fmt.Printf("Charge.ChargeEnergyAdded: %v\n", charge.ChargeEnergyAdded)
			fmt.Printf("Charge.ChargeMilesAddedRated: %v\n", charge.ChargeMilesAddedRated)
			fmt.Printf("Charge.ChargeMilesAddedIdeal: %v\n", charge.ChargeMilesAddedIdeal)
			fmt.Printf("Charge.ChargerVoltage: %v\n", charge.ChargerVoltage)
			fmt.Printf("Charge.ChargerPilotCurrent: %v\n", charge.ChargerPilotCurrent)
			fmt.Printf("Charge.ChargerActualCurrent: %v\n", charge.ChargerActualCurrent)
			fmt.Printf("Charge.ChargerPower: %v\n", charge.ChargerPower)
			fmt.Printf("Charge.TimeToFullCharge: %v\n", charge.TimeToFullCharge)
			fmt.Printf("Charge.TripCharging: %v\n", charge.TripCharging)
			fmt.Printf("Charge.ChargeRate: %v\n", charge.ChargeRate)
			fmt.Printf("Charge.ChargePortDoorOpen: %v\n", charge.ChargePortDoorOpen)
			fmt.Printf("Charge.MotorizedChargePort: %v\n", charge.MotorizedChargePort)
			fmt.Printf("Charge.ScheduledChargingMode: %v\n", charge.ScheduledChargingMode)
			fmt.Printf("Charge.ScheduledDepatureTime: %v\n", charge.ScheduledDepatureTime)
			fmt.Printf("Charge.ScheduledChargingStartTime: %v\n", charge.ScheduledChargingStartTime)
			fmt.Printf("Charge.ScheduledChargingPending: %v\n", charge.ScheduledChargingPending)
			fmt.Printf("Charge.UserChargeEnableRequest: %v\n", charge.UserChargeEnableRequest)
			fmt.Printf("Charge.ChargeEnableRequest: %v\n", charge.ChargeEnableRequest)
			fmt.Printf("Charge.EuVehicle: %v\n", charge.EuVehicle)
			fmt.Printf("Charge.ChargerPhases: %v\n", charge.ChargerPhases)
			fmt.Printf("Charge.ChargePortLatch: %v\n", charge.ChargePortLatch)
			fmt.Printf("Charge.ChargeCurrentRequest: %v\n", charge.ChargeCurrentRequest)
			fmt.Printf("Charge.ChargeCurrentRequestMax: %v\n", charge.ChargeCurrentRequestMax)
			fmt.Printf("Charge.ChargeAmps: %v\n", charge.ChargeAmps)
			fmt.Printf("Charge.OffPeakChargingEnabled: %v\n", charge.OffPeakChargingEnabled)
			fmt.Printf("Charge.OffPeakChargingTimes: %v\n", charge.OffPeakChargingTimes)
			fmt.Printf("Charge.OffPeakHoursEndTime: %v\n", charge.OffPeakHoursEndTime)
			fmt.Printf("Charge.PreconditioningEnabled: %v\n", charge.PreconditioningEnabled)
			fmt.Printf("Charge.PreconditioningTimes: %v\n", charge.PreconditioningTimes)
			fmt.Printf("Charge.ManagedChargingActive: %v\n", charge.ManagedChargingActive)
			fmt.Printf("Charge.ManagedChargingUserCanceled: %v\n", charge.ManagedChargingUserCanceled)
			fmt.Printf("Charge.ManagedChargingStartTime: %v\n", charge.ManagedChargingStartTime)
			fmt.Printf("Charge.ChargePortcoldWeatherMode: %v\n", charge.ChargePortcoldWeatherMode)
			fmt.Printf("Charge.ConnChargeCable: %v\n", charge.ConnChargeCable)
			fmt.Printf("Charge.FastChargerBrand: %v\n", charge.FastChargerBrand)
			fmt.Printf("Charge.MinutesToFullCharge: %v\n", charge.MinutesToFullCharge)
		}

		climate, err := v.ClimateState()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ClimateState(): %s\n", err.Error())
		} else {
			fmt.Printf("Climate.InsideTemp: %v\n", climate.InsideTemp)
			fmt.Printf("Climate.OutsideTemp: %v\n", climate.OutsideTemp)
			fmt.Printf("Climate.DriverTempSetting: %v\n", climate.DriverTempSetting)
			fmt.Printf("Climate.PassengerTempSetting: %v\n", climate.PassengerTempSetting)
			fmt.Printf("Climate.LeftTempDirection: %v\n", climate.LeftTempDirection)
			fmt.Printf("Climate.RightTempDirection: %v\n", climate.RightTempDirection)
			fmt.Printf("Climate.IsAutoConditioningOn: %v\n", climate.IsAutoConditioningOn)
			fmt.Printf("Climate.IsFrontDefrosterOn: %v\n", climate.IsFrontDefrosterOn)
			fmt.Printf("Climate.IsRearDefrosterOn: %v\n", climate.IsRearDefrosterOn)
			fmt.Printf("Climate.FanStatus: %v\n", climate.FanStatus)
			fmt.Printf("Climate.IsClimateOn: %v\n", climate.IsClimateOn)
			fmt.Printf("Climate.MinAvailTemp: %v\n", climate.MinAvailTemp)
			fmt.Printf("Climate.MaxAvailTemp: %v\n", climate.MaxAvailTemp)
			fmt.Printf("Climate.SeatHeaterLeft: %v\n", climate.SeatHeaterLeft)
			fmt.Printf("Climate.SeatHeaterRight: %v\n", climate.SeatHeaterRight)
			fmt.Printf("Climate.SeatHeaterRearLeft: %v\n", climate.SeatHeaterRearLeft)
			fmt.Printf("Climate.SeatHeaterRearRight: %v\n", climate.SeatHeaterRearRight)
			fmt.Printf("Climate.SeatHeaterRearCenter: %v\n", climate.SeatHeaterRearCenter)
			fmt.Printf("Climate.SeatHeaterRearRightBack: %v\n", climate.SeatHeaterRearRightBack)
			fmt.Printf("Climate.SeatHeaterRearLeftBack: %v\n", climate.SeatHeaterRearLeftBack)
			fmt.Printf("Climate.SmartPreconditioning: %v\n", climate.SmartPreconditioning)
			fmt.Printf("Climate.BatteryHeater: %v\n", climate.BatteryHeater)
			fmt.Printf("Climate.BatteryHeaterNoPower: %v\n", climate.BatteryHeaterNoPower)
			fmt.Printf("Climate.ClimateKeeperMode: %v\n", climate.ClimateKeeperMode)
			fmt.Printf("Climate.DefrostMode: %v\n", climate.DefrostMode)
			fmt.Printf("Climate.IsPreconditioning: %v\n", climate.IsPreconditioning)
			fmt.Printf("Climate.RemoteHeaterControlEnabled: %v\n", climate.RemoteHeaterControlEnabled)
			fmt.Printf("Climate.SideMirrorHeaters: %v\n", climate.SideMirrorHeaters)
			fmt.Printf("Climate.WiperBladeHeater: %v\n", climate.WiperBladeHeater)
		}

		drive, err := v.DriveState()
		if err != nil {
			fmt.Fprintf(os.Stderr, "DriveState(): %s\n", err.Error())
		} else {
			fmt.Printf("Drive.ShiftState: %v\n", drive.ShiftState)
			fmt.Printf("Drive.Speed: %v\n", drive.Speed)
			fmt.Printf("Drive.Latitude: %v\n", drive.Latitude)
			fmt.Printf("Drive.Longitude: %v\n", drive.Longitude)
			fmt.Printf("Drive.Heading: %v\n", drive.Heading)
			fmt.Printf("Drive.GpsAsOf: %v\n", drive.GpsAsOf)
			fmt.Printf("Drive.NativeLatitude: %v\n", drive.NativeLatitude)
			fmt.Printf("Drive.NativeLocationSupported: %v\n", drive.NativeLocationSupported)
			fmt.Printf("Drive.NativeLongitude: %v\n", drive.NativeLongitude)
			fmt.Printf("Drive.NativeType: %v\n", drive.NativeType)
			fmt.Printf("Drive.Power: %v\n", drive.Power)
		}

		// v.EnableSentry(): enables Sentry Mode
		// v.FlashLights(): flashes the lights of the vehicle
		gui, err := v.GuiSettings()
		if err != nil {
			fmt.Fprintf(os.Stderr, "GuiSettings(): %s\n", err.Error())
		} else {
			fmt.Printf("Gui.GuiDistanceUnits: %v\n", gui.GuiDistanceUnits)
			fmt.Printf("Gui.GuiTemperatureUnits: %v\n", gui.GuiTemperatureUnits)
			fmt.Printf("Gui.GuiChargeRateUnits: %v\n", gui.GuiChargeRateUnits)
			fmt.Printf("Gui.Gui24HourTime: %v\n", gui.Gui24HourTime)
			fmt.Printf("Gui.GuiRangeDisplay: %v\n", gui.GuiRangeDisplay)
			fmt.Printf("Gui.ShowRangeUnits: %v\n", gui.ShowRangeUnits)
		}

		// v.HonkHorn(): honks the horn of the vehicle
		// v.LockDoors(): locks the doors of the vehicle

	}

	return nil
}
