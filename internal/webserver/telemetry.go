package webserver

// Telemetry represents a single UDP message from the telemetry system.
// For fields like Acceleration, X = right, Y = up, Z = forward.
// For fields like AngularVelocity, X = pitch, Y = yaw, Z = roll.
type Telemetry struct {
	IsRaceOn         int32
	TimestampMS      uint32
	EngineMaxRpm     float32
	EngineIdleRpm    float32
	CurrentEngineRpm float32

	AccelerationX float32 // In the car's local space; X = right, Y = up, Z = forward
	AccelerationY float32
	AccelerationZ float32

	VelocityX float32 // In the car's local space; X = right, Y = up, Z = forward
	VelocityY float32
	VelocityZ float32

	AngularVelocityX float32 // In the car's local space; X = pitch, Y = yaw, Z = roll
	AngularVelocityY float32
	AngularVelocityZ float32

	Yaw   float32
	Pitch float32
	Roll  float32

	NormalizedSuspensionTravelFrontLeft  float32 // Suspension travel normalized: 0.0f = max stretch; 1.0 = max compression
	NormalizedSuspensionTravelFrontRight float32
	NormalizedSuspensionTravelRearLeft   float32
	NormalizedSuspensionTravelRearRight  float32

	TireSlipRatioFrontLeft  float32 // Tire normalized slip ratio, 0 means 100% grip and |ratio| > 1.0 means loss of grip
	TireSlipRatioFrontRight float32
	TireSlipRatioRearLeft   float32
	TireSlipRatioRearRight  float32

	WheelRotationSpeedFrontLeft  float32 // Wheel rotation speed in radians/sec.
	WheelRotationSpeedFrontRight float32
	WheelRotationSpeedRearLeft   float32
	WheelRotationSpeedRearRight  float32

	WheelOnRumbleStripFrontLeft  int32 // = 1 when wheel is on rumble strip, = 0 when off.
	WheelOnRumbleStripFrontRight int32
	WheelOnRumbleStripRearLeft   int32
	WheelOnRumbleStripRearRight  int32

	WheelInPuddleDepthFrontLeft  float32 // = from 0 to 1, where 1 is the deepest puddle.
	WheelInPuddleDepthFrontRight float32
	WheelInPuddleDepthRearLeft   float32
	WheelInPuddleDepthRearRight  float32

	SurfaceRumbleFrontLeft  float32 // Non-dimensional surface rumble values passed to controller force feedback
	SurfaceRumbleFrontRight float32
	SurfaceRumbleRearLeft   float32
	SurfaceRumbleRearRight  float32

	TireSlipAngleFrontLeft  float32 // Tire normalized slip angle, = 0 means 100% grip and |angle| > 1.0 means loss of grip
	TireSlipAngleFrontRight float32
	TireSlipAngleRearLeft   float32
	TireSlipAngleRearRight  float32

	TireCombinedSlipFrontLeft  float32 // Tire normalized combined slip, = 0 means 100% grip and |slip| > 1.0 means loss of grip
	TireCombinedSlipFrontRight float32
	TireCombinedSlipRearLeft   float32
	TireCombinedSlipRearRight  float32

	SuspensionTravelMetersFrontLeft  float32 // Actual suspension travel in meters
	SuspensionTravelMetersFrontRight float32
	SuspensionTravelMetersRearLeft   float32
	SuspensionTravelMetersRearRight  float32

	CarOrdinal          int32 // Unique ID of the car make/model
	CarClass            int32 // Between 0 (D -- worst cars) and 7 (X class -- best cars) inclusive
	CarPerformanceIndex int32 // Between 100 (slowest car) and 999 (fastest car) inclusive
	DrivetrainType      int32 // Corresponds to EDrivetrainType; 0 = FWD, 1 = RWD, 2 = AWD
	NumCylinders        int32 // Number of cylinders in the engine

	// The following are part of V2.   They are likely missing from Forza 7 if the data out type is set to "SLED"
	PositionX float32 // Position (meters)
	PositionY float32
	PositionZ float32

	Speed  float32 // meters / sec
	Power  float32 // watts
	Torque float32 // newton meters

	TireTempFrontLeft  float32
	TireTempFrontRight float32
	TireTempRearLeft   float32
	TireTempRearRight  float32

	Boost            float32
	Fuel             float32
	DistanceTraveled float32
	BestLap          float32
	LastLap          float32
	CurrentLap       float32
	CurrentRaceTime  float32

	LapNumber    uint16
	RacePosition uint8

	Accel     uint8
	Brake     uint8
	Clutch    uint8
	HandBrake uint8
	Gear      uint8
	Steer     int8

	NormalizedDrivingLine       int8
	NormalizedAIBrakeDifference int8
}
